package server

import (
	"context"
	"gotvup/frontend/templates"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomContext struct {
	echo.Context
}

func (c CustomContext) Get(key string) interface{} {
	// grab value from echo.Context
	val := c.Context.Get(key)
	// if it's not nil, return it
	if val != nil {
		return val
	}
	// otherwise, return Request.Context
	return c.Request().Context().Value(key)
}

func buildContext(c echo.Context) context.Context {
	csrf := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	return context.WithValue(c.Request().Context(), templates.CsrfContextKey, csrf)
}

func RenderError(c echo.Context, status int, err error) error {
	return Render(c, status, templates.Error(err))
}

func RenderWithinBase(c echo.Context, status int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	newContext := buildContext(c)
	html, err := templ.ToGoHTML(newContext, t)
	if err != nil {
		return RenderError(c, http.StatusInternalServerError, err)
	}
	err = templates.Base.Execute(buf, templates.BaseContent{
		Main: html,
	})
	if err != nil {
		return RenderError(c, http.StatusInternalServerError, err)
	}
	return c.HTML(status, buf.String())
}

func Render(c echo.Context, status int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	newContext := buildContext(c)
	if err := t.Render(newContext, buf); err != nil {
		return err
	}

	return c.HTML(status, buf.String())
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	RenderError(c, code, err)
}
