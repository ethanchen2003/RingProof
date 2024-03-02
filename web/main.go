package main

import (
	"log"
	"net/http"

	"github.com/phoneapp/security"
	"github.com/phoneapp/service"
)

func main() {
	mux := http.NewServeMux()

	loginPageHandler := http.HandlerFunc(security.LoginPage)
	signinHandler := http.HandlerFunc(security.Login)
	welcomeHandler := http.HandlerFunc(security.Welcome)
	logoutHandler := http.HandlerFunc(security.Logout)

	spamPhoneHandler := http.HandlerFunc(service.SubmitSpamPhone)

	mux.Handle("/", loginPageHandler)
	mux.Handle("/signin", signinHandler)
	mux.Handle("/welcome", welcomeHandler)
	mux.Handle("/logout", logoutHandler)
	mux.Handle("/spamPhone", spamPhoneHandler)

	log.Print("Listening on :8080...")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)

}
