package main

import (
	"fmt"
	"github.com/shuque/getdns"
	"os"
)

func main() {

	var ext getdns.Extension

	ctx, err := getdns.ContextCreate()
	if err != nil {
		fmt.Println("ContextCreate() failed", err)
		return
	}

	qname, qtype := os.Args[1], os.Args[2]

	rc, response := getdns.General(ctx, qname, qtype, ext)
	if rc != getdns.RETURN_GOOD {
		fmt.Printf("getdns.GeneralSync() failed: %d\n", rc)
	} else {
		getdns.PrintResponseDict(response)
	}

	return
}
