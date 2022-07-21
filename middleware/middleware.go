package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"
	"todo/helper"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("sessionId")
		sessionDetails, idErr := helper.FetchSession(sessionId)
		if idErr != nil && idErr == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if idErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if sessionDetails.ExpiryTime.Before(time.Now()) {
			delErr := helper.DeleteSession(sessionId)
			if delErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(delErr)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId := sessionDetails.UserId
		ctx := r.Context()
		key := os.Getenv("userID")
		ctx = context.WithValue(ctx, key, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
