package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/harljos/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON %v", err))
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follow %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerGetFeedFollowsForUser(w http.ResponseWriter, req *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(req.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feed follows %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, req *http.Request, user database.User) {
	params := chi.URLParam(req, "feedFollowID")
	feedFollowID, err := uuid.Parse(params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't parse feed follow id %v", err))
		return
	}

	err = cfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't delete feed follow %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
