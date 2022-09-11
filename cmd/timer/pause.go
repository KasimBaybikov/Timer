package main

import (
	"net/http"
	"time"
)

func (t *Timer) Pause(w http.ResponseWriter, r *http.Request) {
	if t.IsRunning() {
		t.Stop()
		stop := time.Now()
		elasped := stop.Sub(t.start)
		w.Write([]byte(elasped.String()))
	} else {
		w.Write([]byte("The timer has not started yet\n"))
	}
}
