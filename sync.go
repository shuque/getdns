/*
 * Synchronous query functions
 */

package getdns

import (
	"fmt"
	"github.com/miekg/dns"
)

/*
 * General()
 */

func General(ctx *Context, qname, qtype string, ext Extension) (ReturnType, *ResponseDict) {

	rd := new(ResponseDict)
	reply, err := doQuery(ctx, qname, qtype, "IN", ext)
	if err != nil {
		return RETURN_GENERIC_ERROR, rd // temporary
	}
	rd.RepliesTree = append(rd.RepliesTree, reply)
	return RETURN_GOOD, rd

}

/*
 * Address()
 * Lookup AAAA and A records for given qname (in parallel)
 */

func Address(ctx *Context, qname string, ext Extension) (ReturnType, *ResponseDict) {

	var reply, reply1 *Reply
	var err error
	var rc ReturnType = RETURN_GOOD

	qtypeList := [...]string{"AAAA", "A"}
	replyChannel := make(chan *Reply)

	for _, qtype := range qtypeList {
		go func(qtype string) {
			reply, err = doQuery(ctx, qname, qtype, "IN", ext)
			replyChannel <- reply
		}(qtype)
	}

	rd := new(ResponseDict)
	for i := 0; i < len(qtypeList); i++ {
		reply1 = <-replyChannel
		rd.RepliesTree = append(rd.RepliesTree, reply1)
		if reply1.Err != nil {
			rc = RETURN_GENERIC_ERROR
		}
	}

	return rc, rd
}

/*
 * Hostname()
 */

func Hostname(ctx *Context, ipaddr string, ext Extension) (ReturnType, *ResponseDict) {

	rd := new(ResponseDict)
	qname, err := dns.ReverseAddr(ipaddr)
	if err != nil {
		fmt.Printf("Invalid IP address: %s\n", ipaddr)
		return RETURN_GENERIC_ERROR, nil
	}

	reply, err := doQuery(ctx, qname, "PTR", "IN", ext)
	if err != nil {
		return RETURN_GENERIC_ERROR, rd // temporary
	}
	rd.RepliesTree = append(rd.RepliesTree, reply)
	return RETURN_GOOD, rd

}

/*
 * Service()
 */

func Service(ctx *Context, qname string, ext Extension) (ReturnType, *ResponseDict) {

	rd := new(ResponseDict)
	reply, err := doQuery(ctx, qname, "SRV", "IN", ext)
	if err != nil {
		return RETURN_GENERIC_ERROR, rd // temporary
	}
	rd.RepliesTree = append(rd.RepliesTree, reply)
	return RETURN_GOOD, rd

}
