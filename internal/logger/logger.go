package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const logDirectory = "./log"
const logFile = "listener.log"

// GeneralLogger exported
var Info *log.Logger

// ErrorLogger exported
var Error *log.Logger

func init() {
	absPath, err := filepath.Abs(logDirectory)
	if err != nil {
		os.Mkdir(logDirectory, 0755)
	}

	generalLog, err := os.OpenFile(absPath+logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	Info = log.New(generalLog, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(generalLog, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
