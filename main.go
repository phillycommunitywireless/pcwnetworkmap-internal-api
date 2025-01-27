package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: .env file not found. Falling back to environment variables.")
	}

	// Initialize GCP services
	var sheetsService = setUpGoogleSheetsAPI()
	var oauthConfig = setupGoogleOAuth()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", handleRoot)
	e.GET("/get_networkpoints", func(c echo.Context) error { return handleGetNetworkpoints(c, sheetsService) })
	e.GET("/get_level1", func(c echo.Context) error { return handleGetLevel1(c, sheetsService) })
	e.GET("/get_level2", func(c echo.Context) error { return handleGetLevel2(c, sheetsService) })
	e.GET("/get_level3", func(c echo.Context) error { return handleGetLevel3(c, sheetsService) })

	// auth
	e.GET("/login", func(c echo.Context) error { return handleLogin(c, oauthConfig) })
	e.GET("/auth/callback", func(c echo.Context) error { return handleAuthCallback(c, oauthConfig) })

	// route for internal endpoints
	e.GET("/get_internal", func(c echo.Context) error { return handleGetInternal(c, sheetsService) }, echojwt.WithConfig(echojwt.Config{
		SigningKey:  jwtSecret,
		TokenLookup: "header:Authorization:Bearer ,cookie:Authorization",
	}))

	// Start server
	port := os.Getenv("PORT")
	address := ":" + port
	e.Logger.Fatal(e.Start(address))
}

// Handlers

func handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome! Go to /login to authenticate with Google.")
}

func handleLogin(c echo.Context, oauthConfig *oauth2.Config) error {
	url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleAuthCallback(c echo.Context, oauthConfig *oauth2.Config) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "missing code"})
	}

	// Exchange code for token
	token, err := oauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to exchange token"})
	}

	// Get user info
	client := oauthConfig.Client(c.Request().Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch user info"})
	}
	defer resp.Body.Close()

	// Parse user info
	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Failed to decode user info: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to parse user info"})
	}

	// Check authorization
	if userInfo.Email != os.Getenv("USER_1") && userInfo.Email != os.Getenv("USER_2") {
		return c.JSON(http.StatusForbidden, echo.Map{
			"message": "Login Failed - Unauthorized.",
		})
	}

	// Generate JWT
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": userInfo.Email,
		"name":  userInfo.Name,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := jwtToken.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Failed to sign token: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate token"})
	}

	// Set cookie
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 1)
	cookie.Path = "/"
	// needs to be set to 'false' to be visible to browser APIs
	// to make secure requests down the line
	// https://stackoverflow.com/questions/1022112/why-doesnt-document-cookie-show-all-the-cookie-for-the-site
	cookie.HttpOnly = false
	c.SetCookie(cookie)

	// demo - show successful login by returning JWT
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"message": "Login successful",
	// 	"token":   tokenString,
	// })
	// prod - redir to protected URL (map)
	// change this URL to change based off of env
	return c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:4000/")
}

func handleGetInternal(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "internal")
	networkpoints := process_networkpoints_and_internal(sheet_values)
	// networkpoints := process_internal(sheetsService)
	shelled := prep_for_export(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetNetworkpoints(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "networkpoints")
	networkpoints := process_networkpoints_and_internal(sheet_values)
	// networkpoints := process_internal(sheetsService)
	shelled := prep_for_export(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetLevel1(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "level1")
	networkpoints := process_level1(sheet_values)
	// networkpoints := process_internal(sheetsService)
	shelled := prep_for_export_level1(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetLevel2(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "level2")
	networkpoints := process_level2_level3(sheet_values)
	// networkpoints := process_internal(sheetsService)
	shelled := prep_for_export_level2_3(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}

func handleGetLevel3(c echo.Context, sheetsService *sheets.Service) error {
	sheet_values := get_sheet_values(sheetsService, "level3")
	networkpoints := process_level2_level3(sheet_values)
	// networkpoints := process_internal(sheetsService)
	shelled := prep_for_export_level2_3(networkpoints)
	return c.JSON(http.StatusOK, shelled)
}
