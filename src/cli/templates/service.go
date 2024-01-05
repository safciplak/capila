package templates

const Service = `//go:generate generate-interfaces.sh

package {{ .PrivateNameSingular }}Services

import (
	"context"

	"github.com/safciplak/capila/src/apm"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	{{ .PrivateNameSingular }}Repositories "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/repositories"
	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// {{ .PublicNameSingular }}Service contains the necessary repositories.
type {{ .PublicNameSingular }}Service struct {
	Repository {{ .PrivateNameSingular }}Repositories.Interface{{ .PublicNameSingular }}Repository
}

// New{{ .PublicNameSingular }}Service instantiates a new {{ .PublicNameSingular }}Service
func New{{ .PublicNameSingular }}Service(repository {{ .PrivateNameSingular }}Repositories.Interface{{ .PublicNameSingular }}Repository) Interface{{ .PublicNameSingular }}Service {
	return &{{ .PublicNameSingular }}Service{
		Repository: repository,
	}
}

// List returns a slice of {{ .PublicNameSingular }} entities matching the given params
func (service *{{ .PublicNameSingular }}Service) List(ctx context.Context, params *{{ .PrivateNameSingular }}Models.ListRequest) ([]models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Service.List", "service"))

	queryParams := params.ToQueryParams()

	return service.Repository.List(ctx, queryParams)
}

// Read returns a single {{ .PublicNameSingular }} entity matching the params
func (service *{{ .PublicNameSingular }}Service) Read(ctx context.Context, params *{{ .PrivateNameSingular }}Models.BaseRequest) (*models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Service.Read", "service"))

	queryParams := params.ToQueryParams()

	return service.Repository.Read(ctx, queryParams)
}

// Create creates a new {{ .PublicNameSingular }} entity
func (service *{{ .PublicNameSingular }}Service) Create(ctx context.Context, params *{{ .PrivateNameSingular }}Models.CreateRequest) (*models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Service.Create", "service"))

	{{ .PrivateNameSingular }} := params.ToModel()
	err := service.Repository.Create(ctx, {{ .PrivateNameSingular }})

	return {{ .PrivateNameSingular }}, err
}

// Update updates a existing {{ .PublicNameSingular }} entity
func (service *{{ .PublicNameSingular }}Service) Update(ctx context.Context, params *{{ .PrivateNameSingular }}Models.UpdateRequest) (*models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Service.Update", "service"))

	{{ .PrivateNameSingular }} := params.ToModel()
	err := service.Repository.Update(ctx, {{ .PrivateNameSingular }})

	return {{ .PrivateNameSingular }}, err
}

// Delete (soft)deletes a existing {{ .PublicNameSingular }} entity
func (service *{{ .PublicNameSingular }}Service) Delete(ctx context.Context, params *{{ .PrivateNameSingular }}Models.BaseRequest) (*models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Service.Delete", "service"))

	{{ .PrivateNameSingular }} := params.ToModel(true)
	err := service.Repository.Delete(ctx, {{ .PrivateNameSingular }})

	return {{ .PrivateNameSingular }}, err
}`
