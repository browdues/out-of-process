package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/google/uuid"
)

func reader(r net.Conn, eID uuid.UUID) {
	buf := make([]byte, 512)
	defer r.Close()
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		if string(buf[0:n]) == "run" {
			rand.Seed(time.Now().Unix())
			fmt.Printf("Executor %s got run command from server\n", eID)
			duration := rand.Intn(10)
			fmt.Printf("Executor %s running for %d seconds\n", eID, duration)
			time.Sleep(time.Duration(duration) * time.Second)
			fmt.Printf("Executor %s complete\n", eID)
			break
		}
	}
}

func main() {
	executorID := uuid.New()
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go reader(c, executorID)
	for {
		data := fmt.Sprintf("Executor %s alive", executorID)
		_, err = c.Write([]byte(data))
		if err != nil {
			log.Fatal("write error:", err)
		}
		time.Sleep(2 * time.Second)
	}
}
