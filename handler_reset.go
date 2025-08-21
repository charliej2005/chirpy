package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Reset denied", nil)
		return
	}

	cfg.fileserverHits.Store(0)
	cfg.db.DeleteAllUsers(r.Context())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset data"))
}
