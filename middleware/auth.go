package middleware

import (
	"Forum/database"
	"Forum/handlers"
	"context"
	"log"
	"net/http"
	"time"
)

// Log each request's method & path
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Check session + set user role
func SessionMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Retrieve session from cookie
		cookie, err := r.Cookie("session_id")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		//Get session data from database
		session, err := database.GetSessionByID(cookie.Value)
		if err != nil || session.IsDeleted || time.Now().After(session.Expiration) {
			next.ServeHTTP(w, r)
			return
		}

		//Set user info in request context
		ctx := context.WithValue(r.Context(), handlers.UserIDKey, session.UserID)
		ctx = context.WithValue(ctx, handlers.UserRoleKey, session.UserRole)

		//Call next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
