package getdns

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strings"
	"time"
)

/*
 * makeMessage() - construct DNS message structure
 */
func makeMessage(c *Context, qname, qtype, qclass string, ext Extension) *dns.Msg {

	m := new(dns.Msg)
	m.Id = dns.Id()

	if c.restype == RESOLUTION_STUB {
		m.RecursionDesired = true
	} else {
		m.RecursionDesired = false
	}

	if c.adflag {
		m.AuthenticatedData = true
	}

	if c.cdflag {
		m.CheckingDisabled = true
	}

	if ext["dnssec_return_status"] || ext["dnssec_return_only_secure"] || ext["dnssec_return_validation_chain"] {
		opt := new(dns.OPT)
		opt.Hdr.Name = "."
		opt.Hdr.Rrtype = dns.TypeOPT
		opt.SetDo()
		m.Extra = append(m.Extra, opt)
	}

	m.Question = make([]dns.Question, 1)
	qtype_int, ok := dns.StringToType[strings.ToUpper(qtype)]
	if !ok {
		fmt.Printf("%s: Unrecognized query type.\n", qtype)
		return nil
	}
	qclass_int, ok := dns.StringToClass[strings.ToUpper(qclass)]
	if !ok {
		fmt.Printf("%s: Unrecognized query class.\n", qclass)
		return nil
	}
	m.Question[0] = dns.Question{qname, qtype_int, qclass_int}

	return m
}

/*
 * doQuery() - perform DNS query with timeouts and retries as needed
 */
func doQuery(ctx *Context, qname, qtype, qclass string, ext Extension) (r *Reply, err error) {

	var retries = ctx.retries
	var timeout = ctx.timeout

	qname = dns.Fqdn(qname)
	r = new(Reply)
	r.Qname = qname
	r.Qtype = qtype
	r.Qclass = qclass

	m := makeMessage(ctx, qname, qtype, qclass, ext)

	if ctx.tcp {
		r.Transport = "tcp"
		r.Msg, r.Rtt, r.Err = sendRequest(ctx, m, "tcp", timeout)
		return r, r.Err
	}

	r.Transport = "udp"
	for retries > 0 {
		r.Msg, r.Rtt, r.Err = sendRequest(ctx, m, "udp", timeout)
		if r.Err == nil {
			break
		}
		if r.Err == dns.ErrId {
			retries--
			continue
		}
		if nerr, ok := r.Err.(net.Error); ok && !nerr.Timeout() {
			break
		}
		retries--
		timeout = timeout * 2
	}

	if r.Err == dns.ErrTruncated {
		fmt.Println("Truncated; Retrying over TCP ..")
		r.Truncated = true
		r.Retried = true
		r.Transport = "tcp"
		r.Msg, r.Rtt, r.Err = sendRequest(ctx, m, "tcp", timeout)
	}

	return r, r.Err

}

/*
 * sendRequest() - send a DNS query
 */
func sendRequest(ctx *Context, m *dns.Msg, transport string, timeout time.Duration) (response *dns.Msg, rtt time.Duration, err error) {

	c := new(dns.Client)
	c.Timeout = timeout
	c.Net = transport // "udp" or "tcp"
	response, rtt, err = c.Exchange(m, addressString(ctx.server, ctx.port))
	return

}
