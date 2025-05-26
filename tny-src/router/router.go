package router

import (
	"context"
	"log"
	"net/http"

	"github.com/aidenfine/tny/tny-src/services/urls"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

func StartRouter(ctx context.Context, db *pgxpool.Pool) error {
	r := mux.NewRouter()

	// setup cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
		}, AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		}, AllowedHeaders: []string{
			"Content-Type", "Authorization",
		},
		AllowCredentials: true,
	})

	registerRoutes(ctx, r, db)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	}).Methods("GET")
	handler := c.Handler(r)
	log.Println("Server running on port 8080...")

	return http.ListenAndServe(":8080", handler)
}
func registerRoutes(ctx context.Context, r *mux.Router, db *pgxpool.Pool) {
	urls.RegisterUrlsRoutes(ctx, r, db)
}
