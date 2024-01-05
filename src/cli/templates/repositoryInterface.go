package templates

const RepositoryInterface = `package {{ .PrivateNameSingular }}Repositories

import (
	"context"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// Interface{{ .PublicNameSingular }}Repository is the interface implemented by {{ .PublicNameSingular }}Repository
type Interface{{ .PublicNameSingular }}Repository interface {
	// List returns a slice of {{ .PublicNameSingular }} entities matching the given params
	List(ctx context.Context, params *{{ .PrivateNameSingular }}Models.QueryParams) ([]models.{{ .PublicNameSingular }}, error)
	// Read returns a single {{ .PublicNameSingular }} entity matching the params
	Read(ctx context.Context, params *{{ .PrivateNameSingular }}Models.QueryParams) (*models.{{ .PublicNameSingular }}, error)
	// Create creates a new {{ .PublicNameSingular }} entity
	Create(ctx context.Context, {{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}) error
	// Update updates a existing {{ .PublicNameSingular }} entity
	Update(ctx context.Context, {{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}) error
	// Delete (soft)deletes a existing {{ .PublicNameSingular }} entity
	Delete(ctx context.Context, {{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}) error
}
`
