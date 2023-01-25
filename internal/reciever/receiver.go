package reciever

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/sfawcett123/GoConnect/internal/listener"
	log "github.com/sfawcett123/GoConnect/internal/logger"
)

const (
	TYPE = "tcp"
	HOST = "localhost"
)

// Start listening on TCP socket
func Start(channel chan listener.Connection) {

	log.Info.Println("Reciever Starting")
	var conn net.Conn
	var port int = 9000

	select {
	case msg := <-channel:
		port = msg.Port
	default:
	}

	// Connect to Client
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+strconv.Itoa(port))
	if err != nil {
		log.Error.Println(err)
		log.Error.Fatalln("Error Resolving TCP Address ")
	}
	conn, err = net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Error.Println(err)
		log.Error.Fatalln("Error connecting to ", port)
	}
	defer conn.Close()

	// Send some JSON data
	json_data, _ := json.Marshal(map[string]string{
		"AIRCRAFT_HEADING": "1101",
		"BUTTON_X":         "OFF"})

	fmt.Println("Message Sent:", string(json_data))
	_, _ = conn.Write([]byte(string(json_data)))

	json_data2, _ := json.Marshal(map[string]string{
		"AIRCRAFT_HEADING": "1302",
		"BUTTON_X":         "OFF"})

	fmt.Println("Message Sent:", string(json_data2))
	_, _ = conn.Write([]byte(string(json_data2)))

	// Read some JSON data
	for {
		message, _ := bufio.NewReader(conn).ReadString('\r')
		fmt.Println("Message Received:", string(message))
	}
}
