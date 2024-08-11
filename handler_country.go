package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/vimlympics/vimlympics_web/templ"
)

type CountryHandler struct{}

func (h *application) Country(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countrycode")
	sqlCountryCode := sql.NullString{String: countryCode, Valid: true}

	countryRecords, err := h.query.GetCountryDetails(r.Context(), sqlCountryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	leaderboard := templ.CountryBoard(countryRecords, countryCode)
	err = templ.Layout(leaderboard, fmt.Sprintf("%s's Records", countryCode), h.isLoggedIn(r)).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
