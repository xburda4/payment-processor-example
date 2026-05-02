package api

import (
	v1 "fintech-proj/api/v1"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	httpx "go.strv.io/net/http"

	"fintech-proj/util/logger"
)

//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config oapi-codegen-cfg.yaml swagger-ui/openapi.yaml

type SwaggerUIAuth struct {
	Name string `env:"SWAGGER_UI_NAME" validate:"required"`
	//nolint:gosec
	Password string `env:"SWAGGER_UI_PASSWORD" validate:"required"`
}

// Controller handles all /api endpoints.
// It is responsible for routing requests to appropriate handlers.
// Versioned endpoints are handled by subcontrollers.
type Controller struct {
	*chi.Mux

	version     string
	environment string
}

func NewController(
	version string,
	environment string,
) (*Controller, error) {
	controller := &Controller{
		version:     version,
		environment: environment,
	}
	controller.InitRouter()
	return controller, nil
}

func (c *Controller) InitRouter() {
	r := chi.NewRouter()

	v1Handler := v1.NewHandler()

	r.Group(func(r chi.Router) {
		r.Use(httpx.RequestIDMiddleware(func(h http.Header) string {
			return h.Get(httpx.Header.XRequestID)
		}))
		r.Use(httpx.LoggingMiddleware(logger.Default().With(slog.String("caller", "httpx.LoggingMiddleware"))))
		r.Use(httpx.RecoverMiddleware(logger.Default().With(slog.String("caller", "httpx.RecoverMiddleware"))))

		r.Route("/api", func(r chi.Router) {

			r.Mount("/v1", v1Handler)
		})
	})

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	c.Mux = r
}
