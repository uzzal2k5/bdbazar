package utils

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
)

func LogInfo(msg string) {
	infoLogger.Println(msg)
}

func LogError(msg string) {
	errorLogger.Println(msg)
}
