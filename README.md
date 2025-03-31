# Folder Management in Go

## Overview

This project provides a Go-based implementation for managing a hierarchical folder structure using `ltree`-like paths. It supports retrieving child folders and moving folders while maintaining the integrity of the structure. The project is designed to simulate a site directory and document organization system within a PostgreSQL-based platform.

## Features

- **Get All Child Folders**: Retrieve all child folders of a given folder within an organization.
- **Move Folders**: Move a folder and its subtree from one parent to another while maintaining the correct path hierarchy.
- **Error Handling**: Handles invalid folder paths, organization mismatches, and attempts to move a folder to an invalid location.

## Folder Structure

```
| go.mod
| README.md
| main.go
| folder
    | get_folder.go
    | get_folder_test.go
    | move_folder.go
    | static.go
    | sample.json
```

## Getting Started

### 1. Clone the Repository

```sh
git clone https://github.com/yourusername/your-repo.git
cd your-repo
```

### 2. Run the Sample Data Generator

To generate sample data and test the implementation, run the following command:

```sh
go run main.go
```

## Implementation Details

### Component 1: Get All Child Folders

#### Function: `GetAllChildFolders(orgID string, folderName string) ([]Folder, error)`

- Returns all child folders of a specified folder within an organization.
- Ensures that the requested folder exists and belongs to the specified organization.
- Uses an `ltree`-like path structure to determine hierarchical relationships.

#### Example Usage

```go
folders := GetAllChildFolders("org1", "alpha")
```

**Expected Output:**

```json
[
  {"name": "bravo", "path": "alpha.bravo", "orgID": "org1"},
  {"name": "charlie", "path": "alpha.bravo.charlie", "orgID": "org1"},
  {"name": "delta", "path": "alpha.delta", "orgID": "org1"}
]
```

### Component 2: Move Folders

#### Function: `MoveFolder(source string, destination string) error`

- Moves a folder and its entire subtree to a new parent folder.
- Validates the source and destination folders.
- Prevents moving a folder to itself or one of its children.
- Ensures that folders remain within the same organization.

#### Example Usage

```go
err := MoveFolder("bravo", "delta")
```

**Expected Output:**

```json
[
  {"name": "alpha", "path": "alpha", "orgID": "org1"},
  {"name": "bravo", "path": "alpha.delta.bravo", "orgID": "org1"},
  {"name": "charlie", "path": "alpha.delta.bravo.charlie", "orgID": "org1"},
  {"name": "delta", "path": "alpha.delta", "orgID": "org1"}
]
```

## Running Tests

To execute unit tests for `get_folder.go` and `move_folder.go`, run:

```sh
go test ./folder
```

##
