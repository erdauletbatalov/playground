package main

import (
	"encoding/json"
	"log"
	"net"
)

func main() {
	const socketPath = "/var/run/dpdk/yerdaulet/dpdk_telemetry.v2"

	addr, err := net.ResolveUnixAddr("unixpacket", socketPath)
	if err != nil {
		log.Panicf("resolve unix addr: %v", err)
	}

	unixConn, err := net.DialUnix("unixpacket", nil, addr)
	if err != nil {
		log.Panicf("err dialing unix: %v", err)
	}
	defer unixConn.Close()

	buf := make([]byte, 1024)

	var msg initialMessage
	n, err := unixConn.Read(buf)
	if err != nil {
		log.Panicf("err reading the initial message from the socket: %v", err)
	}

	if err := json.Unmarshal(buf, &msg); err != nil {
		log.Panicf("unmarshal error: %v", err)
	}
	if msg.MaxOutputLen <= 0 {
		log.Panicf("invalid max output length: %v", msg.MaxOutputLen)
	}
	buf = make([]byte, msg.MaxOutputLen)

	cmd := "/"
	if _, err := unixConn.Write([]byte(cmd)); err != nil {
		log.Panicf("write the command '%s' to the socket: %v", cmd, err)
	}

	n, err = unixConn.Read(buf)
	if err != nil {
		log.Panicf("read from the socket: %v", err)
	}

}

type initialMessage struct {
	Version      string `json:"version,omitempty"`
	PID          int    `json:"pid,omitempty"`
	MaxOutputLen int    `json:"max_output_len,omit"`
}
