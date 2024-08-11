package main

import (
	"net/http"

	"github.com/vimlympics/vimlympics_web/templ"
)

type LeaderboardHandler struct{}

func (h *application) Leaderboard(w http.ResponseWriter, r *http.Request) {
	indivSummary, err := h.query.GetIndivSummary(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	countrySummary, err := h.query.GetCountrySummary(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	leaderboard := templ.HomeBoards(indivSummary, countrySummary)
	err = templ.Layout(leaderboard, "Leaderboard", h.isLoggedIn(r)).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
