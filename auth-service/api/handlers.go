package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		err := app.ErrorJson(w, err, http.StatusBadRequest)
		if err != nil {
			log.Println(err)
			return
		}
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		err := app.ErrorJson(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		if err != nil {
			log.Println(err)
			return
		}
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		err = app.ErrorJson(w, errors.New("Invalid credentials"), http.StatusBadRequest)
	}

	// Log authentication
	err = app.logRequest("Authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		_ = app.ErrorJson(w, err)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	err = app.WriteJson(w, http.StatusAccepted, payload)
	if err != nil {
		log.Println(err)
		return
	}
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/logger"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
