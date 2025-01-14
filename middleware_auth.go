package main

import (
	"net/http"

	"github.com/JenyBo/golearning/internal/auth"
	"github.com/JenyBo/golearning/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 400, "API key is not found")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, "Failed to get user")
			return
		}

		handler(w, r, user)
	}
}
