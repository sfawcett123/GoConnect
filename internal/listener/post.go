package listener

import (
	"bytes"
	"encoding/json"
	log "github.com/sfawcett123/GoConnect/internal/logger"
	"io/ioutil"
	"net/http"
)

var reqBody []byte
var resp *http.Response

func populateBody() {
	var err error
	reqBody, err = json.Marshal(map[string]string{
		"name":            "GoTest",
		"operatingSystem": "Arduino"})
	if err != nil {
		log.Error.Println("Populate Body Failure ", err)
		return
	}
}

func postRequest(server string) {
	var err error
	resp, err = http.Post(server+"api/v1/API/register", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error.Println("PostRequest Failure ", err)
		return
	}
	log.Info.Println("PostRequest")
	defer resp.Body.Close()
}

func ReadAll() (body []byte) {
	if resp == nil {
		log.Error.Println("No Response")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
