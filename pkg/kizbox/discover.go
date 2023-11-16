package kizbox

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/grandcat/zeroconf"
)

const (
	DefaultService = "_kizboxdev._tcp"
	DefaultDomain  = "local."
	DefaultTimeout = time.Second
)

type Discover struct {
	Service     string          // Service to lookup
	Domain      string          // Lookup domain, default "local"
	Timeout     time.Duration   // Lookup timeout, default 1 second
	Ifaces      []net.Interface // Multicast interface to use
	DisableIPv4 bool            // Whether to disable usage of IPv4 for MDNS operations. Does not affect discovered addresses.
	DisableIPv6 bool            // Whether to disable usage of IPv6 for MDNS operations. Does not affect discovered addresses.
}

func NewDiscover(params ...WithParam) *Discover {
	d := &Discover{
		Service: DefaultService,
		Domain:  DefaultDomain,
		Timeout: DefaultTimeout,
	}

	for _, param := range params {
		param.apply(d)
	}

	return d
}

func (d *Discover) Run(ctx context.Context, ch chan *Server) error {
	t := zeroconf.IPv4AndIPv6

	if d.DisableIPv4 {
		t &^= zeroconf.IPv4
	}
	if d.DisableIPv6 {
		t &^= zeroconf.IPv6
	}

	resolver, err := zeroconf.NewResolver(zeroconf.SelectIfaces(d.Ifaces), zeroconf.SelectIPTraffic(zeroconf.IPType(t)))
	if err != nil {
		return fmt.Errorf("failed to initialize resolver: %w", err)
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			s := &Server{
				Name:    entry.Instance,
				Service: entry.Service,
				Host:    entry.HostName,
				AddrV4:  entry.AddrIPv4,
				AddrV6:  entry.AddrIPv6,
				Port:    entry.Port,
			}

			for _, txt := range entry.Text {
				k, v, _ := strings.Cut(txt, "=")
				switch k {
				case "api_version":
					s.ApiVersion = v
				case "gateway_pin":
					s.GatewayPin = v
				case "fw_version":
					s.FirmwareVersion = v
				}
			}

			ch <- s
		}
	}(entries)

	c, cancel := context.WithTimeout(ctx, d.Timeout)
	defer cancel()

	err = resolver.Browse(ctx, d.Service, d.Domain, entries)
	if err != nil {
		return fmt.Errorf("failed to browse: %w", err)
	}

	<-c.Done()

	return nil
}

type Server struct {
	Name            string
	Service         string
	Host            string
	AddrV4          []net.IP
	AddrV6          []net.IP
	Port            int
	ApiVersion      string
	GatewayPin      string
	FirmwareVersion string
}

type WithParam func(c *Discover)

func (c WithParam) apply(config *Discover) {
	c(config)
}

func WithService(n string) WithParam {
	return func(c *Discover) {
		c.Service = n
	}
}

func WithDomain(d string) WithParam {
	return func(c *Discover) {
		c.Domain = d
	}
}

func WithTimeout(t time.Duration) WithParam {
	return func(c *Discover) {
		c.Timeout = t
	}
}

func WithInterface(ifaces []net.Interface) WithParam {
	return func(c *Discover) {
		c.Ifaces = ifaces
	}
}

func WithDisableIPv4() WithParam {
	return func(c *Discover) {
		c.DisableIPv4 = true
	}
}

func WithDisableIPv6() WithParam {
	return func(c *Discover) {
		c.DisableIPv6 = true
	}
}
