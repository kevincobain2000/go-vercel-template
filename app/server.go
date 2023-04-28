package app

import (
	"net/http"

	"github.com/charmbracelet/log"
	h "github.com/kevincobain2000/go-vercel-template/handlers"
	"github.com/kevincobain2000/go-vercel-template/models"
	"github.com/kevincobain2000/go-vercel-template/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HTTPServer() *echo.Echo {
	log.Info("HTTP Server ")
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = HTTPErrorHandler

	dbMigrate()
	routes(e)
	return e
}

func routes(e *echo.Echo) *echo.Echo {
	log.Info("Registering Routes")
	h := h.NewIndexHandler()
	e.GET("/", h.IndexHandler)
	return e
}

func dbMigrate() {
	log.Info("db migrate")
	pkg.DB().AutoMigrate(&models.User{})
}

// HTTPErrorResponse is the response for HTTP errors
type HTTPErrorResponse struct {
	Error interface{} `json:"error"`
}

// HTTPErrorHandler handles HTTP errors for entire application
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		message = err.Error()
	}

	if err = c.JSON(code, &HTTPErrorResponse{Error: message}); err != nil {
		log.Error("HTTP", "error", err)
	}
}
