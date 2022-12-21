package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	// Replace with the IP address you want to look up
	ip := "8.8.8.8"

	// Look up the PTR records associated with the IP address
	ptrs, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through the PTR records
	for _, ptr := range ptrs {
		// Perform a DNS TXT query for the PTR record's domain name
		txts, err := net.LookupTXT(ptr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Iterate through the returned text records
		for _, txt := range txts {
			// Split the text record into fields
			fields := strings.Split(txt, " ")

			// Check if the first field is "TTL"
			if fields[0] == "TTL" {
				// The second field is the TTL value
				ttl, err := strconv.Atoi(fields[1])
				if err != nil {
					fmt.Println(err)
					continue
				}

				// Print the TTL value
				fmt.Println(ttl)
				break
			}
		}
	}
}
