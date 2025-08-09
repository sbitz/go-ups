package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

)

func main() {
	socketPathRoot := ""
	if runtime.GOOS == "windows" {
		socketPathRoot = "C:/"
	} else {
		socketPathRoot = "/"
	}

	socketPath := socketPathRoot + "tmp/my_uds_socket.sock"
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Received from client: %s\n", buf[:n])
	conn.Write([]byte("Hello from server!"))
}
