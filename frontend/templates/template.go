package templates

import (
	"context"
	"gotvup/frontend/assets"
	"html/template"

	"github.com/labstack/echo/v4/middleware"
)

type BaseContent struct {
	Main template.HTML
}

var Base = template.Must(template.New("index.html").ParseFS(assets.Files, "index.html"))

type CSRFKey string

var CsrfContextKey CSRFKey = CSRFKey(middleware.DefaultCSRFConfig.ContextKey)

func GetCsrf(c context.Context) string {
	if theme, ok := c.Value(CsrfContextKey).(string); ok {
		return theme
	}
	return ""
}
