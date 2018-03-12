package socks5

import (
	"net"

	"context"
)

// NameResolver is used to implement custom name resolution
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
}

// DNSResolver uses the system DNS to resolve host names
type DNSResolver struct{}

// Resolve implements NameResolver interface
func (d DNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		return ctx, nil, err
	}
	return ctx, addr.IP, err
}

// RewriteResolver rewrites host names
type RewriteResolver map[string]string

var dnsResolver = DNSResolver{}

// Resolve implements NameResolver interface
func (r RewriteResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	if r != nil {
		if rewrite, ok := r[name]; ok {
			name = rewrite
		}
	}
	return dnsResolver.Resolve(ctx, name)
}
