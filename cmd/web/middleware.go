package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})

}

// LogIPAddress Log the IP address into the console
func LogIPAddress(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("unable to parse IP from: ", r.RemoteAddr)

		} else {

			log.Println("client Ip", h)
		}

		next.ServeHTTP(w, r)
	})
}

// NoSrve adds CSRF protection to all post request!
func NoSrve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteDefaultMode,
	})

	return csrfHandler

}

// SessionLoad loads and saves the session on every request!
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)

}
