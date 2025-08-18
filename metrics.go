package main

import (
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	hits := strconv.Itoa(int(cfg.fileserverHits.Load()))
	hitString := "Hits: " + hits
	w.Write([]byte(hitString))
}
