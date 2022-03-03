package main

import (
	"github.com/RomaBiliak/lets-go-chat/internal/server"
)

var ser = server.Server{}

func main() {
	ser.Run(":8080")
}

