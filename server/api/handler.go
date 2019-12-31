package api

import (
	"encoding/json"
	"fmt"
	"github.com/Remydeme/esme-devops-project/api/errors"
	"github.com/Remydeme/esme-devops-project/config"
	"github.com/Remydeme/esme-devops-project/db"
	"github.com/Remydeme/esme-devops-project/db/models"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	ExpirationTime = 10000
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

/**
Handler call to check if the server is running
*/
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong..."))
}

func SignUpPage(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/Sinscrire.html"))
	tpl.Execute(w, nil)
	/*body, err := loadFile("templates/signup.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Printf(body, w)*/
}

func SignInPage(w http.ResponseWriter, r *http.Request) {
	body, err := loadFile("templates/Seconnecter.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Printf(body, w)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Add error checking for email that already exists
// Create and return json web token ? or User have to sign In again.
func SignUp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	credential := models.Credential{}

	err := json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		parseError := errors.Error{errors.ParseErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(parseError)
		return
	}

	// checking email adress
	err = credential.CheckEmail()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		emailError := errors.Error{errors.InvalidEmailFormatErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(emailError)
		return
	}

	err = credential.CheckPassword()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		passwordError := errors.Error{errors.InvalidPasswordFormatErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(passwordError)
		return
	}

	err = credential.Add(db.Session)

	// add the user in the database and handle errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		insertionError := errors.Error{errors.DatabaseInsertionErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(insertionError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Perfect you have been registered succesfuly"))
}

/**
Authentification function use to login the user
*/
func SignIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	credential := models.Credential{}

	err := json.NewDecoder(r.Body).Decode(&credential)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		parseError := errors.Error{errors.ParseErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(parseError)
		return
	}

	// checking email adress
	err = credential.CheckEmail()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		emailError := errors.Error{errors.InvalidEmailFormatErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(emailError)
		return
	}

	err = credential.CheckPassword()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		passwordError := errors.Error{errors.InvalidPasswordFormatErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(passwordError)
		return
	}

	credFetch, err := credential.Fetch(db.Session, credential.PhoneNumber)

	if err != nil {
		// Impossible to find the user in the database
		w.WriteHeader(http.StatusInternalServerError)
		failedFetchUserError := errors.Error{errors.InternalServerErrorId, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()}
		json.NewEncoder(w).Encode(failedFetchUserError)
		return
	}

	if credential.ComparePassword([]byte(credFetch.Password)) == false {
		w.WriteHeader(http.StatusUnauthorized)
		BadLoginError := errors.Error{errors.InternalServerErrorId, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()}
		json.NewEncoder(w).Encode(BadLoginError)
		return
	}

	// I create my token and store it in a cookie
	expirationTime := time.Now().Add(ExpirationTime * time.Minute)
	claims := &Claims{
		Username: credential.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// create the JWT
	tokenString, err := token.SignedString(config.Main.JWT.Secret)

	if err != nil {
		// error failed to sign JSON WEB Token
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("JWT error : %s", err.Error())
		jwtSignatureTokenError := errors.Error{errors.JWTSignatureErrorId, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()}
		json.NewEncoder(w).Encode(jwtSignatureTokenError)
		return
	}
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.WriteHeader(http.StatusOK)

	jwtBody := struct {
		Token string
	}{
		tokenString,
	}

	json.NewEncoder(w).Encode(jwtBody)
}

/**
Create a new user.
Read the json and create a user in database
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	person := &models.Person{}

	err := json.NewDecoder(r.Body).Decode(&person)

	log.Printf("Value person : %v \n\n\n", person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parseError := errors.Error{errors.ParseErrorId, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()}
		json.NewEncoder(w).Encode(parseError)
		return
	}

	err = person.Add(db.Session)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error : %s \n", err.Error())
		insertionError := errors.Error{errors.DatabaseInsertionErrorId, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error()}
		json.NewEncoder(w).Encode(insertionError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User have been added succesfuly"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleting .. user have been deleted"))
}

// UPDATE HANDLERS
