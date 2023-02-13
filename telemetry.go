package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

const (
	telemetryVersion = "v2"
	socketName       = "dpdk_telemetry." + telemetryVersion
	defaultPrefix    = "rte"
)

var cmds []string

func readSocket(sock *os.File, bufLen int, echo bool) map[string]interface{} {
	reply := make([]byte, bufLen)
	_, err := sock.Read(reply)
	if err != nil {
		fmt.Println("Error reading from socket:", err)
		sock.Close()
		return nil
	}

	var ret map[string]interface{}
	err = json.Unmarshal(reply, &ret)
	if err != nil {
		fmt.Println("Error decoding reply:", err)
		sock.Close()
		return nil
	}
	if echo {
		fmt.Println(string(reply))
	}
	return ret
}

func getAppName(pid int) string {
	return ""
}

func findSockets(path string) []string {
	return nil
}

func printSocketOptions(prefix string, paths []string) {

}

func getDPDKRuntimeDir(fp string) string {
	return ""
}

func listFp() {

}

func handleSocket(path string) {
	var prompt string
	sock, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return
	}
	if isatty.IsTerminal(os.Stdin.Fd()) {
		prompt = "--> "
		fmt.Println("Connecting to", path)
	}

	if prompt != "" {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(prompt)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "" {
				continue
			}
			_, err := sock.Write([]byte(input))
			if err != nil {
				fmt.Println("Error writing to socket:", err)
				break
			}
			readSocket(sock, 1024, true)
		}
		sock.Close()
	}
}

func main() {
	handleSocket("")
}
