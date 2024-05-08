package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/edwinhuish/go-rest-template/internal/api/gin2"
	models "github.com/edwinhuish/go-rest-template/internal/models/users"
	"github.com/edwinhuish/go-rest-template/internal/pkg/crypto"
	"github.com/edwinhuish/go-rest-template/internal/repos"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// GetUserById godoc
//
//	@Summary		Retrieves user based on given ID
//	@Description	get User by ID
//	@Produce		json
//	@Param			id	path		integer	true	"User ID"
//	@Success		200	{object}	users.User
//	@Router			/api/users/{id} [get]
//	@Security		Authorization Token
func (ctrl *UserController) Find(c *gin2.Context) {
	s := repos.GetUserRepository()
	id := c.Param("id")
	if user, err := s.Get(id); err != nil {
		c.Resp().Fail(errors.New("user not found"))
		log.Println(err)
		return
	} else {
		c.Resp().Success(user)
	}
}

// GetUsers godoc
//
//	@Summary		Retrieves users based on query
//	@Description	Get Users
//	@Produce		json
//	@Param			username	query	string	false	"Username"
//	@Param			firstname	query	string	false	"Firstname"
//	@Param			lastname	query	string	false	"Lastname"
//	@Success		200			{array}	[]users.User
//	@Router			/api/users [get]
//	@Security		Authorization Token
func (ctrl *UserController) List(c *gin2.Context) {
	s := repos.GetUserRepository()
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.Query(&q); err != nil {
		c.Resp().Fail("users not found")
		log.Println(err)
	} else {
		c.Resp().Success(users)
	}
}

func (ctrl *UserController) Create(c *gin2.Context) {
	s := repos.GetUserRepository()
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		Role:      models.UserRole{RoleName: userInput.Role},
	}
	if err := s.Add(&user); err != nil {
		c.Resp().Fail(err)
		log.Println(err)
	} else {
		c.Resp().Success(user)
	}
}

func (ctrl *UserController) Update(c *gin2.Context) {
	s := repos.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("user not found")
		log.Println(err)
	} else {
		user.Username = userInput.Username
		user.Lastname = userInput.Lastname
		user.Firstname = userInput.Firstname
		user.Hash = crypto.HashAndSalt([]byte(userInput.Password))
		user.Role = models.UserRole{RoleName: userInput.Role}
		if err := s.Update(user); err != nil {
			c.Status(http.StatusNotFound)
			c.Resp().Fail(err)
			log.Println(err)
		} else {
			c.Resp().Success(user)
		}
	}
}

func (ctrl *UserController) Delete(c *gin2.Context) {
	s := repos.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		c.Status(http.StatusNotFound)
		c.Resp().Fail("user not found")
		log.Println(err)
	} else {
		if err := s.Delete(user); err != nil {
			c.Status(http.StatusNotFound)
			c.Resp().Fail(err)
			log.Println(err)
		} else {
			c.Status(http.StatusNoContent)
			c.Resp().Success()
		}
	}
}
