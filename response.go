package getdns

import (
	"fmt"
	"github.com/miekg/dns"
	"time"
)

/*
 * Response Dictionary structure: this needs to eventually look
 * like the response dict in the API.
 */
type ResponseDict struct {
	Status      uint16
	Cname       string
	RepliesTree []*Reply
}

type Reply struct {
	Qname        string
	Qtype        string
	Qclass       string
	Transport    string
	Truncated    bool
	Retried      bool
	Timeout      bool
	Rtt          time.Duration
	Err          error
	Msg          *dns.Msg
	Status       uint16
	StatusDNSSEC uint16
}

/*
 * printResponseDict()
 */

func PrintResponseDict(response *ResponseDict) {

	for i, reply := range response.RepliesTree {
		fmt.Printf(">>> Response %d: %s %s %s\n", i, reply.Qname, reply.Qtype, reply.Qclass)
		fmt.Printf("Transport: %s\n", reply.Transport)
		fmt.Println("ResponseTime:", reply.Rtt)
		fmt.Println(reply.Msg)
	}
	return
}

/*
 * JustAddresses() - return list of human readable IP address strings in
 *                   answer section of DNS replies in response dictionary.
 *                   If we implement an API compliant response dict, then
 *                   the output of this function will be a component of that
 *                   dictionary.
 */

func JustAddresses(response *ResponseDict) []string {

	var result []string

	for _, reply := range response.RepliesTree {
		for _, rr := range reply.Msg.Answer {
			if ip4, ok := rr.(*dns.A); ok {
				result = append(result, ip4.A.String())
			} else if ip6, ok := rr.(*dns.AAAA); ok {
				result = append(result, ip6.AAAA.String())
			}
		}
	}

	return result
}
