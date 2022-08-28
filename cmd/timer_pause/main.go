package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KasimBaybikov/Timer/internal/config"
)

func main() {
	conf := config.Config()

	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", conf.Socket())
			},
		},
	}

	response, err := httpc.Get("http://unix" + "/pause_timer")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, response.Body)
}
