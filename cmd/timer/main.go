package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KasimBaybikov/Timer/internal/config"
)

func pause(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pause_timer"))
	fmt.Print("pause_timer")
}

func start(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start_timer"))
	fmt.Print("start_timer")
}

func main() {
	var configFileName string
	configFileName = "/.config/timer/conf.timer"
	home := os.Getenv("HOME")
	configFile, err := os.Open(home + configFileName)
	if err != nil {
		log.Fatal(err)
	}
	conf := config.Config(configFile)

	os.Remove(conf.Socket)

	mux := http.NewServeMux()
	mux.HandleFunc("/start_timer", start)
	mux.HandleFunc("/pause_timer", pause)

	server := http.Server{
		Handler: mux,
	}
	unixListener, err := net.Listen("unix", conf.Socket)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(unixListener)
}
