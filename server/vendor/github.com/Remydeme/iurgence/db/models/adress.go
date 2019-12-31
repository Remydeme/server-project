package models

type Adress struct {
	Type            string `json:"type"`
	Number          string `json:"number"`
	ExtensionNumber string `json:"extensionNumber"`
	Street          string `json:"Street"`
	coordinate      Point  `json:"coordinate"`
	Locality        string `json:"locality"`
	Region          string `json:"region"`
	PostalOfficeBox string `json:"postalOfficeBox"`
	District        string `json:"district"`
	PostalCode      string `json:"postalCode"`
	InseeCode       string `json:"inseeCode"`
	CountryCode     string `json:"countryCode"`
	Apartment       string `json:"apartment"`
	Floor           string `json:"floor"`
	Building        string `json:"building"`
	SecureAccess    string `json:"secureAccess"`
	Comment         string `json:"comment"`
}
