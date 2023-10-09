package main

import "github.com/lmoscheni/go-htmx/server"

func main() {
	server := server.NewServer(":8080")
	server.Start()
}
