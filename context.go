package getdns

import (
	"time"
)

type Context struct {
	restype      ResType
	timeout      time.Duration
	retries      int
	idle_timeout time.Duration
	edns_version uint8
	port         uint16
	tcp          bool
	adflag       bool
	cdflag       bool
	edns         bool
	dnssec       bool
	bufsize      uint16
	server       string
}

func ContextCreate() (c *Context, err error) {
	c = new(Context)
	c.restype = RESOLUTION_STUB
	c.timeout = time.Second * 3
	c.retries = 3
	c.edns_version = 0
	c.port = 53
	c.tcp = false
	c.server, err = getSysResolver()
	if err != nil {
		return c, err
	}
	return c, err
}

func (c *Context) SetServer(server string) {
	c.server = server
	return
}
