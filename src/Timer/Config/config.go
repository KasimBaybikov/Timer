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

func getPort(str string) int {
	args := strings.Split(str, "=")
	port, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	return port
}

func GetConfig(file *os.File) FileConfig {
	var conf FileConfig
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if strings.Contains(str, "port") {
			conf.Port = getPort(str)
		}
	}
	return conf
}
