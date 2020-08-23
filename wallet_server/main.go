package main

import "flag"

func main() {
	port := flag.Uint("port", 8080, "TCP port number for Wallet server")
	domain := flag.String("domain", "0.0.0.0", "Domain to publish app")
	rootPath := flag.String("root", "/", "Root location of app")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain gateway")
	flag.Parse()
	app := NewWalletServer(*domain, uint16(*port), *rootPath, *gateway)
	app.Run()
}
