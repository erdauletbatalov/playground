package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dialer := net.Dialer{Timeout: time.Duration(10 * time.Second)}
			return dialer.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}
	started := time.Now()
	resolver.LookupAddr(context.Background(), os.Args[1])
	fmt.Printf("Query took: %s", time.Since(started))
}
