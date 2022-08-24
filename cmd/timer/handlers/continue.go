package handlers

import "net/http"

func Continue(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start_timer"))
	//need logic
}
