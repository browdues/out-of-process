package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

const SockAddr = "/tmp/echo.sock"

func Server(c net.Conn) {
	// defer c.Close()

	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	buf := make([]byte, 512)

	for {
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		fmt.Println("Server read:", string(data))

		_, err = c.Write([]byte("run"))
		if err != nil {
			log.Fatal("write err: ", err)
		}
	}
}

func main() {

	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	go func() {
		for {
			cmd := exec.Cmd{
				Path:   "/Users/nbrowdues/Workspace/github.com/BrowduesMan85/out-of-process/executor/executor",
				Stdout: os.Stdout,
				Stdin:  os.Stdin,
			}
			cmd.Start()
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		// Accept new cnnections, dispatching them to the server
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept err:", err)
		}

		go Server(conn)
	}
}
