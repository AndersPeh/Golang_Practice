package folder

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestMoveFolder(t *testing.T) {
	// Initialize sample organization IDs
	orgID1 := uuid.FromStringOrNil(DefaultOrgID)
	orgID2 := uuid.Must(uuid.NewV4())

	// Create sample folders
	folders := []Folder{
		{Name: "alpha", Path: "alpha", OrgID: orgID1},
		{Name: "bravo", Path: "alpha.bravo", OrgID: orgID1},
		{Name: "charlie", Path: "alpha.bravo.charlie", OrgID: orgID1},
		{Name: "delta", Path: "alpha.delta", OrgID: orgID1},
		{Name: "echo", Path: "alpha.delta.echo", OrgID: orgID1},
		{Name: "foxtrot", Path: "foxtrot", OrgID: orgID2},
		{Name: "golf", Path: "golf", OrgID: orgID1},
	}

	// Define test cases
	tests := []struct {
		name        string
		source      string
		destination string
		expectError bool
		expected    []Folder
	}{
		{
			name:        "Move bravo under delta",
			source:      "bravo",
			destination: "delta",
			expectError: false,
			expected: []Folder{
				{Name: "alpha", Path: "alpha", OrgID: orgID1},
				{Name: "bravo", Path: "alpha.delta.bravo", OrgID: orgID1},
				{Name: "charlie", Path: "alpha.delta.bravo.charlie", OrgID: orgID1},
				{Name: "delta", Path: "alpha.delta", OrgID: orgID1},
				{Name: "echo", Path: "alpha.delta.echo", OrgID: orgID1},
				{Name: "foxtrot", Path: "foxtrot", OrgID: orgID2},
				{Name: "golf", Path: "golf", OrgID: orgID1},
			},
		},
		{
			name:        "Move bravo under golf",
			source:      "bravo",
			destination: "golf",
			expectError: false,
			expected: []Folder{
				{Name: "alpha", Path: "alpha", OrgID: orgID1},
				{Name: "bravo", Path: "golf.bravo", OrgID: orgID1},
				{Name: "charlie", Path: "golf.bravo.charlie", OrgID: orgID1},
				{Name: "delta", Path: "alpha.delta", OrgID: orgID1},
				{Name: "echo", Path: "alpha.delta.echo", OrgID: orgID1},
				{Name: "foxtrot", Path: "foxtrot", OrgID: orgID2},
				{Name: "golf", Path: "golf", OrgID: orgID1},
			},
		},
		{
			name:        "Move bravo under charlie (invalid)",
			source:      "bravo",
			destination: "charlie",
			expectError: true,
		},
		{
			name:        "Move bravo under bravo (invalid)",
			source:      "bravo",
			destination: "bravo",
			expectError: true,
		},
		{
			name:        "Move bravo under foxtrot (different org)",
			source:      "bravo",
			destination: "foxtrot",
			expectError: true,
		},
		{
			name:        "Move invalid_folder under delta",
			source:      "invalid_folder",
			destination: "delta",
			expectError: true,
		},
		{
			name:        "Move bravo under invalid_folder",
			source:      "bravo",
			destination: "invalid_folder",
			expectError: true,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Make a copy of the folders to prevent modifying the original data
			foldersCopy := make([]Folder, len(folders))
			copy(foldersCopy, folders)

			// Initialize the driver with the copied folders
			driver := NewDriver(foldersCopy)
			updatedFolders, err := driver.MoveFolder(test.source, test.destination)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Did not expect error but got: %v", err)
				return
			}

			// Compare the updated folders with the expected folders
			if !compareFolders(updatedFolders, test.expected) {
				t.Errorf("Folders after move do not match expected result")
				t.Logf("Expected: %+v", test.expected)
				t.Logf("Got: %+v", updatedFolders)
			}
		})
	}
}

// compareFolders compares two slices of folders for equality.
func compareFolders(a, b []Folder) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps for easier comparison
	mapA := make(map[string]Folder)
	mapB := make(map[string]Folder)

	for _, folder := range a {
		mapA[folder.Name] = folder
	}
	for _, folder := range b {
		mapB[folder.Name] = folder
	}

	for name, folderA := range mapA {
		folderB, exists := mapB[name]
		if !exists {
			return false
		}
		if folderA.Path != folderB.Path || folderA.OrgID != folderB.OrgID {
			return false
		}
	}

	return true
}
