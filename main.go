package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sfkshan/my-go/models"
	"github.com/sfkshan/my-go/routers"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// DB migration.
	db, err := gorm.Open("sqlite3", "./gorm.db")

	db.AutoMigrate(&models.User{})

	defer func() {
		fmt.Println(err)
		db.Close()
	}()

	// Echo app instantiation
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = renderer

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Hello world test template
	e.GET("/", func(c echo.Context) error {
		fmt.Println("=======")
		fmt.Println(c.QueryParam("name"))
		fmt.Println("=======")
		return c.Render(http.StatusOK, "home.html", c.QueryParam("name"))
	})

	e.GET("/p", func(c echo.Context) error {
		return c.Render(http.StatusOK, "products.html", c.QueryParam("name"))
	})

	routers.Init(e)

	e.Logger.Fatal(e.Start(":3500"))
}
