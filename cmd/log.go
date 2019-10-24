package cmd

import (
	"log"
	"os"
)

var infoLog *log.Logger
var logFile *os.File

func logInit() {
	fileName := "datarw/Agenda.log"
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)

	if err != nil {
		log.Fatalln("Open file error")
	}
	infoLog = log.New(logFile, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Println("Cmd deleteuser called")
}

func logSave(str string, logType string) {
	infoLog.SetPrefix(logType)
	if curUser != nil {
		infoLog.Println("curUser: " + curUser.Name + "  " + str)
	} else {
		infoLog.Println(str)
	}

}