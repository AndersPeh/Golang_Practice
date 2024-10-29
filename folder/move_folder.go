package folder

import (
	"errors"
	"strings"
)

// MoveFolder moves a folder from one parent to another within the same organization.
func (f *Driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Step 1: Find the source folder
	var sourceFolder *Folder
	for i := range f.folders {
		if f.folders[i].Name == name {
			sourceFolder = &f.folders[i]
			break
		}
	}
	if sourceFolder == nil {
		return nil, errors.New("Source folder does not exist")
	}

	// Step 2: Find the destination folder
	var destFolder *Folder
	for i := range f.folders {
		if f.folders[i].Name == dst {
			destFolder = &f.folders[i]
			break
		}
	}
	if destFolder == nil {
		return nil, errors.New("Destination folder does not exist")
	}

	// Step 3: Check if both folders are in the same organization
	if sourceFolder.OrgID != destFolder.OrgID {
		return nil, errors.New("Cannot move folder to a different organization")
	}

	// Step 4: Check for invalid moves
	if sourceFolder.Path == destFolder.Path {
		return nil, errors.New("Cannot move a folder to itself")
	}
	if strings.HasPrefix(destFolder.Path+".", sourceFolder.Path+".") {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}

	// Step 5: Update paths of source folder and its descendants
	oldPathPrefix := sourceFolder.Path
	newPathPrefix := destFolder.Path + "." + sourceFolder.Name

	for i := range f.folders {
		if f.folders[i].OrgID != sourceFolder.OrgID {
			continue
		}
		if f.folders[i].Path == sourceFolder.Path || strings.HasPrefix(f.folders[i].Path+".", oldPathPrefix+".") {
			// Update the path by replacing the old prefix with the new prefix
			suffix := strings.TrimPrefix(f.folders[i].Path, oldPathPrefix)
			f.folders[i].Path = newPathPrefix + suffix
		}
	}

	// Step 6: Return the updated list of folders
	return f.folders, nil
}
