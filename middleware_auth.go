package main

import (
	"fmt"
	"net/http"

	"github.com/harljos/rss_aggregator/internal/auth"
	"github.com/harljos/rss_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		apiKey, err := auth.GetApiKey(req.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Couldn't find API key %v", err))
			return
		}

		user, err := cfg.DB.GetUserByApiKey(req.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't find user %v", err))
			return
		}

		handler(w, req, user)
	}
}
