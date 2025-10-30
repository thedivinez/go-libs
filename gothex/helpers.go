package gothex

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Configs struct {
	Port         string `env:"PORT" validate:"required"`
	AppTitle     string `env:"APP_TITLE" validate:"required"`
	SessionAge   string `env:"SESSION_AGE" validate:"required"`
	SignInPage   string `env:"SIGNIN_PAGE" validate:"required"`
	GothexSecret string `env:"GOTHEX_SECRET" validate:"required"`
	AfterSignin  string `env:"AFTER_SIGNIN_PAGE" validate:"required"`
}

var customErrorPageContent []ErrorPageContent

func IsCustomErrorPages(code int) *ErrorPageContent {
	for _, page := range customErrorPageContent {
		if page.Code == code {
			return &page
		}
	}
	return nil
}

func GetAuthSession(c echo.Context) (*sessions.Session, bool) {
	if session, err := session.Get("__auth__", c); err != nil {
		return nil, false
	} else {
		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			return session, true
		} else {
			return session, false
		}
	}
}

func SignIn(c echo.Context, redirect string, user map[string]any) error {
	session, _ := GetAuthSession(c)
	if jsonString, err := json.Marshal(user); err != nil {
		return err
	} else {
		session.Values = map[any]any{"authenticated": true, "user": jsonString}
	}
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return HxRedirect(c, redirect)
}

func SignOut(c echo.Context, redirect string) error {
	session, _ := GetAuthSession(c)
	session.Values = map[any]any{"authenticated": false}
	delete(session.Values, "user")
	session.Options.MaxAge = -1
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return HxRedirect(c, redirect)
}

type TemplBody func(ctx context.Context, options ...any) templ.Component
type Page interface {
	Render(ctx context.Context, w io.Writer) error
}

func Render(c echo.Context, body TemplBody, data ...any) error {
	ctx := context.WithValue(c.Request().Context(), ContextKey{Key: "X-Title"}, c.Get("X-Title"))
	return body(ctx, data...).Render(ctx, c.Response().Writer)
}

func RenderWithTitle(c echo.Context, title string, body TemplBody, data ...any) error {
	c.Set("X-Title", title)
	return Render(c, body, data...)
}

func ShowComponent(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func IsHxRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func HxReload(c echo.Context) error {
	c.Response().Header().Set("HX-Refresh", "true")
	c.Response().WriteHeader(http.StatusMovedPermanently)
	return nil
}

func HxRedirect(c echo.Context, location string) error {
	c.Response().Header().Set("HX-Redirect ", location)
	return c.NoContent(http.StatusOK)
}

//go:embed "error.templ"
var errorComponentSource string

// Simple templ-to-html converter for your specific use case
func RenderErrorComponent(w io.Writer, message string) error {
	// Simple template replacement - extend based on your needs
	html := strings.ReplaceAll(errorComponentSource, "{ message }", message)
	html = strings.ReplaceAll(html, "templ ErrorComponent(message string) {", "")
	html = strings.ReplaceAll(html, "}", "")
	html = strings.TrimSpace(html)

	_, err := fmt.Fprint(w, html)
	return err
}
