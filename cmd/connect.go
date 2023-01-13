package main

import (
	"github.com/sfawcett123/GoConnect/internal/listener"
	"github.com/sfawcett123/GoConnect/internal/reciever"
)

func main() {
	ch := make(chan listener.Connection, 3)

	go listener.Start("https://localhost:27464/", ch)
	go reciever.Start(ch)

	for {
	}
}
