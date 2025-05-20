package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"ridebooking/internal/utils"
	"strings"

	"github.com/golang-jwt/jwt"
)

type key string

const UserContextKey key = "email"

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func JwtMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		var claims utils.MyClaims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserContextKey, claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
