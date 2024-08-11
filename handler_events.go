package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/vimlympics/vimlympics_web/db"
	"github.com/vimlympics/vimlympics_web/model"
	"github.com/vimlympics/vimlympics_web/templ"
)

type EventHandler struct{}

func (h *application) Events(w http.ResponseWriter, r *http.Request) {
	eventType := r.PathValue("eventtype")
	eventLevel := r.PathValue("level")

	eventType64, err := strconv.ParseInt(eventType, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	eventLevel64, err := strconv.ParseInt(eventLevel, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	eventParams := db.GetEventDetailsParams{
		EventType:  eventType64,
		EventLevel: eventLevel64,
	}

	eventRecord, err := h.query.GetEventDetails(r.Context(), eventParams)
	if len(eventRecord) == 0 {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	eventName := model.EventType(eventType64)
	leaderboard := templ.EventBoards(eventRecord, eventName)
	err = templ.Layout(leaderboard, fmt.Sprintf("%s %s Records", eventName.String(), eventLevel), h.isLoggedIn(r)).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
