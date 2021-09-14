package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	server := createServer()
	conn, err := net.Dial(server.Network, server.Port)

	if err != nil {
		os.Exit(1)
	}

	wg.Add(2)

	fmt.Print("Connected to the server, choose a nickname: ")

	buf := bufio.NewReader(os.Stdin)
	tcpBuf := bufio.NewReader(conn)

	nickname, _ := buf.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	_, _ = conn.Write([]byte(nickname))

	go func() {
		defer wg.Done()
		for {
			msg, _ := buf.ReadString('\n')
			_, _ = conn.Write([]byte(msg))
		}
	}()
	go func() {
		defer wg.Done()
		for {
			msg, _ := tcpBuf.ReadString('\n')
			print(msg)
		}
	}()

	wg.Wait()
}