package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	//서버에 연결
	port := ":8080"
	conn, err := net.Dial("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	//사용자로부터 입력값 받기
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(conn, text+"\n")
		//fmt.Fprintf(conn, "hello \n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("->: " + message)

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting..!")
			return
		}
	}

}
