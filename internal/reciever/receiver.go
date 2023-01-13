package reciever

import (
	"fmt"
	"github.com/sfawcett123/GoConnect/internal/listener"
	log "github.com/sfawcett123/GoConnect/internal/logger"
)

func Start(channel chan listener.Connection) {

	log.Info.Println("Reciever Starting")

	for {
		select {
		case msg := <-channel:
			fmt.Println("New Port ", msg.Port)
		default:
		}
	}
}
