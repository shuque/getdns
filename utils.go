package getdns

import (
	"fmt"
	"github.com/miekg/dns"
	"strconv"
	"strings"
)

/*
 * getSysResolver() - obtain (1st) system default resolver address
 */
func getSysResolver() (resolver string, err error) {
	config, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err == nil {
		resolver = config.Servers[0]
	} else {
		fmt.Println("Error processing /etc/resolv.conf: " + err.Error())
	}
	return
}

/*
 * addressString() - return address:port string
 */
func addressString(addr string, port uint16) string {
	if strings.Index(addr, ":") == -1 {
		return addr + ":" + strconv.Itoa(int(port))
	} else {
		return "[" + addr + "]" + ":" + strconv.Itoa(int(port))
	}
}
