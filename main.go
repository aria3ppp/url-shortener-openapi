package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aria3ppp/url-shortener-openapi/internal/core/usecase"
	"github.com/aria3ppp/url-shortener-openapi/internal/generator"
	"github.com/aria3ppp/url-shortener-openapi/internal/oapi"
	"github.com/aria3ppp/url-shortener-openapi/internal/repository"
	"github.com/aria3ppp/url-shortener-openapi/internal/server"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)
	generator := generator.NewRandomStringGenerator(6)

	serviceUseCases := usecase.NewService(repo, generator)

	//--------------------------------------------------------------------------

	// handler := handler.NewHandler(serviceUseCases)

	// router := echo.New()
	// helper.HandleRoutes(router, handler)

	// if err := router.Start(":" + os.Getenv("SERVER_PORT")); err != nil {
	// 	panic(err)
	// }

	//--------------------------------------------------------------------------

	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	// swagger.Servers = nil

	serverImpl := server.New(serviceUseCases)

	e := echo.New()
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
	}))

	oapi.RegisterHandlers(e, serverImpl)

	if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil {
		panic(err)
	}
}
