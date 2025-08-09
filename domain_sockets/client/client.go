package main

import (
	"fmt"
	"net"
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
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Hello from client!"))
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Received from server: %s\n", buf[:n])
}
