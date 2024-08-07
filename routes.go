package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	staticFS := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", staticFS))
	mux.HandleFunc("GET /", app.Leaderboard)
	mux.HandleFunc("POST /submit/{user}/{eventtype}/{level}/{timems}", app.Submission)
	mux.HandleFunc("GET /indiv/{user}", app.Indiv)
	mux.HandleFunc("GET /signup", app.SignUp)
	mux.HandleFunc("GET /country/{countrycode}", app.Country)

	return mux
}
