package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/charliej2005/chripy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	resp, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, entry := range resp {
		chirp := Chirp{
			ID:        entry.ID,
			CreatedAt: entry.CreatedAt,
			UpdatedAt: entry.UpdatedAt,
			Body:      entry.Body,
			UserID:    entry.UserID,
		}
		chirps = append(chirps, chirp)
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")
	fmt.Println(chirpID)
	uuidID, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}
	resp, err := cfg.db.GetChirp(r.Context(), uuidID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't retrieve chirps", err)
		return
	}

	chirp := Chirp{
		ID:        resp.ID,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
		Body:      resp.Body,
		UserID:    resp.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirp)
}

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleaned,
		UserID: params.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func validateChirp(body string) (string, error) {
	const maxChirpLength = 140
	if len(body) > maxChirpLength {
		return "", errors.New("Chirp is too long")
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(body, badWords)
	return cleaned, nil
}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}
