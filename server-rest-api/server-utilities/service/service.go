package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

const privateKeyPath = "server-utilities/service/keys/key.rsa"
const pubKeyPath = "server-utilities/service/keys/key.rsa.pub"

var SignKey []byte
var VerifyKey []byte

var (
	errorWhileDecoding  = errors.New("Error while parsing body")
	errorBadCredentials = errors.New("Bad credentials")
)

func init() {
	initKey()
}

// Function
// Infos: This function is used to load the private and public key
func initKey() {
	var err error

	SignKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Printf(" error : %s", err)
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

//UserCredentials : "Structure that store user login information"
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Token:"Strore the token"
type Token struct {
	Token string `json:"token"`
}

//"Login : service that verify the argument pass to login and initialise a session for this user"
func Login(ctx *iris.Context) (*Token, error) {

	var user UserCredentials

	//decode request into UserCredentials struct
	r := (*ctx).Request()

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		(*ctx).StatusCode(iris.StatusForbidden)
		log.Printf("Error while decoding %s ", err)
		return nil, errorWhileDecoding
	}

	//validate user credentials
	if strings.ToLower(user.Username) != "alexcons" {
		if user.Password != "kappa123" {
			(*ctx).StatusCode(iris.StatusForbidden)
			log.Println("Invalid credentials")
			return nil, errorBadCredentials
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin":    true,
		"username": user.Username,
		"password": user.Password,
	})

	tokenString, err := token.SignedString(SignKey)

	if err != nil {
		(*ctx).StatusCode(iris.StatusInternalServerError)
		log.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
		return &Token{}, errors.New("Error signing the token : " + err.Error())
	}

	//create a token instance using the token string
	response := Token{tokenString}

	return &response, nil

}

//Infos: This function goal is to get the Claims that contains the user information from the jwt
func GetClaims(ctx *iris.Context) (jwt.Claims, error) {
	user, ok := (*ctx).Values().Get("jwt").(*jwt.Token)
	if ok != true {
		(*ctx).StatusCode(iris.StatusInternalServerError)
		return nil, errors.New("Assertion type failed")
	}
	claims := user.Claims
	return claims, nil
}
