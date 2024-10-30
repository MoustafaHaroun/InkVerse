package middleware

import (
	"context"
	"net/http"

	"github.com/MoustafaHaroun/InkVerse/internal/server/user"
	"github.com/MoustafaHaroun/InkVerse/pkg/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// func Authentication(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		tokenString := r.Header.Get("Authorization")

// 		if tokenString == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString = tokenString[len("Bearer "):]

// 		_, err := auth.VerifyToken(tokenString)
// 		if err != nil {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func WithJWTAuth(handlerFunc http.HandlerFunc, repository user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		token, err := auth.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)

		userID, err := uuid.Parse(str)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := repository.GetByID(userID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
