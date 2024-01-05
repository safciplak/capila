package templates

const Models = `package models

// {{ .PublicNameSingular }} is the base model for {{ .PrivateNameSingular }} records.
type {{ .PublicNameSingular }} struct {
	BaseTableModel
}

// Columns returns the columns used by the model
func ({{ .PublicNameSingular }}) Columns() []string {
	return []string{"guid", "name"}
}
`
