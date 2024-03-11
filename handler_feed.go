package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/harljos/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	feed, err := cfg.DB.CreateFeed(req.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}
