//go:generate oapi-codegen --config ../../openapi/types.server.cfg.yaml 	../../openapi/url-shortener-api.yaml
//go:generate oapi-codegen --config ../../openapi/server.cfg.yaml 			../../openapi/url-shortener-api.yaml
//go:generate oapi-codegen --config ../../openapi/embedded-spec.cfg.yaml 	../../openapi/url-shortener-api.yaml

package oapi
