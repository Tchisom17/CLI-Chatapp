package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	 Nickname string
	 Conn net.Conn
}

func log(msg string) {
	start := time.Now()
	fmt.Println(start.String() + " " + msg)
}

func main(){
	var clients []Client

	server := createServer()
	serverConnection, err := net.Listen(server.Network, server.Port)

	if err != nil {
		os.Exit(1)
	}

	log("Server started")

	for {
		conn, acceptErr := serverConnection.Accept()
		if acceptErr != nil {
			os.Exit(1)
		}

		log( "New client has connected from " + conn.RemoteAddr().String())

		go func() {
			//read the nickname and store and store the client nickname and connection
			buf := bufio.NewReader(conn)
			nickname, _ := buf.ReadString('\n')
			currentClient := Client{Nickname: strings.TrimSuffix(nickname, "\n"), Conn: conn}
			clients = append(clients, currentClient)

			for {
				//wait for message to get to the server
				msg, errOnMsg := buf.ReadString('\n')
				if msg == "" {
					continue
				}else if msg == " "{
					continue
				}
				if errOnMsg != nil {
					log("Client " + currentClient.Nickname + " disconnected")
					break
				}
				log("Message from " + currentClient.Nickname + ": " + strings.TrimSuffix(msg, "\n"))

				//broadcast message to all the clients
				for _, c := range clients {
					if c.Nickname != currentClient.Nickname {
						_, _ = c.Conn.Write([]byte(currentClient.Nickname + ": " + msg))
					}
				}
			}
		}()
	}
}
