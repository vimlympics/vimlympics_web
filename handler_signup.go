package main

import (
	"database/sql"
	"net/http"

	"github.com/dghubble/gologin/v2/github"
	"github.com/vimlympics/vimlympics_web/db"
	"github.com/vimlympics/vimlympics_web/model"
	"github.com/vimlympics/vimlympics_web/templ"

	"github.com/google/uuid"
)

type LoginHandler struct{}

func (h *application) isLoggedIn(r *http.Request) bool {
	githubUser := h.session.GetString(r.Context(), "githubUser")
	return githubUser != ""
}

func (h *application) Profile(w http.ResponseWriter, r *http.Request) {
	githubUser := h.session.GetString(r.Context(), "githubUser")

	if githubUser == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	profileData, _ := h.query.GetUserProfileData(r.Context(), githubUser)
	var userCountry string

	if !profileData.Country.Valid {
		userCountry = "Unknown"
	} else {
		userCountry = profileData.Country.String
	}

	profile := templ.Profile(githubUser, userCountry, profileData.ApiKey)

	err := templ.Layout(profile, "Login", h.isLoggedIn(r)).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *application) LogOut(w http.ResponseWriter, r *http.Request) {
	h.session.Remove(r.Context(), "githubUser")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *application) issueSession() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userExists, _ := h.query.GetUser(r.Context(), *githubUser.Login)
		if userExists == 0 {
			key := uuid.New().String()
			userParams := db.CreateUserParams{
				Username: *githubUser.Login,
				ApiKey:   key,
			}

			_, err = h.query.CreateUser(r.Context(), userParams)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		h.session.Put(r.Context(), "githubUser", githubUser.Login)

		http.Redirect(w, r, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func (h *application) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	country := r.FormValue("country")
	githubUser := h.session.GetString(r.Context(), "githubUser")

	if _, validCountry := model.ISO3166[country]; !validCountry {
		http.Error(w, "invalid country code", http.StatusBadRequest)
		return
	}

	countryUpdate := db.UpdateUserCountryParams{
		Username: githubUser,
		Country:  sql.NullString{String: country, Valid: true},
	}

	_, err := h.query.UpdateUserCountry(r.Context(), countryUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedSuccess := templ.UpdateSuccess(country)
	err = updatedSuccess.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
