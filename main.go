package main

// See page 254.

// Chat is a server that lets clients chat with each other.

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// tempDirsQ := []string{}
// var rmdirs []func()
// for _, dir := range tempDirsQ {
// 	os.MkdirAll(dir, 0755)
// 	rmdirs = append(rmdirs, func() {
// 		os.RemoveAll(dir) // Примечание: неверно!
// 	})
// }

// select {

// }

// type A struct {
// 	B int
// }

// M := map[int]A{}

// M[10] = A{10}

// M[10].B = 20

// var count = 0
// for count < 10 { // Начало области видимости num
// 	var num = rand.Intn(10) + 1
// 	fmt.Println(num)
// 	count++
// } // Конец области видимости num

// x := "hello"
// for _, x := range x {
// 	x := x + 'А' - 'а'
// 	fmt.Printf("%c", x) // "HELLO" (по букве за итерацию)
// }
// }
