package main

import (
	"github.com/Noah-Wilderom/go-logger-service/data"
	"net/http"
)

type JsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JsonPayload
	_ = app.readJson(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.ErrorJson(w, err)
		return
	}

	response := JsonResponse{
		Error:   false,
		Message: "Logged",
	}

	_ = app.WriteJson(w, http.StatusAccepted, response)
}
