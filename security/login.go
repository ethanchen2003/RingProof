package security

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Password string
	Username string
}

type session struct {
	username string
	expiry   time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

var sessions = map[string]session{}

var sessionkey = "session_token"

// 30 minutes
var sessionTimeout = 60 * 30

// hard code to store users. Will move to database table
var inMemeoryUsers = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// for GET
func LoginPage(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "templates/login", new(User))
}

func Login(response http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := &User{Password: password, Username: username}

	userSession, ok := process_user(user)

	if !ok {
		response.WriteHeader(http.StatusUnauthorized)
		//show login fail page
		http.Redirect(response, r, "/", http.StatusFound)
		return
	}

	set_session(response, *userSession)
	http.Redirect(response, r, "/welcome", http.StatusFound)
	//renderTemplate(response, "templates/welcome", user)
}

func renderTemplate(response http.ResponseWriter, tmpl string, user *User) {
	t, err := template.ParseFiles(tmpl + ".html")

	if err != nil {
		fmt.Print(err)
		return
	}
	t.Execute(response, user)
}

func process_user(user *User) (username *string, ok bool) {
	//TODO get user from database
	expectedPassword, ok := inMemeoryUsers[user.Username]
	if !ok || expectedPassword != user.Password {
		return nil, false
	}

	//after implementing the database, we can get full user info from database
	return &user.Username, true
}

func set_session(response http.ResponseWriter, username string) {

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1800 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: username,
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(response, &http.Cookie{
		Name:    sessionkey,
		Value:   sessionToken,
		Expires: expiresAt,
	})
}

func getSessionToken(w http.ResponseWriter, r *http.Request) (*string, bool) {
	c, err := r.Cookie(sessionkey)

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return nil, false
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return nil, false
	}

	sessionToken := c.Value

	// We then get the session from our session map
	userSession, exists := sessions[sessionToken]

	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return nil, false
	}
	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return nil, false
	}

	//renew session time
	expiresAt := time.Now().Add(1800 * time.Second)
	// Set the token in the session map, along with the user whom it represents
	sessions[sessionToken] = session{
		username: userSession.username,
		expiry:   expiresAt,
	}

	return &userSession.username, true
}

func Welcome(w http.ResponseWriter, r *http.Request) {

	userSession, ok := getSessionToken(w, r)
	if !ok {
		return
	}

	user := &User{Username: *userSession}
	renderTemplate(w, "templates/welcome", user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(sessionkey)
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// remove the users session from the session map
	delete(sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	http.SetCookie(w, &http.Cookie{
		Name:    sessionkey,
		Value:   "",
		Expires: time.Now(),
	})
	renderTemplate(w, "templates/login", new(User))
}

func securityFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, ok := getSessionToken(w, r)
		if !ok {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("end of filter")
	})
}
