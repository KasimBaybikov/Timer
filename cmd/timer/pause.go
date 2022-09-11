package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func (t *Timer) Pause(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile(t.historyFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file = os.Stdout
	}
	if t.IsRunning() {
		t.Stop()
		now := time.Now()
		stop := time.Now().Add(time.Minute * 40).Add(time.Hour * 10)
		elasped := stop.Sub(t.start)
		elasped = elasped.Round(time.Second)
		fmt.Print(file.Name())
		fmt.Fprintln(file, now.Format("1 January 2006"), elasped.String())
		w.Write([]byte("You worked: " + elasped.String()))
	} else {
		w.Write([]byte("The timer has not started yet\n"))
	}
}
