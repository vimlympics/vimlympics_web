package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vimlympics/vimlympics_web/levels"
	"github.com/vimlympics/vimlympics_web/model"
)

type LevelHandler struct{}

func (h *application) GetLevel(w http.ResponseWriter, r *http.Request) {
	UrlType := r.PathValue("eventtype")
	UrlLevel := r.PathValue("level")

	IntType, _ := strconv.Atoi(UrlType)
	IntLevel, _ := strconv.Atoi(UrlLevel)

	eventType := model.EventType(IntType)
	eventLevel := model.Level(IntLevel)

	level, err := events.GetEvent(eventLevel, eventType)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.MarshalIndent(level, " ", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *application) ListEvents(w http.ResponseWriter, r *http.Request) {
	events := events.ListEvents()
	jsonBytes, err := json.MarshalIndent(events, " ", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
