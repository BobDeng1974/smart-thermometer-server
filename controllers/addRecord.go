package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unsign3d/smart-thermometer-server/utils"

	"models"
)

func AddRecord(w http.ResponseWriter, r *http.Request) {
	database := utils.Database()
	collection := database.Collection("records")
	record := parseBody(w, r)
	models.SaveRecord(collection, record)
	w.WriteHeader(http.StatusCreated)
}

func parseBody(w http.ResponseWriter, r *http.Request) models.Record {
	recordBody := models.Record{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&recordBody)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, err.Error())
	}
	return recordBody
}
