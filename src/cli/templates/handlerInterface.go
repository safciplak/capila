package templates

const HandlerInterface = `// code generated by ifacemaker; DO NOT EDIT.

package {{ .PrivateNameSingular }}Handlers

import (
	"github.com/gin-gonic/gin"
)

// Interface{{ .PublicNameSingular }}Handler is the interface implemented by {{ .PublicNameSingular }}Handler
type Interface{{ .PublicNameSingular }}Handler interface {
	// List returns a list of {{ .PublicNameSingular }} entities
	List() gin.HandlerFunc
	// Read returns a single {{ .PublicNameSingular }} entity
	Read() gin.HandlerFunc
	// Create creates a new {{ .PublicNameSingular }} entity
	Create() gin.HandlerFunc
	// Update updates a existing {{ .PublicNameSingular }} entity
	Update() gin.HandlerFunc
	// Delete (soft)deletes a existing {{ .PublicNameSingular }} entity
	Delete() gin.HandlerFunc
}
`
