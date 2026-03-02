package config

import (
	"io"
	"log"
	"os"
)

func Logging() *log.Logger {
	var file *os.File
	var err error

	file, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	var multiWriter io.Writer = io.MultiWriter(os.Stdout, file)

	return log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
}
