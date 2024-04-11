package main

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// App contain the reference for all reusable value in the application.
// Handlers are defined as a method of this type. Therefore they can access
// the values inside App.
type App struct {
	config *oauth2.Config
}

func main() {

	// Get the client id and secret from google cloud credential.
	// url: https://console.cloud.google.com/apis/credentials
	// access them from environment variable
	clientid := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SEC")

	// Oauth config manage the OAuth flow. You have to register
	// the redirect url in the OAuth provider. For the endpoint,
	// there are many provider specific package inside the
	// golang.org/x/oauth2 package
	conf := &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8000/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	// Instantiating the App type
	app := App{config: conf}

	// Creating a new ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("GET /auth/login", app.loginHandler)

	mux.HandleFunc("GET /auth/oauth", app.oAuthHandler)

	mux.HandleFunc("GET /auth/callback", app.oAuthCallbackHandler)

	http.ListenAndServe(":8000", mux)
}
