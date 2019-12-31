package models

import (
	"github.com/Remydeme/esme-devops-project/db"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	bongo.DocumentBase `bson:",inline"`
	PhoneNumber        string            `json:"phoneNumber" bson:"phoneNumber"`
	DeviceNumber       string            `json:"deviceNumber" bson:"deviceNumber"`
	ActiveSessionId    string            `json:"activeSessionID" bson:"activeSession"`
	Firstname          string            `json:"firstName" bson:"firstname"`
	Lastname           string            `json:"lastName" bson:"lastname"`
	PrivacyRequested   bool              `json:"privacyRequested" bson:"privacyRequest"`
	Gender             string            `json:"gender" bson:"gender"`
	Birthdate          string            `json:"birthdate" bson:"birthdate"`
	Profession         string            `json:"profession" bson:"profession"`
	Company            string            `json:"company" bson:"company"`
	Siret              string            `json:"siret" bson:"siret"`
	Adress             Adress            `json:"adress" bson:"adress"`
	Medical            MedicalInfo       `json:"medicalInformations" bson:"medical"`
	EmergencyContacts  EmergencyContacts `json:"emergencyContacts" bson:"emergencyContacts"`
	Position           Position          `json:"position" bson:"position"`
	AccessData         AccessData        `json:"accessData" bson:"accessData"`
	comment            string            `json:"comment" bson:"comment"`
	Pictures           Pictures          `json:"pictures" bson:"pictures"`
}

// function that check is a user already exists in the database
// Improve this method this method is not optimized because it fetch the object and load it
// It's better to just check if the user exists
// use the official documentation of mangoDB
func (p *Person) exists(session *bongo.Connection, key string) bool {
	tempPerson := &Person{}
	err := session.Collection("Person").FindOne(bson.M{"phoneNumber": key}, &tempPerson)
	if err != nil {
		return false
	}
	return true
}

// function used to add a new Person to the database
func (p *Person) Add(session *bongo.Connection) error {

	// need to check if a user with the same number exists in the database
	if p.exists(session, p.PhoneNumber) == true {
		return db.ErrorObjectAlreadyExists
	}
	// save the object to the database
	err := session.Collection("Person").Save(p)
	if err != nil {
		vErr, ok := err.(*bongo.ValidationError)
		if ok == true {
			log.Println("Validation erross are:", vErr.Errors)
			return db.ErrorValidation // Errors is an array of error => []error
		} else {
			log.Println("Got a real error:", err.Error())
			return db.ErrorOccured
		}
	}
	return nil
}

func (p *Person) Delete(session *bongo.Connection, key string) error {

	if p.exists(session, p.PhoneNumber) == true {
		return db.ErrorObjectAlreadyExists
	}

	changeInfo, err := session.Collection("Person").Delete(bson.M{"phoneNumber": key})

	if err != nil {
		return db.ErrorOccured
	}

	log.Printf("Deleted %d documents", changeInfo.Removed)

	return nil
}
