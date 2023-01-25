package listener

import (
	"fmt"
	log "github.com/sfawcett123/GoConnect/internal/logger"
)

type Connection struct {
	Port      int
	Ip        string
	Connected bool
}

// Start initialise listener
func Start(server string, channel chan Connection) {
	log.Info.Println("Starting..")

	port := Request(server)

	channel <- Connection{port, "0.0.0.0", true}

	fmt.Println("Listening on port: ", port)
}

// Request gets a port from the Listener
func Request(server string) (port int) {

	populateBody()

	postRequest(server)

	return GetPort()
}
