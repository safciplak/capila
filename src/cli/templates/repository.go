package templates

const Repository = `//go:generate generate-interfaces.sh

package {{ .PrivateNameSingular }}Repositories

import (
	"context"

	"github.com/safciplak/capila/src/apm"
	"github.com/safciplak/capila/src/database"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// {{ .PublicNameSingular }}Repository is used for future purposes as a receiver
type {{ .PublicNameSingular }}Repository struct {
	database *database.Connection
}

// New{{ .PublicNameSingular }}Repository initializes the service. database *database.Connection
func New{{ .PublicNameSingular }}Repository(db *database.Connection) Interface{{ .PublicNameSingular }}Repository {
	return &{{ .PublicNameSingular }}Repository{
		database: db,
	}
}

// List returns a slice of {{ .PublicNameSingular }} entities matching the given params
func (repo *{{ .PublicNameSingular }}Repository) List(ctx context.Context, params *{{ .PrivateNameSingular }}Models.QueryParams) ([]models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Repository.List", "repository"))

	{{ .PrivateNamePlural }} := make([]models.{{ .PublicNameSingular }}, 0)
	err := repo.database.Read.ModelContext(ctx, &{{ .PrivateNamePlural }}).
		Where("{{ .PrivateNameSingular }}.isdeleted = false").
		Select()

	return {{ .PrivateNamePlural }}, err
}

// Read returns a single {{ .PublicNameSingular }} entity matching the params
func (repo *{{ .PublicNameSingular }}Repository) Read(ctx context.Context, params *{{ .PrivateNameSingular }}Models.QueryParams) (*models.{{ .PublicNameSingular }}, error) {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Repository.Read", "repository"))

	{{ .PrivateNameSingular }} := new(models.{{ .PublicNameSingular }})
	err := repo.database.Read.ModelContext(ctx, {{ .PrivateNameSingular }}).
		Where("{{ .PrivateNameSingular }}.guid = ?", params.GUID).
		Where("{{ .PrivateNameSingular }}.isdeleted = false").
		First()

	return {{ .PrivateNameSingular }}, err
}

// Create creates a new {{ .PublicNameSingular }} entity
func (repo *{{ .PublicNameSingular }}Repository) Create(ctx context.Context, {{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}) error {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Repository.Create", "repository"))

	_, err := repo.database.Write.ModelContext(ctx, {{ .PrivateNameSingular }}).
		Returning({{ .PrivateNameSingular }}.BaseTableModel.Returning()).
		Insert()

	return err
}

// Update updates a existing {{ .PublicNameSingular }} entity
func (repo *{{ .PublicNameSingular }}Repository) Update(ctx context.Context, {{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}) error {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Repository.Update", "repository"))

	_, err := repo.database.Write.ModelContext(ctx, {{ .PrivateNameSingular }}).
		Column(""). // @TODO: select columns to update
		Returning({{ .PrivateNameSingular }}.BaseTableModel.Returning()).
		Where("{{ .PrivateNameSingular }}.guid = ?", {{ .PrivateNameSingular }}.GUID).
		Where("{{ .PrivateNameSingular }}.isdeleted = false").
		Update()

	return err
}

// Delete (soft)deletes a existing {{ .PublicNameSingular }} entity
func (repo *{{ .PublicNameSingular }}Repository) Delete(ctx context.Context, {{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}) error {
	defer apm.End(apm.Start(ctx, "{{ .PublicNameSingular }}Repository.Delete", "repository"))

	_, err := repo.database.Write.ModelContext(ctx, {{ .PrivateNameSingular }}).
		Column("isdeleted").
		Returning({{ .PrivateNameSingular }}.BaseTableModel.Returning()).
		Where("{{ .PrivateNameSingular }}.guid = ?", {{ .PrivateNameSingular }}.GUID).
		Where("{{ .PrivateNameSingular }}.isdeleted = false").
		Update()

	return err
}
`
