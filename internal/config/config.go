package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	configFileName        = "/.config/timer/conf.timer"
	defaultSocketName     = "/timer.socket"
	defaulHistoryFileName = "/timer.log"
	emptyLine             = ""
)

type FileConfig struct {
	socketPath      string
	historyFilePath string
}

func newFileConfig() FileConfig {
	home := os.Getenv("HOME")
	return FileConfig{
		socketPath:      home + defaultSocketName,
		historyFilePath: home + defaulHistoryFileName,
	}
}

func socket(str string) string { // В случае изменения не придется переписывать код, надо только подправить логику
	return str
}

func history(filePath string) string {
	history := filePath
	return history
}

func (f FileConfig) Socket() string {
	return f.socketPath
}

func (f FileConfig) History() string {
	return f.historyFilePath
}

func Config() FileConfig {
	home := os.Getenv("HOME")
	conf := newFileConfig()

	configFile, err := os.Open(home + configFileName)
	if err != nil {
		return conf
	}

	fileScanner := bufio.NewScanner(configFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == emptyLine {
			continue
		}
		args := strings.Split(line, "=")
		if len(args) != 2 {
			log.Fatal("Invalid argument: " + line)
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
