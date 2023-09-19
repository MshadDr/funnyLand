package middleware

import (
	"context"
	"gitlab.com/M.darvish/funtory/internal/model/repository"
	"gitlab.com/M.darvish/funtory/pkg/jwt"
	"net/http"
	"strings"
)

// ValidateJwtAuthToken and username to request context
func ValidateJwtAuthToken(userRepo repository.IUser, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract the token
		jwtAuthHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(jwtAuthHeader, "Bearer ")

		// Check token and retrieve username in successful case
		username, err := jwt.ValidationAuthToken(token)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := userRepo.GetByUsername(username)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add the token and username to the request context
		ctx := context.WithValue(r.Context(), "username", username)
		ctx = context.WithValue(ctx, "userId", user.ID)
		ctx = context.WithValue(ctx, "phone", user.Phone)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
