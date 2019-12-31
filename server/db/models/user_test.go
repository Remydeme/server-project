package models

import "testing"
import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

var users = []*Person{
	&Person{DeviceNumber: "ID"},
	&Person{DeviceNumber: "ID", PhoneNumber: "number", ActiveSessionId: "session id", Firstname: "remy", Lastname: "deme", PrivacyRequested: true, Gender: "male", Birthdate: "06/08/1996"},
}

// test create a user structure
func TestCreateUser(t *testing.T) {

	for _, user := range users {
		fmt.Printf("%v", user)
		assert.Equal(t, user.DeviceNumber, "ID", "Unexpected value id")
	}
}
