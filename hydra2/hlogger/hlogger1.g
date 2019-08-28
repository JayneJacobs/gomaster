package hlogger

import (
	"log"
	"os"
)

type hydraLogger struct {
	*log.Logger
	filename string
}

//Hlogger creates a log using GetInstance with a name string
var Hlogger *hydraLogger

//var once sync.Once

//GetInstance create a singleton instance of the hydra logger
func GetInstance(name string) *hydraLogger {

	Hlogger = createLogger(name + "hydralogger.log")

	return Hlogger
}

//Create a logger instance
func createLogger(fname string) *hydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	return &hydraLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Llongfile),
	}
}
