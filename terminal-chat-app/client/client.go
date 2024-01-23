package main

import (
	"fmt"
	"net"
	"time"
)

const (
	ServerAddress = "localhost:5000"
)

type Message struct {
	Sender    string
	Content   string
	Timestamp string
}

func main() {
	username := getUsername()

	conn, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}

	defer conn.Close()

	go readServerMessages(conn)
	sendMessages(conn, username)
}

func getUsername() string {
	fmt.Print("Enter your username: ")
	var username string
	fmt.Scanln(&username)
	return username
}

func readServerMessages(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading server message:", err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}
}

func sendMessages(conn net.Conn, username string) {
	for {
		var message string
		fmt.Scanln(&message)

		if message == "quit" {
			conn.Write([]byte(message))
			time.Sleep(500 * time.Millisecond) // Allow time for the message to be sent
			return
		}

		timestamp := time.Now().Format("15:04:05")
		message = fmt.Sprintf("[%s] %s: %s\n", timestamp, username, message)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
