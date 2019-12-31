package models

type Contact struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	Number       string `json:"number"`
}

type EmergencyContacts struct {
	Contacts *[]Contact
}
