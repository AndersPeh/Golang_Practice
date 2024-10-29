package folder

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
)

// Folder represents a folder with a name, path, and organization ID.
type Folder struct {
	Name  string    `json:"name"`
	Path  string    `json:"paths"` // Map "paths" from JSON to "Path" field
	OrgID uuid.UUID `json:"org_id"`
}

// Driver provides methods to manipulate folders.
type Driver struct {
	folders []Folder
}

// NewDriver creates a new Driver with the given folders.
func NewDriver(folders []Folder) *Driver {
	return &Driver{folders: folders}
}

// GetFoldersByOrgID returns all folders for a given organization ID.
func (d *Driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	var res []Folder
	for _, f := range d.folders {
		if f.OrgID == orgID {
			res = append(res, f)
		}
	}
	return res
}

// PrettyPrint prints folders in a readable JSON format.
func PrettyPrint(folders []Folder) {
	jsonBytes, err := json.MarshalIndent(folders, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
