package main

import (
	"fmt"
	"github.com/shuque/getdns"
	"os"
)

func main() {

	var ext = make(getdns.Extension)
	ext["dnssec_return_status"] = true

	ctx, err := getdns.ContextCreate()
	if err != nil {
		fmt.Println("ContextCreate() failed", err)
		return
	}
	ctx.SetServer("8.8.8.8")

	qname, qtype := os.Args[1], os.Args[2]

	rc, response := getdns.General(ctx, qname, qtype, ext)
	if rc != getdns.RETURN_GOOD {
		fmt.Printf("getdns.GeneralSync() failed: %d\n", rc)
	} else {
		for i, reply := range response.RepliesTree {
			fmt.Println("Response:", i)
			fmt.Println(reply)
		}
	}

	return
}
