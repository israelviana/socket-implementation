package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write a message: ")
	message, _ := reader.ReadString('\n')

	fmt.Fprintf(conn, message)

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Server response: %s", response)
}
