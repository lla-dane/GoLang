package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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

var clients = make(map[net.Conn]string)

func main() {
	listener, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Chat server started on", ServerAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go hanadleClient(conn)
	}
}

func hanadleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Client connected from %s\n", conn.RemoteAddr())

	username := readUsername(conn)
	clients[conn] = username

	broadcast(fmt.Sprintf("%s joined the chat", username))

	for {
		message := readMessage(conn)
		if message == "quit" {
			delete(clients, conn)
			broadcast(fmt.Sprintf("%s left the chat", username))
			return
		}
		timestamp := time.Now().Format("15:04:05")
		completedMessage := fmt.Sprintf("[%s] %s: %s", timestamp, username, message)
		broadcast(completedMessage)
	}
}

func readUsername(conn net.Conn) string {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error raeding the message: ", err)
		return ""
	}
	return strings.TrimSpace(string(buffer[:n]))
}

func readMessage(conn net.Conn) string {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reeading message", err)
		return ""
	}
	return string(buffer[:n])
}

func broadcast(message string) {
	for conn := range clients {
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error broadcasting message: ", err)
		}
	}
}
