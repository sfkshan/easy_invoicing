package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sfkshan/my-go/controllers/resources"
	"github.com/sfkshan/my-go/models"
	model "gopkg.in/jeevatkm/go-model.v1"
)

func Login(c echo.Context) error {
	userForm := &resources.User{}
	c.Bind(userForm)

	user := models.User{}

	// source, destination
	errs := model.Copy(&user, userForm)
	fmt.Println("Errors:", errs)

	return c.JSON(http.StatusOK, user)
}
