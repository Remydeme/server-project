package models

import (
	"errors"
	"github.com/Remydeme/iurgence/db"
	"github.com/badoux/checkmail"
	"github.com/go-bongo/bongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Credential struct {
	bongo.DocumentBase `bson:",inline"`
	Password           string `json:"password" bson:"password"`
	Email              string `json:"email" bson:"email"`
	PhoneNumber        string `json:"phoneNumber" bson:"phoneNumber"`
}

func (c *Credential) HashAndSalt() ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (c *Credential) CheckEmail() error {
	err := checkmail.ValidateFormat(c.Email)
	log.Printf("Email : %s ", c.Email)
	return err
}

func (c *Credential) CheckPassword() error {
	passwordLen := len(c.Password)
	log.Printf("Email : %s ", c.Password)
	if passwordLen < 8 {
		return errors.New("Password format is invalid. Too short.")
	}
	return nil
}

func (c *Credential) ComparePassword(hashPassword []byte) bool {

	plainPwd := []byte(c.Password)

	err := bcrypt.CompareHashAndPassword(hashPassword, plainPwd)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// check if a user already exists in the database
// checking if the user e-mail or phone number is in the database
func (c *Credential) Exists(session *bongo.Connection, key string) bool {

	cred := &Credential{}

	err := session.Collection("Credential").FindOne(bson.M{"phoneNumber": key}, &cred)

	if err != nil {
		// this means that the user doesn't exist
		return false
	}
	return true
}

// Look for the user in the database using it's phoneNumber as key
func (c *Credential) Fetch(session *bongo.Connection, key string) (*Credential, error) {

	cred := &Credential{}

	err := session.Collection("Credential").FindOne(bson.M{"phoneNumber": key}, &cred)

	if err != nil {
		return nil, err
	}
	return cred, nil
}

// func used to create a new object
func (c *Credential) Add(session *bongo.Connection) error {

	if c.Exists(session, c.PhoneNumber) == true {
		return db.ErrorObjectAlreadyExists
	}

	hash, err := c.HashAndSalt()

	if err != nil {
		return err
	}

	// store the hash
	c.Password = string(hash)

	err = session.Collection("Credential").Save(c)

	if err != nil {

		// check it's validation error
		validationErr, ok := err.(*bongo.ValidationError)
		if ok == true {
			log.Printf("Validation error : %s \n", validationErr.Error())
			return db.ErrorValidation
		} else {
			log.Printf("An error occured : %s \n", err.Error())
			return db.ErrorOccured
		}
	}

	return nil
}
