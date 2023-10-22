package controllers

import (
	"net/http"
	"strconv"

	"github.com/eduardo-paes/clean-go/entities"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Users map[int]*entities.User
	Seq   int
}

func NewUserController() *UserController {
	return &UserController{
		Users: make(map[int]*entities.User),
		Seq:   1,
	}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	u := &entities.User{ID: uc.Seq}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user data")
	}

	uc.Users[u.ID] = u
	uc.Seq++
	return c.JSON(http.StatusCreated, u)
}

func (uc *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, ok := uc.Users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUsers(c echo.Context) error {
	skip, err := strconv.Atoi(c.QueryParam("skip"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid 'skip' parameter")
	}

	take, err := strconv.Atoi(c.QueryParam("take"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid 'take' parameter")
	}

	allUsers := []*entities.User{}
	end := skip + take
	if end > len(uc.Users) {
		end = len(uc.Users)
	}

	for i := skip; i < end; i++ {
		allUsers = append(allUsers, uc.Users[i])
	}

	if len(allUsers) == 0 {
		return c.JSON(http.StatusNotFound, "No users found")
	}

	return c.JSON(http.StatusOK, allUsers)
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	u := &entities.User{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user data")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, ok := uc.Users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	user.Name = u.Name
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	_, ok := uc.Users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	delete(uc.Users, id)
	return c.NoContent(http.StatusNoContent)
}
