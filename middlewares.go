package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

/* thank you 161 */
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		// pass the request to the next handler
		next.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"path":           r.RequestURI,
			"execution_time": time.Since(startTime).String(),
			"remote_addr":    r.RemoteAddr,
		}).Info("received a new http request")
	})
}

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "middleware here")
		var cookie *http.Cookie
		var err error

		/* if we are just on homepage or logging in, no need for authentication */
		path := r.URL.Path
		if path == "/" || path == "/login" || path == "/signup" {
			goto end
		}

		cookie, err = r.Cookie("session_token")
		if err != nil {
			goto end
		}
		_ = cookie
		/* scan the database for the session cookie... if doesn't match
		user_id, do a redirect to homepage */

	end:
		next.ServeHTTP(w, r)
		return
	})
}
