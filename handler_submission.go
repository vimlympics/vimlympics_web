package main

import (
	"net/http"
	"strconv"

	"github.com/vimlympics/vimlympics_web/db"
)

type SubmissionHandler struct{}

func (h *application) Submission(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	eventType := r.PathValue("eventtype")
	eventLevel := r.PathValue("level")
	timeMs := r.PathValue("timems")
	if user == "" || eventType == "" || eventLevel == "" || timeMs == "" {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	eventType64, err := strconv.ParseInt(eventType, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	eventLevel64, err := strconv.ParseInt(eventLevel, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	timeMs64, err := strconv.ParseInt(timeMs, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	submissions := db.SubmitScoreParams{
		Username:   user,
		EventType:  eventType64,
		EventLevel: eventLevel64,
		Timems:     timeMs64,
	}

	res, err := h.query.SubmitScore(r.Context(), submissions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if row == 0 {
		http.Error(w, "No rows affected", http.StatusInternalServerError)
	}
	_, err = w.Write([]byte("Submission received"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
