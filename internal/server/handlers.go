package server

import (
	"gotvup/frontend/assets"
	"gotvup/frontend/templates"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	fileServer := http.FileServer(http.FS(assets.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))
	e.GET("/", helloWorldHandler)
}

func helloWorldHandler(c echo.Context) error {
	//return Render(c, http.StatusOK, templates.HelloForm())
	return RenderWithinBase(c, http.StatusOK, templates.HelloForm())
}
