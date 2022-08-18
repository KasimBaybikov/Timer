package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KasimBaybikov/Timer/src/timer/config"
)

func main() {
	args := os.Args[1:]
	var configFileName string
	if len(args) == 0 {
		configFileName = "/.conf.timer"
	} else {
		configFileName = args[0]
	}
	home := os.Getenv("HOME")
	fmt.Println(home + configFileName)
	configFile, err := os.Open(home + configFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(configFile)
	conf := config.GetConfig(configFile)
	fmt.Printf("%+v", conf)
	//http.ListenAndServe()
}
