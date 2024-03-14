package main

import (
	"os"

	"github.com/mesxx/Fiber_User_Management_API/servers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return ":" + port
}

func main() {
	server := servers.Server()
	server.Listen(getPort())
}
