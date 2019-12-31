package models

type AccessData struct {
	Building   string `json:"building"`
	Coordinate Point  `json:"coordinate"`
	StreetName string `json:"streetname"`
	PostalCode string `json:"postalCode"`
}
