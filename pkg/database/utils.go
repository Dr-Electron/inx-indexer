package database

import (
	"fmt"

	"github.com/iotaledger/hive.go/core/ioutils"
)

// Exists checks if the database folder exists and is not empty.
func Exists(dbPath string) (bool, error) {

	dirExists, isDirectory, err := ioutils.PathExists(dbPath)
	if err != nil {
		return false, fmt.Errorf("unable to check database path (%s): %w", dbPath, err)
	}
	if !dirExists {
		return false, nil
	}
	if !isDirectory {
		return false, fmt.Errorf("database path exists but is not a directory (%s)", dbPath)
	}

	// directory exists, but maybe database doesn't exist.
	// check if the directory is empty (needed for example in docker environments)
	dirEmpty, err := ioutils.DirectoryEmpty(dbPath)
	if err != nil {
		return false, fmt.Errorf("unable to check database path (%s): %w", dbPath, err)
	}

	return !dirEmpty, nil
}
