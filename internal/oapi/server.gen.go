// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package oapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /link)
	CreateLink(ctx echo.Context) error
	// Your GET endpoint
	// (GET /link/{shortened_string})
	GetLink(ctx echo.Context, shortenedString ShortenedString) error
	// Your GET endpoint
	// (GET /link/{shortened_string}/user)
	GetLinkUser(ctx echo.Context, shortenedString ShortenedString) error

	// (POST /user)
	CreateUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateLink converts echo context to params.
func (w *ServerInterfaceWrapper) CreateLink(ctx echo.Context) error {
	var err error

	ctx.Set(Username_passwordScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateLink(ctx)
	return err
}

// GetLink converts echo context to params.
func (w *ServerInterfaceWrapper) GetLink(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shortened_string" -------------
	var shortenedString ShortenedString

	err = runtime.BindStyledParameterWithLocation("simple", false, "shortened_string", runtime.ParamLocationPath, ctx.Param("shortened_string"), &shortenedString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shortened_string: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetLink(ctx, shortenedString)
	return err
}

// GetLinkUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetLinkUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "shortened_string" -------------
	var shortenedString ShortenedString

	err = runtime.BindStyledParameterWithLocation("simple", false, "shortened_string", runtime.ParamLocationPath, ctx.Param("shortened_string"), &shortenedString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter shortened_string: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetLinkUser(ctx, shortenedString)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/link", wrapper.CreateLink)
	router.GET(baseURL+"/link/:shortened_string", wrapper.GetLink)
	router.GET(baseURL+"/link/:shortened_string/user", wrapper.GetLinkUser)
	router.POST(baseURL+"/user", wrapper.CreateUser)

}
