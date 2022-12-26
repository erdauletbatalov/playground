package main

import (
	"fmt"
	"time"

	"github.com/miekg/dns"
)

func main() {
	// Set the DNS server to use for the request
	dnsServer := "8.8.8.8:53"

	// Read the array of IP addresses
	ips := []string{"216.58.211.3", "200.140.59.10", "188.235.48.62"}

	// Create a channel to receive the DNS responses
	responses := make(chan *dns.Msg)

	// Send a DNS PTR request for each IP address
	for _, ip := range ips {
		// Create a DNS PTR record with the IP address
		question := new(dns.Msg)
		question.SetQuestion(ip+".in-addr.arpa.", dns.TypePTR)

		// Send the DNS request in a goroutine
		go func() {
			c := new(dns.Client)
			msg, _, err := c.Exchange(question, dnsServer)
			if err != nil {
				fmt.Printf("Error sending DNS request: %s\n", err)
				return
			}
			responses <- msg
		}()
	}

	// Wait for all responses to arrive
	for i := 0; i < len(ips); i++ {
		select {
		case msg := <-responses:
			// Print the response from the DNS server
			for _, rr := range msg.Answer {
				if rr.Header().Rrtype == dns.TypePTR {
					// This is a DNS PTR record, so print the PTR query and TTL value
					fmt.Printf("PTR query: %s\n", rr.(*dns.PTR).Ptr)
					fmt.Printf("TTL of DNS PTR record: %d\n", rr.Header().Ttl)
				}
			}

		case <-time.After(5 * time.Second):
			// Timeout after 5 seconds
			fmt.Println("Timeout waiting for DNS response")
		}
	}
}
