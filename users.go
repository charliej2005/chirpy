package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	type returnVals struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
	}

	respondWithJSON(w, http.StatusCreated, returnVals{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.UTC(),
		UpdatedAt: user.UpdatedAt.UTC(),
		Email:     user.Email,
	})
}
