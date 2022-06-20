package main

import (
	"os"

	"github.com/damocles217/user_service/src/user"
)

func main() {
	server := user.CreateServer()

	port := os.Getenv("PORT")
	url := os.Getenv("URL")

	server.Run(url + ":" + port)
}
