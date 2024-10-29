package folder

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gofrs/uuid"
	"github.com/lucasepe/codename"
)

// These are all helper methods and fixed types.
// There's no real need for you to be editing these, but feel free to tweak it to suit your needs.
// If you do make changes here, be ready to discuss why these changes were made.

// how many trees you want to generate
const MaxRootSet = 4

// maximum possible children per node
const MaxChild = 4

// max depth of the tree
const MaxDepth = 5

// the default orgID that we will be using for testing
const DefaultOrgID = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"

// GenerateData generates a random set of folders.
func GenerateData() []Folder {
	rng, _ := codename.DefaultRNG()
	tree := []Folder{}

	for i := 0; i < MaxRootSet; i++ {
		orgID := uuid.FromStringOrNil(DefaultOrgID)
		if i%3 == 0 {
			orgID = uuid.Must(uuid.NewV4())
		}

		name := codename.Generate(rng, 0)

		subtree := generateTree(1, []Folder{
			{
				Name:  name,
				OrgID: orgID,
				Path:  name,
			},
		})
		tree = append(tree, subtree...)
	}

	return tree
}

func generateTree(depth int, tree []Folder) []Folder {
	rng, _ := codename.DefaultRNG()

	if depth >= MaxDepth {
		return tree
	}

	var newTree []Folder
	for _, t := range tree {
		numOfChild := rng.Int()%MaxChild + 1
		for i := 0; i < numOfChild; i++ {
			name := codename.Generate(rng, 0)
			child := Folder{
				Name:  name,
				OrgID: t.OrgID,
				Path:  t.Path + "." + name,
			}
			newTree = append(newTree, child)
			// Recursively generate child trees
			subtree := generateTree(depth+1, []Folder{child})
			newTree = append(newTree, subtree...)
		}
	}

	return newTree
}

// MarshalJson marshals data into pretty JSON format.
func MarshalJson(b interface{}) []byte {
	s, _ := json.MarshalIndent(b, "", "\t")
	return s
}

// PrettyPrint prints data in a pretty JSON format.
func PrettyPrint(b interface{}) {
	s := MarshalJson(b)
	fmt.Println(string(s))
}

// GetSampleData reads folders from sample.json.
func GetSampleData() []Folder {
	_, filename, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "sample.json")

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	folders := []Folder{}
	err = json.Unmarshal(jsonByte, &folders)
	if err != nil {
		panic(err)
	}

	return folders
}

// WriteSampleData writes data to sample.json.
func WriteSampleData(data interface{}) {
	b := MarshalJson(data)
	_, filename, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "sample.json")

	err := os.WriteFile(filePath, b, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
