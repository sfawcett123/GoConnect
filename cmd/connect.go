package main

import (
	"fmt"
	"github.com/sfawcett123/GoConnect/internal/listener"
	log "github.com/sfawcett123/GoConnect/internal/logger"
)

func main() {
	log.Info.Println("Starting..")

	url := "https://localhost:27464/"

	port, err := listener.Request(url)

	fmt.Println(err)
	fmt.Println(port)
}
