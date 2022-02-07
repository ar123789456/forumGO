package middleware

import (
	"forum/models"
	"net/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var user controllers.UserController
		var userSession models.UserSession

		c, err := r.Cookie("session_token")
		if err != nil {
			r.Method = http.MethodGet
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		_, err = userSession.GET(c.Value)
		if err != nil {
			r.Method = http.MethodGet
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		noExpired, err := userSession.Reconciliation()
		if err != nil {
			r.Method = http.MethodGet
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}
		if !noExpired {
			r.Method = http.MethodGet
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
