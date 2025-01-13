package errHandler

import (
	"fmt"
	"log"
	"os"
)

func HandleError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Error:", err)

	logToFile(err)
}

func logToFile(err error) {
	file, fileErr := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Println("Unable to open log file:", fileErr)
		return
	}
	defer file.Close()

	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(err)
}
