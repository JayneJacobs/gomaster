package printerrors

import (
	"fmt"
	"log"
)

// PrintFatalError is an error handler
func PrintFatalError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error happened while processing file", err)
	}
}
