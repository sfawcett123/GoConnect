package listener

import (
	"fmt"
	log "github.com/sfawcett123/GoConnect/internal/logger"
)

// Start initialise listener
func Start(server string) {
	log.Info.Println("Starting..")

	port := Request(server)

	fmt.Println("listening on port ", port)
}

// Request gets a port from the Listener
func Request(server string) (port int) {
	populateBody()

	postRequest(server)

	return GetPort()
}
