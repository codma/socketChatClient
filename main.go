package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	port := ":8080"
	conn, err := net.Dial("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	ch := make(chan string)
	go handleReciveMsg(conn, ch)
	handleSendMsg(conn)
}

func handleSendMsg(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(conn, "%s", text+"\n")

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting..!")
			conn.Close()
			return
		}
	}
}

func handleReciveMsg(conn net.Conn, ch chan string) {
	for {
		select {
		case message := <-ch:
			fmt.Printf("from another user : %s\n", message)
		default:
			go func(net.Conn, chan string) {
				message, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				ch <- message
			}(conn, ch)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
