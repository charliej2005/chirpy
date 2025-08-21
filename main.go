package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/charliej2005/chripy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database\n")
		return
	}
	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"
	apiCfg := apiConfig{
		db:             dbQueries,
		platform:       platform,
		fileserverHits: atomic.Int32{},
	}
	fileSever := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	mux := http.NewServeMux()
	mux.Handle("/app/", fileSever)

	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)

	mux.HandleFunc("GET /api/chirps", apiCfg.handlerRetrieveChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerRetrieveChirp)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerChirpsCreate)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)

	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

type apiConfig struct {
	db             *database.Queries
	platform       string
	fileserverHits atomic.Int32
}
