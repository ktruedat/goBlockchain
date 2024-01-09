package main

import (
	"flag"
)

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	server := NewBlockchainServer(uint16(*port))
	server.Run()
}
