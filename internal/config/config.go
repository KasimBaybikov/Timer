package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const configFileName = "/.config/timer/conf.timer"

type FileConfig struct {
	socketPath      string
	historyFilePath string
}

func newFileConfig() FileConfig {
	return FileConfig{
		socketPath:      "/tmp/timer.socket",
		historyFilePath: "/tmp/timer.log"}
}

func socket(str string) string { // В случае изменения не придется переписывать код, надо только подправить логику
	return str
}

func history(filePath string) string {
	history := filePath
	return history
}

func (f FileConfig) Socket() string { //Инкапсулирую данные от пользователя
	return f.socketPath
}

func (f FileConfig) History() string {
	return f.historyFilePath
}

func Config() FileConfig {
	home := os.Getenv("HOME")
	configFile, err := os.Open(home + configFileName)
	if err != nil {
		log.Fatal(err)
	}
	conf := newFileConfig()
	fileScanner := bufio.NewScanner(configFile)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		args := strings.Split(str, "=")
		if len(args) != 2 {
			log.Fatal("Invalid argument: " + str)
		}
		switch args[0] {
		case "socket":
			conf.socketPath = socket(args[1])
		case "history":
			conf.historyFilePath = history(args[1])
		default:
			log.Fatal("unknown argument: ", args[0])
		}
	}
	return conf
}
