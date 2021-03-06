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

	ipaddr := os.Args[1]

	rc, response := getdns.Hostname(ctx, ipaddr, ext)
	if rc != getdns.RETURN_GOOD {
		fmt.Printf("getdns.HostnameSync() failed: %d\n", rc)
	} else {
		for i, reply := range response.RepliesTree {
			fmt.Println("Response:", i)
			fmt.Println(reply)
		}
	}

	return
}
