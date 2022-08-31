package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf handles csrfToken
func NoSurf(next http.Handler) http.Handler {
	csrft := nosurf.New(next)

	csrft.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrft
}

// SessionLoad loads an saves session on request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
