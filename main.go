package main

import "github.com/magrininicolas/gobank/api"

func main() {
	server := api.NewApiServer(":3000")
	server.Run()
}
