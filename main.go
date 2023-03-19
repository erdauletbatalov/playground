package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	var temp string
	for i, val := range os.Args[1:] {
		if i == len(os.Args[1:])-1 {
			temp += val
			break
		}
		temp += val + " "
	}

	fmt.Printf("%.10fs elapsed, arguments: %v \n", time.Since(start).Seconds(), temp)

	start = time.Now()
	result := strings.Join(os.Args[1:], " ")

	fmt.Printf("%.10fs elapsed, arguments: %v \n", time.Since(start).Seconds(), result)

}

// Results:

// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000080000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000008000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000078000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000008000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000091000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000007000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000081000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000008000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000014000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000007000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000040000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000008000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000013000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000011000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000013000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000008000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000014000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000008000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// erdauletbatalov@DESKTOP-45RK1O4:~/edu/playground$ go run main.go hello pidor idi nahui  ldfslkd lfjs ldf sldfj
// 0.0000013000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
// 0.0000009000s elapsed, arguments: hello pidor idi nahui ldfslkd lfjs ldf sldfj
