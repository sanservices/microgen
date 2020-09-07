package api_version

import (
	"context"
	"html/template"
	"net/http"

	_ "goproposal/files/swaggerui/api_version/statik" // statik file

	"github.com/rakyll/statik/fs"

	"github.com/gorilla/mux"
	"github.com/san-services/apilogger"
)

type DocsModel struct {
	ServicePrefix string
}

func (h Handler) getDocs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lg := apilogger.New(ctx, "")

	tmpl, err := template.New("index.html").ParseFiles("files/swaggerui/api_version/index.html")
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
	api_version := r.PathPrefix("/api_version").Subrouter()

	statikFS, err := fs.NewWithNamespace("api_version")
	if err == nil {
		r.StrictSlash(true).PathPrefix("/api_version/docs/").Handler(
			http.StripPrefix("/api_version/docs/", http.FileServer(statikFS)),
		)
	}

	api_version.HandleFunc("/docs", h.getDocs).Methods(http.MethodGet)

	// ----------------------------------------------------------------
	// Add here your handler routes.
	// ----------------------------------------------------------------
}
