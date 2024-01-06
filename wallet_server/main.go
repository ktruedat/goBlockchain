package main

import "flag"

func main() {
	port := flag.Uint("port", 8080, "TCP port number for Wallet Server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain Gateway")
	flag.Parse()

	server := NewWalletServer(uint16(*port), *gateway)
	server.Run()
}
