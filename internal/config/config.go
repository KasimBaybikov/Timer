package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type FileConfig struct {
	Socket string
}

func newFileConfig() FileConfig {
	return FileConfig{Socket: "/etc/timer.socket"}
}

func socket(str string) string {
	args := strings.Split(str, "=")
	socket := args[1]
	return socket
}

func Config() FileConfig {
	var configFileName string
	configFileName = "/.config/timer/conf.timer"
	home := os.Getenv("HOME")
	configFile, err := os.Open(home + configFileName)
	if err != nil {
		log.Fatal(err)
	}
	conf := newFileConfig()
	fileScanner := bufio.NewScanner(configFile)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if strings.Contains(str, "socket") {
			conf.Socket = socket(str)
		}
	}
	return conf
}
