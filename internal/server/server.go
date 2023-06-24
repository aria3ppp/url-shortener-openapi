package server

import (
	"errors"
	"net/http"

	"github.com/aria3ppp/url-shortener-openapi/internal/core/domain"
	domain_errors "github.com/aria3ppp/url-shortener-openapi/internal/core/errors"
	"github.com/aria3ppp/url-shortener-openapi/internal/core/port"
	"github.com/aria3ppp/url-shortener-openapi/internal/oapi"
	"github.com/aria3ppp/url-shortener-openapi/internal/validate"
	"github.com/labstack/echo/v4"
)

type Server struct {
	serviceUseCases port.ServiceUseCases
}

var _ oapi.ServerInterface = &Server{}

func New(serviceUseCases port.ServiceUseCases) *Server {
	return &Server{serviceUseCases: serviceUseCases}
}

func (s *Server) CreateLink(c echo.Context) error {
	// parse and validate url and shortened string
	var body oapi.CreateLinkRequestBody
	if httpError := (&echo.DefaultBinder{}).BindBody(c, &body); httpError != nil {
		return httpError
	}
	if err := validate.CreateLinkRequestBody(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if body.ShortenedString == nil {
		body.ShortenedString = new(string)
		*body.ShortenedString = ""
	}

	// fetch username:password off the basic authorization
	username, password, ok := c.Request().BasicAuth()
	if !ok {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"basic authorization not provided",
		)
	}

	link, err := s.serviceUseCases.CreateLink(
		body.Url,
		*body.ShortenedString,
		&domain.User{Username: username, Password: password},
	)
	if err != nil {
		if errors.Is(err, domain_errors.ErrUserNotFound) ||
			errors.Is(err, domain_errors.ErrIncorrectPassword) {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"invalid username or password",
			)
		}
		if errors.Is(err, domain_errors.ErrUsedShortenedString) {
			return echo.NewHTTPError(
				http.StatusConflict,
				"shortened string have used",
			)
		}
		return echo.NewHTTPError(http.StatusInternalServerError).
			SetInternal(err)
	}

	return c.JSON(http.StatusOK, oapi.CreateLinkResponseBody{
		ShortenedString: link.ShortenedString,
		Url:             link.URL,
		Username:        link.Username,
	})
}

func (s *Server) GetLink(
	c echo.Context,
	shortenedString oapi.ShortenedString,
) error {
	link, err := s.serviceUseCases.GetLink(shortenedString)
	if err != nil {
		if errors.Is(err, domain_errors.ErrLinkNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError).
			SetInternal(err)
	}

	return c.Redirect(http.StatusPermanentRedirect, link.URL)
}

func (s *Server) GetLinkUser(
	c echo.Context,
	shortenedString oapi.ShortenedString,
) error {
	user, err := s.serviceUseCases.GetLinkUser(shortenedString)
	if err != nil {
		if errors.Is(err, domain_errors.ErrLinkNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "link not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError).
			SetInternal(err)
	}

	return c.JSON(http.StatusOK, oapi.GetLinkUserResponseBody{
		Username: user.Username,
	})
}

func (s *Server) CreateUser(c echo.Context) error {
	// parse and validate username and password
	var body oapi.CreateUserRequestBody
	if httpError := (&echo.DefaultBinder{}).BindBody(c, &body); httpError != nil {
		return httpError
	}
	if err := validate.CreateUserRequestBody(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := s.serviceUseCases.CreateUser(&domain.User{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		if errors.Is(err, domain_errors.ErrUsernameTaken) {
			return echo.NewHTTPError(
				http.StatusConflict,
				"username have taken",
			)
		}
		return echo.NewHTTPError(http.StatusInternalServerError).
			SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
