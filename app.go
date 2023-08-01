package main

import (
	"belajar-go/config"
	"belajar-go/server"
)

func main() {
	config.InitConfig()
	server.Serve()
}
