package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KasimBaybikov/Timer/internal/config"
)

func pause(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pause_timer"))
	//need logic
}

func start(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start_timer"))
	//need logic
}

func main() {
	conf := config.Config()
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
	println("Timer is running")
	server.Serve(unixListener)
}
