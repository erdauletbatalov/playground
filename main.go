package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var cmds map[string]bool

func main() {
	args := os.Args[1:]
	socketPath := "/var/run/dpdk/yerdaulet/dpdk_telemetry.v2"
	if len(args) > 0 {
		socketPath = args[0]
	}
	handleSocket(args, socketPath)
}

func readSocket(conn net.Conn, size int, prompt string) (map[string]interface{}, error) {
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
	// if _, ok := response["output"]; ok {
	fmt.Printf("%s\n", response)
	// }
	return response, nil
}

func handleSocket(args []string, socketPath string) {
	var prompt string
	prompt = "--> "
	fmt.Println("Connecting to", socketPath)

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println("Error connecting to", socketPath)
		conn.Close()
		if _, err := os.Stat(socketPath); os.IsNotExist(err) {
			return
		}
		sockets := findSockets(filepath.Dir(socketPath))
		if len(sockets) > 0 {
			fmt.Println("\nOther DPDK telemetry sockets found:")
			// 	printSocketOptions(args, sockets)
			// } else {
			// 	listFp(args)
		}
		return
	}
	defer conn.Close()
	response, err := readSocket(conn, 1024, prompt)
	if err != nil {
		log.Fatal(err)
	}
	outputBufLen := int(response["max_output_len"].(float64))
	appName := int(response["pid"].(float64))
	if appName != 0 && prompt != "" {
		fmt.Printf("Connected to application: \"%s\"\n", appName)
	}
	conn.Write([]byte("/"))
	response, err = readSocket(conn, outputBufLen, "")
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
			if cmds[text] {
				conn.Write([]byte(text))
				response, err = readSocket(conn, outputBufLen, "")
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Invalid command. Type /help for list of commands.")
			}
		}
	}
}

func findSockets(path string) []string {
	var sockets []string
	files, _ := os.ReadDir(path)
	for _, file := range files {
		fmt.Println(path + file.Name())
		if !IsSocket(path + file.Name()) {
			continue
		}
		sockets = append(sockets, fmt.Sprintf("%s/%s", path, file.Name()))
	}
	return sockets
}

// Check if a file is a socket.
func IsSocket(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal("ERROR - ", err)
	}
	return fileInfo.Mode().Type() == fs.ModeSocket
}

type initialMessage struct {
	Version      string `json:"version,omitempty"`
	PID          int    `json:"pid,omitempty"`
	MaxOutputLen int    `json:"max_output_len,omit"`
}
