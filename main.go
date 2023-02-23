package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var cmds map[string]bool

func main() {
	args := os.Args[1:]
	socketPath := "/var/run/dpdk/yerdaulet/dpdk_telemetry.v2"
	if len(args) > 0 {
		socketPath = args[0]
	}
	handleSocket(socketPath)
}

func readSocket(conn net.Conn, size int) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	buf := make([]byte, size)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(buf[:n], &response)
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range response {
		fmt.Printf("%s\t%s\n", key, value)
	}
	fmt.Println()
	return response, nil
}

func handleSocket(socketPath string) {
	prompt := "--> "
	fmt.Println("Connecting to", socketPath)

	addr, err := net.ResolveUnixAddr("unixpacket", socketPath)
	if err != nil {
		log.Panicf("resolve unix addr: %v", err)
	}

	conn, err := net.DialUnix("unixpacket", nil, addr)
	if err != nil {
		log.Panicf("err dialing unix: %v", err)
		conn.Close()
		return
	}
	defer conn.Close()
	response, err := readSocket(conn, 1024)
	if err != nil {
		log.Fatal(err)
	}
	outputBufLen := int(response["max_output_len"].(float64))
	appName := int(response["pid"].(float64))
	if appName != 0 && prompt != "" {
		fmt.Printf("Connected to application: \"%d\"\n", appName)
	}
	conn.Write([]byte("/"))
	response, err = readSocket(conn, outputBufLen)
	if err != nil {
		log.Fatal(err)
	}
	cmds = make(map[string]bool)
	for _, cmd := range response["/"].([]interface{}) {
		cmds[cmd.(string)] = true
	}
	for {
		var text string
		fmt.Print(prompt)
		fmt.Scanln(&text)
		text = strings.TrimSpace(text)
		if text == "quit" {
			break
		}
		if len(text) > 0 && text[0] == '/' {
			// if cmds[text] {
			conn.Write([]byte(text))
			_, err = readSocket(conn, outputBufLen)
			if err != nil {
				log.Fatal(err)
			}
			// } else {
			// 	fmt.Println("Invalid command. Type /help for list of commands.")
			// }
		}
	}
}

type initialMessage struct {
	Version      string `json:"version,omitempty"`
	PID          int    `json:"pid,omitempty"`
	MaxOutputLen int    `json:"max_output_len,omit"`
}
