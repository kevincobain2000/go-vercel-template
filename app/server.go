package app

import (
	"net/http"

	h "github.com/kevincobain2000/go-vercel-template/handlers"
	"github.com/kevincobain2000/go-vercel-template/models"
	"github.com/kevincobain2000/go-vercel-template/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func HTTPServer() *echo.Echo {
	log.Info("HTTP Server ")
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = HTTPErrorHandler

	// set logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "HTTP   [${time_custom}] latency=${latency_human} method=${method} uri=${uri} status=${status} error=${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

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
	if err := pkg.DB().AutoMigrate(&models.User{}); err != nil {
		log.Fatal("db migrate", "error", err)
	}
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
