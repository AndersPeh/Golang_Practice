package folder

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestGetAllChildFolders(t *testing.T) {
	// Initialize sample organization IDs
	orgID1 := uuid.FromStringOrNil(DefaultOrgID)
	orgID2 := uuid.Must(uuid.NewV4())

	// Create sample folders
	folders := []Folder{
		{Name: "alpha", Path: "alpha", OrgID: orgID1},
		{Name: "bravo", Path: "alpha.bravo", OrgID: orgID1},
		{Name: "charlie", Path: "alpha.bravo.charlie", OrgID: orgID1},
		{Name: "delta", Path: "alpha.delta", OrgID: orgID1},
		{Name: "echo", Path: "echo", OrgID: orgID1},
		{Name: "foxtrot", Path: "foxtrot", OrgID: orgID2},
	}

	// Initialize the driver with the folders
	driver := NewDriver(folders)

	// Define test cases
	tests := []struct {
		orgID       uuid.UUID
		folderName  string
		expected    []Folder
		expectError bool
	}{
		{
			orgID:      orgID1,
			folderName: "alpha",
			expected: []Folder{
				{Name: "bravo", Path: "alpha.bravo", OrgID: orgID1},
				{Name: "charlie", Path: "alpha.bravo.charlie", OrgID: orgID1},
				{Name: "delta", Path: "alpha.delta", OrgID: orgID1},
			},
			expectError: false,
		},
		{
			orgID:      orgID1,
			folderName: "bravo",
			expected: []Folder{
				{Name: "charlie", Path: "alpha.bravo.charlie", OrgID: orgID1},
			},
			expectError: false,
		},
		{
			orgID:       orgID1,
			folderName:  "charlie",
			expected:    []Folder{},
			expectError: false,
		},
		{
			orgID:       orgID1,
			folderName:  "echo",
			expected:    []Folder{},
			expectError: false,
		},
		{
			orgID:       orgID1,
			folderName:  "invalid_folder",
			expected:    nil,
			expectError: true,
		},
		{
			orgID:       orgID1,
			folderName:  "foxtrot",
			expected:    nil,
			expectError: true,
		},
	}

	// Iterate over test cases
	for _, test := range tests {
		result, err := driver.GetAllChildFolders(test.orgID, test.folderName)
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for folder '%s', but got none", test.folderName)
			}
			continue
		}

		if err != nil {
			t.Errorf("Did not expect error for folder '%s', but got: %v", test.folderName, err)
			continue
		}

		if len(result) != len(test.expected) {
			t.Errorf("For folder '%s', expected %d child folders, got %d", test.folderName, len(test.expected), len(result))
			continue
		}

		// Check that the expected folders are in the result
		expectedPaths := make(map[string]bool)
		for _, folder := range test.expected {
			expectedPaths[folder.Path] = true
		}

		for _, folder := range result {
			if !expectedPaths[folder.Path] {
				t.Errorf("Unexpected folder '%s' in result for '%s'", folder.Path, test.folderName)
			}
		}
	}
}
