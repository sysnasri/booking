package handlers

import (
	"log"
	"net/http"

	"github.com/sysnasri/booking/pkg/config"
	"github.com/sysnasri/booking/pkg/models"
	"github.com/sysnasri/booking/pkg/render"
)

// TemplateData holds data sent from handlers to template!

//const PortNumber = ":8080"

// Repository creates a new repository
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {

	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	rIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", rIP)
	log.Println("Ip in home page is", rIP)

	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
	// n, err := fmt.Fprintf(w, "Home Page!")
	// CheckErr(err)
	// log.Println("number of bytes written", n)

}

// About is the about page handlere

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perfom some logic

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello Again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	log.Println("IP in about page is", remoteIP)

	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{StringMap: stringMap})

	// s := AddValues(2, 2)

	// n, err := fmt.Fprintf(w, fmt.Sprintf("This is about page and sum is %d", s))
	// CheckErr(err)
	// log.Println("number of bytes written", n)

}

// Reservation renders the make-reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl.html", &models.TemplateData{})

}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl.html", &models.TemplateData{})

}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl.html", &models.TemplateData{})

}

// Availability renders the search availability
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl.html", &models.TemplateData{})

}

// Contact renders the search availability
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl.html", &models.TemplateData{})

}

// Addvalues returns the sum of values
func AddValues(x, y int) int {
	return x + y
}

// func HttpHandlers() {

// 	http.HandleFunc("/", Repo.Home)
// 	http.HandleFunc("/divide", helper.Devidering)
// 	http.HandleFunc("/about", Repo.About)

// }
