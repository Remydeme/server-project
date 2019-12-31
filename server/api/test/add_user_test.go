package main

import "testing"

import (
	"bytes"
	"encoding/json"
	"github.com/Remydeme/esme-devops-project/api"
	"github.com/Remydeme/esme-devops-project/db/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var users = []*models.Person{
	&models.Person{DeviceNumber: "ID"},
	&models.Person{DeviceNumber: "ID", PhoneNumber: "number", ActiveSessionId: "session id", Firstname: "remy", Lastname: "deme", PrivacyRequested: true, Gender: "male", Birthdate: "06/08/1996"},
}

// test create to add new user in the database
// request : POST /user/add/
func TestAddNewUser(t *testing.T) {
	handler := api.CreateUser
	person, _ := json.Marshal(users[1])
	body := bytes.NewBuffer(person)
	req, _ := http.NewRequest("POST", "user/add/", body)
	w := httptest.NewRecorder()
	handler(w, req)

	response := w.Result()
	bod, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, response.StatusCode, http.StatusOK, "Unexpected status code ")
	assert.Equal(t, "application/vnd.api+json", response.Header.Get("Content-Type"), "bad content type")
	assert.Equal(t, "/", string(bod), "User have been added succesfuly")
}
