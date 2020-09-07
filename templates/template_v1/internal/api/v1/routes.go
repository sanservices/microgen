package v1

import (
	"context"
	"html/template"
	"net/http"

	_ "goproposal/files/swaggerui/v1/statik" // statik file

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/san-services/apilogger"
)

type DocsModel struct {
	ServicePrefix string
}

func (h Handler) getDocs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lg := apilogger.New(ctx, "")

	tmpl, err := template.New("index.html").ParseFiles("files/swaggerui/v1/index.html")
	if err != nil {
		lg.Info(apilogger.LogCatStartUp, err)
		http.Error(w, "Error: couldn't find template to serve",
			http.StatusInternalServerError)
		return
	}

	data := DocsModel{
		ServicePrefix: h.conf.Service.PathPrefix}

	err = tmpl.Execute(w, data)
	if err != nil {
		lg.Info(apilogger.LogCatStartUp, err)
	}
}

func (h Handler) InitRoutes(ctx context.Context, r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	statikFS, err := fs.New()
	if err == nil {
		r.StrictSlash(true).PathPrefix("/v1/docs/").Handler(
			http.StripPrefix("/v1/docs/", http.FileServer(statikFS)),
		)
	}

	v1.HandleFunc("/docs", h.getDocs).Methods(http.MethodGet)

	// person endpoint example for Get method
	// swagger:route GET /persons persons
	//
	// Get User info
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  400:
	//  200: personsResponse
	v1.HandleFunc("/persons", h.persons).Methods(http.MethodGet)

	// person endpoint example for Get method
	// swagger:route GET /persons/:id person
	//
	// Get User info
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  400:
	//  200: personResponse
	v1.HandleFunc("/persons/{id}", h.person).Methods(http.MethodGet)

	// addPerson endpoint example for POST method
	// swagger:route POST /persons addPerson
	//
	// Post User info
	//
	// Produces:
	//  - application/json
	//
	// Security:
	//  - api_key:
	//
	// Responses:
	//  400:
	//  201:
	v1.HandleFunc("/persons", h.addPerson).Methods(http.MethodPost)
}
