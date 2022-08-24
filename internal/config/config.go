package config

import (
	"bufio"
	"os"
	"strings"
)

type FileConfig struct {
	Socket string
}

func newFileConfig() FileConfig {
	return FileConfig{Socket: "/etc/timer.socket"}
}

func Socket(str string) string {
	args := strings.Split(str, "=")
	socket := args[1]
	return socket
}

func Config(file *os.File) FileConfig {
	conf := newFileConfig()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if strings.Contains(str, "socket") {
			conf.Socket = Socket(str)
		}
	}
	return conf
}
