package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func setUpGoogleSheetsAPI() *sheets.Service {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// get fname of sheetsAPI credentials service account
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", os.Getenv("GCP_SHEETS_CREDENTIAL_PATH"))

	// Retrieve the value to verify it's set
	// value := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	// fmt.Println(value)

	ctx := context.Background()

	// Use service account or OAuth2 credentials
	client, err := google.DefaultClient(ctx, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
	}

	return sheetsService
}

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func setupGoogleOAuth() *oauth2.Config {
	// load .env file from given path
	// we keep it empty it will load .env from current directory

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Replace with your own Google OAuth2 credentials
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GCP_OAUTH_CLIENT"),
		ClientSecret: os.Getenv("GCP_OAUTH_SECRET"),
		RedirectURL:  os.Getenv("GCP_OAUTH_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return oauthConfig
}
