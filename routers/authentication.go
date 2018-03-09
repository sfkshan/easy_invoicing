package routers

import (
	"github.com/labstack/echo"
	"github.com/sfkshan/my-go/controllers"
)

func SetAuthenticationRoutes(e *echo.Echo) {
	e.POST("/login", controllers.Login)
}
