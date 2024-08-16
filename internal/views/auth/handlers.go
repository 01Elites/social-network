package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/session"

	"github.com/gofrs/uuid"
	"golang.org/x/oauth2"
)

type SignUpRequst struct {
	UserName       string    `json:"user_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	NickName       string    `json:"nick_name"`
	ProfilePrivacy string    `json:"profile_privacy"`
	Image          string    `json:"image"`
	About          string    `json:"about"`
}

/*
	SignUp handles the HTTP POST request for user sign-up.
	It reads the request body, unmarshals the JSON data, validates the sign-up data,
	SignUpRequst struct defines the structure of the JSON request body expected for user sign-up.

	Body: {
		"user_name": string, // optional, will be generated if not provided
		"email": string,
		"password": string,
		"first_name": string,
		"last_name": string,
		"gender": "male" | "female"
		"date_of_birth": 0, // unix timestamp
		"nick_name": string // optional
		"about": string, // optional
		"profile_privacy": "public" | "private"
		}

hashes the password, adds the user to the database, and returns the appropriate response.
If any error occurs during the process, it returns the corresponding HTTP error status code.
*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	var data SignUpRequst

	// Create a JSON decoder for the request body
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	// Decode the JSON directly from the request body
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("Error decoding sign up request: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}

	if err := ValidatePassword(data.Password); err != nil {
		log.Printf("Password validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.ProfilePrivacy == "" {
		data.ProfilePrivacy = models.ProfilePrivacy.Public
	}

	if data.Image != "" && data.Image != "null" { // Save the image if it exists
		data.Image, err = helpers.SaveBase64Image(data.Image)
		if err != nil {
			fmt.Println("Error with Profile Image:\n", err)
		}
	}

	userProfile := models.UserProfile{
		NickName:       data.NickName,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Gender:         data.Gender,
		DateOfBirth:    data.DateOfBirth,
		ProfilePrivacy: data.ProfilePrivacy,
		Avatar:         data.Image,
		About:          data.About,
	}

	if err := helpers.ValidateUserProfileData(&userProfile); err != nil {
		log.Printf("Validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before adding it
	hash, err := HashPassword(data.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
	data.Password = hash
	data.Email = strings.ToLower(data.Email) // Convert email to lowercase for consistency

	// Generate a unique username
	username, err := database.GenerateUniqueUsername(data.FirstName, data.LastName)
	if err != nil {
		log.Fatalf("Failed to generate unique username: %v", err)
	}
	user := models.User{
		UserName: username,
		Email:    data.Email,
		Password: data.Password,
		Provider: models.Provider.Manual,
	}

	if err := helpers.ValidateUserFields(&user); err != nil {
		log.Printf("Validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.SignUpUser(user, userProfile); err != nil {
		fmt.Println(user, userProfile)
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			if strings.Contains(err.Error(), "user_user_name_key") {
				log.Printf("Username already exists: %v", err)
				helpers.HTTPError(w, "Username already exists", http.StatusConflict)
				return
			} else if strings.Contains(err.Error(), "unique_email_provider") {
				log.Printf("Email already exists: %v", err)
				helpers.HTTPError(w, "Email already exists", http.StatusConflict)
				return
			}
			log.Printf("User already exists: %v", err)
			helpers.HTTPError(w, "User already exists", http.StatusConflict)
			return
		} else if strings.Contains(err.Error(), "SQLSTATE 22P02") {
			log.Printf("Invalid data format: %v", err)
			helpers.HTTPError(w, "Invalid data format", http.StatusBadRequest)
			return
		}
		log.Printf("Error signing up user: %v", err)
		helpers.HTTPError(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	// AddClient(data.UserName)
	// SendUsersList()
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Signup successful")
}

type SignInRequst struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
SignIn handles the login request.
It reads the request body and responds with a success(200) or Unauthorized(401).

SignInRequst Body:

	{
	    "email": string,
	    "password": string
	}
*/
func SignIn(w http.ResponseWriter, r *http.Request) {
	var data SignInRequst

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&data); err != nil {
		log.Printf("Error decoding sign up request: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}

	data.Email = strings.ToLower(data.Email) // Convert email to lowercase for consistency

	user, err := database.GetManualUser(data.Email)
	if err != nil {
		log.Printf("Error getting manual user: %v", err)
		helpers.HTTPError(w, "username or password is incorrect", http.StatusUnauthorized)
		return
	}
	hashMatched := CheckPasswordHash(data.Password, user.Password)
	if !hashMatched {
		log.Printf("Password does not match")
		helpers.HTTPError(w, "username or password is incorrect", http.StatusUnauthorized)
		return
	}
	// Create a seesion for this user
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error creating session UUID: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}

	// Add session to database
	if err := database.AddUserSession(user.UserID, sessionUUID.String()); err != nil {
		log.Printf("Error adding session: %v", err)
		helpers.HTTPError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set a cookie with a session token that can be used to authenticate access without logging in
	// session.SetSessionCookie(w, sessionUUID.String())

	session.SetAutherizationHeader(w, sessionUUID.String())
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Signin successful")
}

// LogOut handles the user logout functionality.
// It retrieves the session from the token , deletes the session from the database,
// and clears the session cookie.
func LogOut(w http.ResponseWriter, r *http.Request) { // Get the session token from the cookie
	token, err := session.ExtractToken(r)
	if err != nil {
		log.Printf("Error extracting token: %v", err)
		helpers.HTTPError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Delete the session from the database
	if err := database.DeleteUserSession(token); err != nil {
		log.Printf("Error deleting session: %v", err)
		helpers.HTTPError(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	// Clear the session cookie
	session.ClearAutherizationHeader(w)
	// AddClient(data.UserName)
	io.WriteString(w, "LogOut success")
}

var oauth2Config = &oauth2.Config{
	ClientID:     "09f2da37-9d1c-414e-bee5-6a58942c8e17",
	ClientSecret: "IbjBSXycBHFCltksPvW2XhrAJJa9VS2MRlR30vQ8ufEI",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://learn.reboot01.com/git/login/oauth/authorize",
		TokenURL: "https://learn.reboot01.com/git/login/oauth/access_token",
	},
	RedirectURL: "http://localhost:8081/api/auth/gitea/callback",
	Scopes:      []string{"read:user"},
}

func GiteaLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
}

func GiteaCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging token: %v", err)
		helpers.HTTPError(w, "Failed to authenticate", http.StatusInternalServerError)
		return
	}

	// Get the user info from Gitea
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://learn.reboot01.com/git/api/v1/user")
	if err != nil {
		log.Printf("Error getting user info: %v", err)
		helpers.HTTPError(w, "Failed to retrieve user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	gitea_user := models.GiteaUser{}

	if err := json.NewDecoder(resp.Body).Decode(&gitea_user); err != nil {
		log.Printf("Error decoding user info: %v", err)
		helpers.HTTPError(w, "Failed to retrieve user info", http.StatusInternalServerError)
		return
	}

	userProfile := models.UserProfile{
		NickName:       gitea_user.Login,
		ProfilePrivacy: "public",
		Avatar:         gitea_user.AvatarURL,
		Gender:         "male",
		FirstName:      strings.Split(gitea_user.FullName, " ")[0],
		LastName:       strings.Split(gitea_user.FullName, " ")[1],
	}

	user := models.User{
		UserName:  gitea_user.Login,
		Email:     gitea_user.Email,
		Password:  "",
		Provider:  models.Provider.Reboot,
		Following: make(map[string]bool),
	}
	// Check if the user exists in the database
	exists, userId := database.GetUserIDByProvider(user.UserName)
	if !exists {
		// Add the user to the database
		if err := database.SignUpUser(user, userProfile); err != nil {
			log.Printf("Error signing up user: %v", err)
			helpers.HTTPError(w, "Internal Server error", http.StatusInternalServerError)
			return
		}
	} else {
		sessionUUID, err := uuid.NewV4()
		if err != nil {
			log.Printf("Error creating session UUID: %v", err)
			helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
			return
		}
		if err := database.AddUserSession(userId, sessionUUID.String()); err != nil {
			log.Printf("Error adding session: %v", err)
			helpers.HTTPError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Set a cookie with a session token that can be used to authenticate access without logging in
		session.SetAutherizationHeader(w, sessionUUID.String())
		redirectURL := fmt.Sprintf("http://localhost:3000/?session_id=%s", sessionUUID.String())
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}

//gihtub login handler

const (
	githubClientID     = "1e7aa20b6eaf1123d2ab"
	githubClientSecret = "8b4215e898cbb5f6fab983bcd7e23f88cba46db6"
	githubRedirectURI  = "http://localhost:8081/api/auth/github/login"
)

func HandleGithubLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	models.Code = r.URL.Query().Get("code")
	fmt.Println(models.Code)
	if models.Code == "" {
		params := url.Values{}
		params.Add("client_id", githubClientID)
		params.Add("redirect_uri", githubRedirectURI)
		params.Add("scope", "user:email") // Include user:email scope
		params.Add("state", "github")
		redirectURL := "https://github.com/login/oauth/authorize?client_id=1e7aa20b6eaf1123d2ab"
		Testing(w, r, redirectURL)
	} else {
		HandleGithubCallback(w, r)
	}
}

func Testing(w http.ResponseWriter, r *http.Request, redirect string) {
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func HandleGithubCallback(w http.ResponseWriter, r *http.Request) {

	fmt.Println(models.Code)

	tokenURL := "https://github.com/login/oauth/access_token"
	data := url.Values{}
	data.Set("code", models.Code)
	data.Set("client_id", githubClientID)
	data.Set("client_secret", githubClientSecret)
	data.Set("redirect_uri", githubRedirectURI)
	data.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		log.Fatal("Error getting token:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
		return
	}

	log.Println("Token response:", string(body)) // Log token response for debugging

	// Extract access token from token response
	accessToken := ExtractAccessToken(string(body))
	if accessToken == "" {
		log.Fatal("Access token not found in token response")
		return
	}

	// Use the access token for further requests
	userInfoURL := "https://api.github.com/user"
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal("Error getting user info:", err)
		return
	}
	defer resp.Body.Close()

	userInfoBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading user info response:", err)
		return
	}

	var userInfo map[string]interface{}
	if err := json.Unmarshal(userInfoBody, &userInfo); err != nil {
		log.Fatal("Error parsing user info:", err)
		return
	}
	fmt.Println(userInfo)
	Fullname := strings.Split(userInfo["name"].(string), " ")
	userProfile := models.UserProfile{
		NickName:       userInfo["login"].(string),
		ProfilePrivacy: "public",
		Avatar:         userInfo["avatar_url"].(string),
		Gender:         "male",
		FirstName:      Fullname[0],
		LastName:       Fullname[1],
	}

	user := models.User{
		UserName:  userInfo["login"].(string),
		Email:     userInfo["email"].(string),
		Password:  "",
		Provider:  models.Provider.Github,
		Following: make(map[string]bool),
	}
	// Check if the user exists in the database
	exists, userId := database.GetUserIDByProvider(user.UserName)
	if !exists {
		// Add the user to the database
		if err := database.SignUpUser(user, userProfile); err != nil {
			log.Printf("Error signing up user: %v", err)
			helpers.HTTPError(w, "Internal Server error", http.StatusInternalServerError)
			return
		}
	} else {
		sessionUUID, err := uuid.NewV4()
		if err != nil {
			log.Printf("Error creating session UUID: %v", err)
			helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
			return
		}
		if err := database.AddUserSession(userId, sessionUUID.String()); err != nil {
			log.Printf("Error adding session: %v", err)
			helpers.HTTPError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Set a cookie with a session token that can be used to authenticate access without logging in
		session.SetAutherizationHeader(w, sessionUUID.String())
		redirectURL := fmt.Sprintf("http://localhost:3000/?session_id=%s", sessionUUID.String())
		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}

func ExtractAccessToken(body string) string {
	params, err := url.ParseQuery(body)
	if err != nil {
		log.Println("Error parsing token response:", err)
		return ""
	}
	return params.Get("access_token")
}
