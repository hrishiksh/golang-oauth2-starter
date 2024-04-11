# Golang OAuth2 Starter

This is a starter repository from where a beginner can take inspiration to add OAuth2 system in their Golang application. This project doesn't use any 3rd party package. For OAuth2 flow, the official [`golang.org/x/oauth2`](https://pkg.go.dev/golang.org/x/oauth2) package is used.

## Why this repository exists

1. When I was learning, I didn't find a good source and code example to impliment OAuth2 in Golang.
2. There are very easy to use package like [Goth](https://github.com/markbates/goth), but I don't want to add an additional 3rd party dependency.
3. Token refresh and building HTTP client using the access token is not clear in the documentation.

## Getting Started

1. Clone this repository

   ```bash
   git clone https://github.com/hrishiksh/golang-oauth2-starter.git
   ```

2. Download the requirements

   ```bash
   go mod download
   go mod tidy
   ```

3. Get Client ID and Client Secret from the Oauth provider. I am using Google for example. Go to the [Google cloud credential page](https://console.cloud.google.com/apis/credentials) for getting client id and secret.

4. Add the Client ID and secret as an environment variable and run the application

   ```bash
   CLIENT_ID="clientid" CLIENT_SEC="secret" go run .
   ```

## Important things to remember

1. Don't add your logo in Google OAuth consent screen. Otherwise you have to varify your app from Google.
