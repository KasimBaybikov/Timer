package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KasimBaybikov/Timer/internal/config"
)

func main() {
	conf := config.Config()
	os.Remove(conf.Socket())

	_, err := os.Stat(conf.History())
	if err != nil {
		_, err := os.Create(conf.History())
		if err != nil {
			log.Fatal(err)
		}
	}

	timer := newTimer()
	timer.SetHistoryFilePath(conf.History())

	mux := http.NewServeMux()
	mux.HandleFunc("/timer_continue", timer.Continue)
	mux.HandleFunc("/timer_pause", timer.Pause)

	server := http.Server{
		Handler: mux,
	}
	unixListener, err := net.Listen("unix", conf.Socket())
	if err != nil {
		log.Fatal(err)
	}
	println("Timer is running")
	server.Serve(unixListener)
}
