package main

import (
	"net/http"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/github"
)

func (app *application) routes() http.Handler {
	stateConfig := gologin.DebugOnlyCookieConfig
	mux := http.NewServeMux()

	staticFS := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", staticFS))
	mux.HandleFunc("POST /submit/{user}", app.Submission)
	mux.HandleFunc("GET /indiv/{user}", app.Indiv)
	mux.HandleFunc("GET /event/{eventtype}/{level}", app.Events)
	mux.HandleFunc("GET /country/{countrycode}", app.Country)

	mux.HandleFunc("GET /level/{eventtype}/{level}", app.GetLevel)
	mux.HandleFunc("GET /listlevel/", app.ListEvents)

	mux.Handle("GET /login", github.StateHandler(stateConfig, github.LoginHandler(app.oauth, nil)))
	mux.Handle("GET /github/callback", github.StateHandler(stateConfig, github.CallbackHandler(app.oauth, app.issueSession(), nil)))
	mux.HandleFunc("PATCH /profile/updatecountry", app.UpdateCountry)
	mux.HandleFunc("GET /profile", app.Profile)
	mux.HandleFunc("GET /logout", app.LogOut)
	mux.HandleFunc("GET /", app.Leaderboard)
	return mux
}
