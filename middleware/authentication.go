package middleware

import (
	"forum/controllers"
	"net/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user controllers.UserController
		_, err := r.Cookie("session_token")

		if err != nil {
			r.Method = http.MethodGet
			user.LogIn(w, r)
		}
		next.ServeHTTP(w, r)
	})
}
