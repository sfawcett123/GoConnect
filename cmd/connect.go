package main

import (
	"github.com/sfawcett123/GoConnect/internal/listener"
	"github.com/sfawcett123/GoConnect/internal/reciever"
	"time"
)

func main() {
	ch := make(chan listener.Connection, 3)

	listener.Start("https://localhost:27464/", ch)

	go reciever.Start(ch)

	for i := 0; i < 100; i++ {
		time.Sleep(5 * time.Second)
	}
}
