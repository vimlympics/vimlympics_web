package main

import (
	"fmt"
	"net/http"

	"github.com/vimlympics/vimlympics_web/templ"
)

type IndivHandler struct{}

func (h *application) Indiv(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	indivRecord, err := h.query.GetIndivDetails(r.Context(), user)
	if len(indivRecord) == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	leaderboard := templ.IndivBoards(indivRecord, user)
	err = templ.Layout(leaderboard, fmt.Sprintf("%s's Records", user), h.isLoggedIn(r)).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
