package listener

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sfawcett123/GoConnect/internal/logger"
)

var reqBody []byte
var newBoard Board

type Board struct {
	Name string `json:"name,omitempty"`
	IP   string `json:"ip_address,omitempty"`
	Port string `json:"port,omitempty"`
	Os   string `json:"os,omitempty"`
}

func populateBody() {
	var err error
	reqBody, err = json.Marshal(map[string]string{
		"name":            "GoTest",
		"operatingSystem": "Arduino"})
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
}

func GetPort() (port int) {
	port, _ = strconv.Atoi(newBoard.Port)
	return port
}
