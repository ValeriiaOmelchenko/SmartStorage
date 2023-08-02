package main

import (
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Active",
		Message: "Back end is working",
		Version: "0.0.0",
	}
	log.Println("someone just hitted this endpoint")
	_ = app.writeJSON(w, http.StatusOK, payload)
}
