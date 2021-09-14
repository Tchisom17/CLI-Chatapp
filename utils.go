package main

type Server struct {
	Network string
	Port string
}

func createServer() Server{
	server := Server{ Port: ":5000", Network: "tcp"}
	return server
}
