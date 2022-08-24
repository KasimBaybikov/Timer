package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KasimBaybikov/Timer/cmd/timer/handlers"
	"github.com/KasimBaybikov/Timer/internal/config"
)

func main() {
	conf := config.Config()
	os.Remove(conf.Socket)

	mux := http.NewServeMux()
	mux.HandleFunc("/start_timer", handlers.Continue)
	mux.HandleFunc("/pause_timer", handlers.Pause)

	server := http.Server{
		Handler: mux,
	}
	unixListener, err := net.Listen("unix", conf.Socket)
	if err != nil {
		log.Fatal(err)
	}
	println("Timer is running")
	server.Serve(unixListener)
}
