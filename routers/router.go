package routers

import (
	"github.com/labstack/echo"
)

// Init : adds all the routes to echo
func Init(e *echo.Echo) {
	SetAuthenticationRoutes(e)
}
