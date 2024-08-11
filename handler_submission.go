package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/vimlympics/vimlympics_web/db"
)

type SubmissionHandler struct{}

func (h *application) Submission(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	type BodyData struct {
		EventType  int64  `json:"eventtype"`
		EventLevel int64  `json:"level"`
		TimeMs     int64  `json:"timems"`
		ApiKey     string `json:"apikey"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := BodyData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user == "" || data.EventType == 0 || data.EventLevel == 0 || data.ApiKey == "" {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	submissions := db.SubmitScoreParams{
		Username:   user,
		EventType:  data.EventType,
		EventLevel: data.EventLevel,
		Timems:     data.TimeMs,
		ApiKey:     data.ApiKey,
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
