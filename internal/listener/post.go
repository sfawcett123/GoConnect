package listener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sfawcett123/GoConnect/internal/logger"
)

var reqBody []byte
var newBoard Board
var postdata PostData

type Board struct {
	Name string `json:"name,omitempty"`
	IP   string `json:"ip_address,omitempty"`
	Port string `json:"port,omitempty"`
	Os   string `json:"os,omitempty"`
}

type PostData struct {
	Name            string
	OperatingSystem string
	Outputs         []string
}

func populateBody() {
	var err error

	postdata = PostData{Name: "AH interface board", OperatingSystem: "Linux", Outputs: []string{"PLANE LATITUDE"}}
	reqBody, err = json.Marshal(postdata)
	if err != nil {
		log.Error.Println("Populate Body Failure %1", err)
		return
	}
}

func postRequest(server string) {
	var err error
	resp, err := http.Post(server+"api/v1/API/register", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error.Println("PostRequest Failure ", err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&newBoard)
	if err != nil {
		log.Error.Println("Decoder Error ", err)
	}

	fmt.Println("Board Information: ", newBoard)
}

func GetPort() (port int) {
	port, _ = strconv.Atoi(newBoard.Port)
	return port
}
