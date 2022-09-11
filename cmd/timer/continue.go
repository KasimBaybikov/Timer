package main

import (
	"net/http"
	"time"
)

func (t *Timer) Continue(w http.ResponseWriter, r *http.Request) {
	if !t.IsRunning() {
		t.Start()
		t.start = time.Now()
		w.Write([]byte("start_timer\n"))
	} else {
		w.Write([]byte("The timer is already running\n"))
	}
}
