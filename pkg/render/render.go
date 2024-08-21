package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sysnasri/booking/pkg/config"
	"github.com/sysnasri/booking/pkg/models"
)

var app *config.AppConfig

// NewTemplates set the config for template packages
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td

}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()

	}

	// get the template cache from App Config

	// create a template cache

	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ie := config.AppConfig{}
	// tc := ie.TemplateCache

	// get requested template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template Cache!")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// pt, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl.html")
	// err := pt.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing in template!", err)
	// 	return
	// }

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl.html")
	if err != nil {

		return myCache, err

	}

	// range through all files ending with *.page.tmpl.html

	for _, page := range pages {

		fileName := filepath.Base(page)

		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {

			return myCache, err

		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")
		if err != nil {

			return myCache, err

		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl.html")
			if err != nil {

				return myCache, err

			}
		}
		myCache[fileName] = ts

	}
	return myCache, nil

}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {

// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache!

// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("Creating template and adding to cache")

// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 	} else {
// 		// we have the template in the cache!

// 		log.Println("Using cached template!")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error {

// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl
// 	return nil

// }
