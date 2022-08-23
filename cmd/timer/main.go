package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KasimBaybikov/Timer/cmd/timer/config"
)

func main() {
	var configFileName string
	configFileName = "/.config/timer/.conf.timer"
	home := os.Getenv("HOME")
	configFile, err := os.Open(home + configFileName)
	if err != nil {
		log.Fatal(err)
	}
	conf := config.Config(configFile)
	fmt.Printf("%+v", conf)
	//http.ListenAndServe()
}
