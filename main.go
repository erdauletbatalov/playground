package main

// See page 254.

// Chat is a server that lets clients chat with each other.

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	names    = make(map[string]string)
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
	who := conn.RemoteAddr().String()
	ch := make(chan string) // outgoing client messages
	exit := make(chan struct{})

	go clientWriter(conn, ch, exit)
	go clientReader(conn, who, ch)

	// NOTE: ignoring potential errors from input.Err()
	<-exit

	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string, exit chan<- struct{}) {
	for {
		select {
		case <-time.After(10 * time.Second):
			exit <- struct{}{}
			return
		case msg := <-ch:
			fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
		}
	}
}

func clientReader(conn net.Conn, who string, ch chan string) {
	var name string
	conn.Write([]byte("Enter your name: "))
	input := bufio.NewScanner(conn)
	for input.Scan() {
		if _, ok := names[who]; !ok {
			name = input.Text()
			names[who] = name
			ch <- "You are " + name
			conn.Write([]byte(fmt.Sprintf("%s has arrived, %d connections now", name, len(names))))
			messages <- fmt.Sprintf("%s has arrived, %d connections now", name, len(names))
			entering <- ch
			continue
		}
		messages <- name + ": " + input.Text()
	}
	if _, ok := names[who]; ok {
		leaving <- ch
		messages <- name + " has left"
		delete(names, who)
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
