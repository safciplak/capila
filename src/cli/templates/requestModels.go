package templates

const RequestModels = `package {{ .PrivateNameSingular }}Models

import (
	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/convert"

	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// QueryParams contains the allowed query params per request.
type QueryParams struct {
	GUID     string
	Language string
}

// BaseRequest is the validator for default request param check.
type BaseRequest struct {
	Language *string {{ .LanguageBinding }}
	GUID     string  {{ .BaseRequestGUIDBinding }}
}

// Validate makes sure the correct data is being submitted
func (request *BaseRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindUri(request)
}

// ToQueryParams transforms a request to QueryParams
func (request *BaseRequest) ToQueryParams() *QueryParams {
	return &QueryParams{
		GUID:     request.GUID,
		Language: convert.PointerToString(request.Language),
	}
}

// ToModel transforms a request to a model
func (request *BaseRequest) ToModel(isDeleted bool) *models.{{ .PublicNameSingular }} {
	return &models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID:      request.GUID,
			IsDeleted: isDeleted,
		},
	}
}

// ListRequest is the validator for list (getAll) requests.
type ListRequest struct {
	Language *string {{ .LanguageBinding }}
}

// Validate makes sure the correct data is being submitted
func (request *ListRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindQuery(request)
}

// ToQueryParams transforms a request to QueryParams
func (request *ListRequest) ToQueryParams() *QueryParams {
	return &QueryParams{
		Language: convert.PointerToString(request.Language),
	}
}

// CreateRequest is the validator for create requests.
type CreateRequest struct {
	Language *string {{ .LanguageBinding }}
	GUID     string  {{ .CreateRequestGUIDBinding }}
}

// Validate makes sure the correct data is being submitted
func (request *CreateRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(request)
}

// ToModel transforms a request to a model
func (request *CreateRequest) ToModel() *models.{{ .PublicNameSingular }} {
	return &models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID: request.GUID,
		},
	}
}

// UpdateRequest is the validator for update requests.
type UpdateRequest struct {
	Language *string {{ .LanguageBinding }}
	GUID     string  {{ .UpdateRequestGUIDBinding }}
}

// Validate makes sure the correct data is being submitted
func (request *UpdateRequest) Validate(ctx *gin.Context) error {
	request.GUID = ctx.Params.ByName("guid")

	return ctx.ShouldBindJSON(request)
}

// ToModel transforms a request to a model
func (request *UpdateRequest) ToModel() *models.{{ .PublicNameSingular }} {
	return &models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID: request.GUID,
		},
	}
}`
