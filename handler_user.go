package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/harljos/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	user, err := cfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (cfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, req *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, req *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsByUser(req.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get posts %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
