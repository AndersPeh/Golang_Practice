package folder

import (
	"github.com/gofrs/uuid"
)

// Folder represents a folder with a name, path, and organization ID.
type Folder struct {
	Name  string    `json:"name"`
	Path  string    `json:"path"`
	OrgID uuid.UUID `json:"org_id"`
}

// Driver holds all the folders and provides methods to manipulate them.
type Driver struct {
	folders []Folder
}

// NewDriver creates a new Driver instance with the provided folders.
func NewDriver(folders []Folder) *Driver {
	return &Driver{folders: folders}
}

// GetFoldersByOrgID returns all folders that belong to a specific organization.
func (f *Driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	var result []Folder
	for _, folder := range f.folders {
		if folder.OrgID == orgID {
			result = append(result, folder)
		}
	}
	return result
}
