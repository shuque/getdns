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

	qname := os.Args[1]

	rc, response := getdns.Address(ctx, qname, ext)
	if rc != getdns.RETURN_GOOD {
		fmt.Printf("getdns.AddressSync() failed: %d\n", rc)
	} else {
		for _, addr := range getdns.JustAddresses(response) {
			fmt.Println(addr)
		}
		// getdns.PrintResponseDict(response)
		// fmt.Printf("JustAddresses: %s\n", getdns.JustAddresses(response))
	}

	return
}
