package main

import (
	"flag"
	"mywsapp/server"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	server := server.NewWebsocketServer(*addr)
	server.Serve()
}
