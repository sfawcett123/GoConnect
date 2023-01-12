package listener

import (
	"fmt"
)

// Request gets a port from the Listener
func Request(server string) (port int, err error) {
	port = -1

	populateBody()

	postRequest(server)

	body := ReadAll()

	fmt.Println(string(body))

	return
}
