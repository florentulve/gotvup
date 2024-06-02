package server

import (
	"fmt"
	"gotvup/frontend/templates"
	"os"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

func StartWebServer() {
	port := os.Getenv("PORT")

	e := echo.New()
	e.HTTPErrorHandler = CustomHTTPErrorHandler
	e.Use(emw.Logger())
	e.Use(emw.Timeout())
	e.Use(emw.BodyLimit("2M"))
	e.Use(emw.CORS())
	e.Use(emw.CSRFWithConfig(emw.CSRFConfig{
		TokenLookup: "form:_csrf",
		ContextKey:  string(templates.CsrfContextKey),
	}))
	e.Use(emw.Gzip())
	e.Use(emw.SecureWithConfig(emw.SecureConfig{
		ContentSecurityPolicy: "default-src 'self'",
	}))

	RegisterRoutes(e)

	e.Start(fmt.Sprintf(":%s", port))
}
