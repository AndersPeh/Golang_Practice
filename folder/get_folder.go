package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

// GetAllChildFolders returns all child folders of a given folder within an organization.
func (f *Driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Step 1: Find the target folder by name and orgID
	var targetFolder *Folder
	for _, folder := range f.folders {
		if folder.OrgID == orgID && folder.Name == name {
			targetFolder = &folder
			break
		}
	}
	if targetFolder == nil {
		return nil, errors.New("Folder does not exist in the specified organization")
	}

	// Step 2: Collect all child folders
	var childFolders []Folder
	targetPathPrefix := targetFolder.Path + "."

	for _, folder := range f.folders {
		if folder.OrgID != orgID {
			continue
		}
		if strings.HasPrefix(folder.Path, targetPathPrefix) {
			childFolders = append(childFolders, folder)
		}
	}

	// Step 3: Return the list of child folders
	return childFolders, nil
}
