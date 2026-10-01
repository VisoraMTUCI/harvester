package handler

import "net/http"

func GetEventHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
}
