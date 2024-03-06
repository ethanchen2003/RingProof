package main

import (
	"log"
	"net/http"

	"github.com/phoneapp/security"
	"github.com/phoneapp/service"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", enableCORS(http.HandlerFunc(security.LoginPage)))
	mux.Handle("/signin", enableCORS(http.HandlerFunc(security.Login)))
	mux.Handle("/welcome", enableCORS(http.HandlerFunc(security.Welcome)))
	mux.Handle("/logout", enableCORS(http.HandlerFunc(security.Logout)))
	mux.Handle("/spamPhone", enableCORS(http.HandlerFunc(service.SubmitSpamPhone)))

	log.Print("Listening on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
