package models_test

import (
	"testing"

	"github.com/mikhailtse/go-simple-webservice/models"
)

func Test_GetUsers_ReturnEmptyList(t *testing.T) {
	users := models.GetUsers()
	if len(users) != 0 {
		t.Error("did not get empty list of the users")
	}
}

func Test_GetUsers_ReturnsNonEmptyList(t *testing.T) {
	models.AddUser(models.User{FirstName: "Samwell", LastName: "Tarly"})
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("did not get non-empty list of users")
	}
}

func Test_AddUser_ReturnsAddedUser(t *testing.T) {
	firstName := "John"
	lastName := "Snow"
	u, err := models.AddUser(models.User{FirstName: firstName, LastName: lastName})
	if err != nil {
		t.Error(err)
	}
	if u.FirstName != firstName {
		t.Errorf("did not get expected result. Got: '%v', wanted: '%v'", u.FirstName, firstName)
	}
	if u.LastName != lastName {
		t.Errorf("did not get expected result. Got: '%v', wanted: '%v'", u.LastName, lastName)
	}
	if u.Id == 0 {
		t.Errorf("did not get expected result. Got: '%v', wanted: int != 0", u.Id)
	}
}

func Test_AddUser_ReturnsErrorWhenIdIsNotZero(t *testing.T) {
	firstName := "False John"
	lastName := "Wrong Snow"
	_, err := models.AddUser(models.User{Id: 1, FirstName: firstName, LastName: lastName})
	if err == nil {
		t.Error("did not get expected error")
	}
}

func Test_AddUser_IncrementsId(t *testing.T) {
	u1, _ := models.AddUser(models.User{FirstName: "Daenerys", LastName: "Targaryen"})
	u2, _ := models.AddUser(models.User{FirstName: "Joffrey", LastName: "Baratheon"})
	if u1.Id == u2.Id {
		t.Error("got users with the same ids")
	}
	if u1.Id+1 != u2.Id {
		t.Error("did not increment user id by 1")
	}
}
