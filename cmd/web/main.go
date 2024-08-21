package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sysnasri/booking/pkg/config"
	"github.com/sysnasri/booking/pkg/handlers"
	"github.com/sysnasri/booking/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//var session *scs.SessionManager

	// change this to true if it is production

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template Cache!")
	}

	app.TemplateCache = tc

	app.PortNumber = ":8080"
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// r,err := helper.Devide(100.0,10.0)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(r)

	// handlers.HttpHandlers()
	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: routes(&app),
	}
	fmt.Printf("starting new application on port: %s", app.PortNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	//http.ListenAndServe(app.PortNumber, nil)

}
