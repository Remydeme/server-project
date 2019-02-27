package database

type Users struct {
	Login    string `json:login`
	Password string `json:password`
}
