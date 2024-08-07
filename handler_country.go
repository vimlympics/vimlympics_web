package main

import (
	"fmt"
	"net/http"

	"github.com/vimlympics/vimlympics_web/templ"
)

type CountryHandler struct{}

func (h *application) Country(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countrycode")
	countryRecords, err := h.query.GetCountryDetails(r.Context(), countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	leaderboard := templ.CountryBoard(countryRecords, countryCode)
	err = templ.Layout(leaderboard, fmt.Sprintf("%s's Records", countryCode), false).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
