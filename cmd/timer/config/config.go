package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type FileConfig struct {
	Port int
}

func newFileConfig() FileConfig {
	return FileConfig{Port: 8080}
}

func Port(str string) int {
	args := strings.Split(str, "=")
	port, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	return port
}

func Config(file *os.File) FileConfig {
	conf := newFileConfig()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if strings.Contains(str, "port") {
			conf.Port = Port(str)
		}
	}
	return conf
}
