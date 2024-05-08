package controllers

import (
	"log"
	"net/http"

	"github.com/edwinhuish/go-rest-template/internal/api/gin2"
	"github.com/edwinhuish/go-rest-template/internal/persistence"
	"github.com/edwinhuish/go-rest-template/internal/pkg/crypto"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ctrl *AuthController) Login(c *gin2.Context) {

	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("user not found")
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			c.Status(http.StatusForbidden)
			c.Resp().Fail("user and password not match")
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		c.Resp().Success(token)
	}
}
