package handlers

import "net/http"

func Pause(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pause_timer"))
	//need logic
}
