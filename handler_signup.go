package main

import (
	"net/http"

	"github.com/vimlympics/vimlympics_web/templ"
)

type SignupHandler struct{}

func (h *application) SignUp(w http.ResponseWriter, r *http.Request) {
	signUp := templ.SignUp()
	err := templ.Layout(signUp, "SignUp", true).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
