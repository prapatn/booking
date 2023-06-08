package controllers

import (
	"booking/entities/users"
	"booking/usecases"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	user := new(users.Register)
	err := c.Bind(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = user.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user.HashPassword()
	err = usecases.CreateUser(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func SetUserGrade(c echo.Context) error {
	user := new(users.Grade)
	err := c.Bind(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = user.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user.Sum()
	err = usecases.UpdateGrade(user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}
