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
	//fmt.Printf("%+v", conf)
	os.Remove(conf.Socket())
	timer := newTimer()

	mux := http.NewServeMux()
	mux.HandleFunc("/start_timer", timer.Continue)
	mux.HandleFunc("/pause_timer", timer.Pause)

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
