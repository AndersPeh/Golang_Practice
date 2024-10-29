package main

import (
	"fmt"

	"github.com/AndersPeh/Anders_sc/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// Use the default OrgID from static.go
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// Generate sample data or read from sample.json
	res := folder.GetSampleData() // Or folder.GenerateData() to generate new data

	// Initialize the driver with the folders
	folderDriver := folder.NewDriver(res)

	// Get folders by OrgID
	orgFolders := folderDriver.GetFoldersByOrgID(orgID)
	fmt.Printf("Folders for OrgID: %s\n", orgID)
	folder.PrettyPrint(orgFolders)

	// Example usage of GetAllChildFolders
	childFolders, err := folderDriver.GetAllChildFolders(orgID, "alpha")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nChild folders of 'alpha':")
		folder.PrettyPrint(childFolders)
	}

	// Example usage of MoveFolder
	updatedFolders, err := folderDriver.MoveFolder("bravo", "delta")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFolders after moving 'bravo' under 'delta':")
		folder.PrettyPrint(updatedFolders)
	}
}
