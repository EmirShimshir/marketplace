package logrus

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Level string
}

const (
	defaultPerm = 0777
	logsDir     = "./logs"
	name        = "logs.txt"
)

func InitLogrus(cfg *Config) {
	log.SetFormatter(&log.JSONFormatter{})

	err := os.MkdirAll(logsDir, defaultPerm)
	if err != nil {
		fmt.Println("Failed to create log dir")
		panic(err)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/%s", logsDir, name), os.O_APPEND|os.O_CREATE|os.O_RDWR, defaultPerm)
	if err != nil {
		fmt.Println("Failed to create log file")
		panic(err)
	}
	log.SetOutput(f)

	switch cfg.Level {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

}
