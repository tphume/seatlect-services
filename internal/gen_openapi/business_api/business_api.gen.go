// Package business_api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package business_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Business defines model for Business.
type Business struct {
	Id           *string   `json:"_id,omitempty"`
	Address      *string   `json:"address,omitempty"`
	BusinessName *string   `json:"businessName,omitempty"`
	Description  *string   `json:"description,omitempty"`
	DisplayImage *string   `json:"displayImage,omitempty"`
	Images       *[]string `json:"images,omitempty"`
	Location     *struct {
		Latitude  *string `json:"latitude,omitempty"`
		Longitude *string `json:"longitude,omitempty"`
	} `json:"location,omitempty"`
	Type *[]string `json:"type,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /business/{businessId})
	GetBusinessBusinessId(ctx echo.Context, businessId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetBusinessBusinessId converts echo context to params.
func (w *ServerInterfaceWrapper) GetBusinessBusinessId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "businessId" -------------
	var businessId string

	err = runtime.BindStyledParameter("simple", false, "businessId", ctx.Param("businessId"), &businessId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter businessId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBusinessBusinessId(ctx, businessId)
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

	router.GET(baseURL+"/business/:businessId", wrapper.GetBusinessBusinessId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/4xTTW/bMAz9Kwa3YxC53Wk6FgWGAMP+wFAMik3HLGRJo+gAQZD/PlC2k6b1gJ708aj3",
	"SD7qDE0cUgwYJIM9Q256HFzZPo2ZAuayTxwTshCW0x9qdZFTQrCQhSkc4LIB17Y8P/iA7We2X27A1YAW",
	"c8OUhGJYxykn7067wR3WCUiRIk6Cw3oW84Vjdic9+9i4RfG+Ru+EZGzXpXwMh/+hN5W4f8VGbrKfT+wj",
	"h15R6GJ5TOIVuxq0gSNyLmXAw7be1soYEwaXCCx829bbR9hActIXdbOYYc7LbtdeFDmg6HLnBfxAqRat",
	"SpPgoTSt2p+q3TMUKS43u3YKX6KfruxFnt2AgpzB/n6vsXuuYldJj9WbqkgRzRo2EMrcwP4tI+PfkRhb",
	"sMIjbubhXfPkRYNziiFP7j7WtS5NDIKhlOxS8jQNg3nN00Tc+L4ydmDhi7n9FjN/FXNNuJh0X9Zq2xiF",
	"CY/YVnlsGsy5G71X1wtBRj4uPRrZg4VeJFljdFZ9H7PYc4osF+MSmeMDvNf8qXHVRKOj4Zjc3k9168PJ",
	"4M6NXsDC97quVfrl8i8AAP//C0X0mwsEAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
