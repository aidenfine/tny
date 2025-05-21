package urls

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterUrlsRoutes(r *mux.Router, db *sqlx.DB) {
	urlsRouter := r.PathPrefix("/v1/urls").Subrouter()
	urlsRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		CreateShortUrl(w, r, db)
	}).Methods("POST")

}
