package gothex

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type GothexRouter struct {
	*echo.Echo
	configs     *Configs
	CookieStore *sessions.CookieStore
	Protected   func(next echo.HandlerFunc) echo.HandlerFunc
}

type ErrorPageContent struct {
	Code      int
	Title     string
	ErrorType string
	Message   string
}

type ContextKey struct {
	Key string
}

func createRouter(configs *Configs) *GothexRouter {
	router := echo.New()
	router.Static("/", "public")
	router.StaticFS("/lib", echo.MustSubFS(assetFiles, "assets"))
	sessionAge, _ := strconv.Atoi(configs.SessionAge)
	cookieStore := sessions.NewCookieStore([]byte(configs.GothexSecret))
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.MaxAge = sessionAge
	cookieStore.Options.Path = configs.AfterSignin
	router.Use(session.Middleware(cookieStore))
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("X-Title", configs.AppTitle)
			return next(c)
		}
	})

	router.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			if !strings.HasPrefix(c.Request().RequestURI, "/api") {
				errPage := ErrorPageContent{Code: code}
				errPage.ErrorType = he.Message.(string)
				errPage.Title = fmt.Sprintf("Error (%d)", code)
				if page := IsCustomErrorPages(code); page != nil {
					errPage = *page
				}
				c.Set("X-Title", errPage.Title)
				c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
				tmpl.Execute(c.Response().Writer, errPage)
				return
			}
		}
		c.JSON(code, err)
	}

	return &GothexRouter{
		router,
		configs,
		cookieStore,
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				session, authenticated := GetAuthSession(c)
				if !authenticated {
					if IsHxRequest(c) {
						return SignOut(c, configs.SignInPage)
					}
					return c.Redirect(http.StatusMovedPermanently, configs.SignInPage)
				}
				sessionAge, _ := strconv.Atoi(configs.SessionAge)
				session.Options.MaxAge = sessionAge
				if err := session.Save(c.Request(), c.Response()); err != nil {
					return err
				}
				return next(c)
			}
		},
	}
}

func NewGothexRouter() *GothexRouter {
	configs := &Configs{}
	if err := configs.ReadFromEnv(); err != nil {
		log.Fatal("failed to get configs", err)
	}
	return createRouter(configs)
}

func (c *Configs) ReadFromEnv() error {
	godotenv.Load()
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := env.Parse(c); err != nil {
		return errors.WithStack(err)
	} else {
		return validate.Struct(c)
	}
}

func NewRouterWithConfigs(configs *Configs) *GothexRouter {
	return createRouter(configs)
}

func (r *GothexRouter) WithCustomErrorPageContent(pages ...ErrorPageContent) *GothexRouter {
	customErrorPageContent = pages
	return r
}

func (r *GothexRouter) WithErrorHandler(handler echo.HTTPErrorHandler) *GothexRouter {
	r.HTTPErrorHandler = handler
	return r
}

func (r *GothexRouter) WithNoCache() *GothexRouter {
	r.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Response().Header().Set("Pragma", "no-cache")
			c.Response().Header().Set("Expires", "0")
			return next(c)
		}
	})
	return r
}

func (r *GothexRouter) Run() error {
	return r.Start(fmt.Sprintf(":%s", r.configs.Port))
}
