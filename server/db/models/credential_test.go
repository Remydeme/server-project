package models

import (
	"testing"
)

type Test struct {
	Value     Credential
	TestValue string
	Expected  string
}

var cred = Credential{Password: "123456", Email: "demeremy@gmail.com"}

func TestCreateCredential(t *testing.T) {

	hash, err := cred.HashAndSalt()

	if err != nil {
		t.Errorf("Failed to hash password : %s \n", err.Error())
		return
	}
	if cred.ComparePassword(hash) != true {
		t.Errorf("Failes bad password %s \n", cred.Password)
	}
}
