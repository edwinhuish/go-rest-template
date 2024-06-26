package test

import (
	"fmt"
	"testing"

	"github.com/edwinhuish/go-rest-template/internal/config"
	"github.com/edwinhuish/go-rest-template/internal/db"
	models "github.com/edwinhuish/go-rest-template/internal/models/users"
	"github.com/edwinhuish/go-rest-template/internal/repos"
)

var userTest models.User

func Setup() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM users")
}

func TestAddUser(t *testing.T) {
	Setup()
	user := models.User{
		Firstname: "Antonio",
		Lastname:  "Paya",
		Username:  "antonio",
		Hash:      "hash",
		Role:      models.UserRole{RoleName: "user"},
	}
	s := repos.GetUserRepository()
	if err := s.Add(&user); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	userTest = user
}

func TestGetAllUsers(t *testing.T) {
	s := repos.GetUserRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	db.SetupDB()
	db.SetupDB()
	s := repos.GetUserRepository()
	if _, err := s.Get(fmt.Sprint(userTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
