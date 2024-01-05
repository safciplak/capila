package templates

const Handler = `//go:generate generate-interfaces.sh

package {{ .PrivateNameSingular }}Handlers

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/convert"
	capilaContext "github.com/safciplak/capila/src/http/context"
	"github.com/safciplak/capila/src/http/handlers"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	{{ .PrivateNameSingular }}Services "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/services"
)

// {{ .PublicNameSingular }}Handler is the receiver for the handler functions
type {{ .PublicNameSingular }}Handler struct {
	service {{ .PrivateNameSingular }}Services.Interface{{ .PublicNameSingular }}Service
}

// New{{ .PublicNameSingular }}Handler initializes the handler.
func New{{ .PublicNameSingular }}Handler(service {{ .PrivateNameSingular }}Services.Interface{{ .PublicNameSingular }}Service) Interface{{ .PublicNameSingular }}Handler {
	return &{{ .PublicNameSingular }}Handler{
		service,
	}
}

// List returns a list of {{ .PublicNameSingular }} entities
func (handler *{{ .PublicNameSingular }}Handler) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &{{ .PrivateNameSingular }}Models.ListRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.List(ctx, request)
		})
	}
}

// Read returns a single {{ .PublicNameSingular }} entity
func (handler *{{ .PublicNameSingular }}Handler) Read() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &{{ .PrivateNameSingular }}Models.BaseRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Read(ctx, request)
		})
	}
}

// Create creates a new {{ .PublicNameSingular }} entity
func (handler *{{ .PublicNameSingular }}Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &{{ .PrivateNameSingular }}Models.CreateRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Create(ctx, request)
		})
	}
}

// Update updates a existing {{ .PublicNameSingular }} entity
func (handler *{{ .PublicNameSingular }}Handler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &{{ .PrivateNameSingular }}Models.UpdateRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Update(ctx, request)
		})
	}
}

// Delete (soft)deletes a existing {{ .PublicNameSingular }} entity
func (handler *{{ .PublicNameSingular }}Handler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &{{ .PrivateNameSingular }}Models.BaseRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Delete(ctx, request)
		})
	}
}`
