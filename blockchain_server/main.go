package main

import (
	"flag"
)

func main() {
	port := flag.Uint("port", 5000, "TCP port number for Blockchain server")
	domain := flag.String("domain", "0.0.0.0", "Domain to publish app")
	rootPath := flag.String("root", "/", "Root location of app")
	flag.Parse()
	app := NewBlockchainServer(*domain, uint16(*port), *rootPath)
	app.Run()
}
