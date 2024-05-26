package cat

import (
	"errors"
	"os"
)

func Cat(file string) (string, error) {
	// If no filepath is given, return with error
	if len(file) == 0 {
		return "", errors.New("filepath required")
	}

	// Analyze file
	info, err := os.Stat(file)

	// Return with error if file is invalid
	if errors.Is(err, os.ErrNotExist) { // Does not exist
		return "", errors.New("file not found: " + file)

	} else if info.Mode().Perm()&0444 != 0444 { // Is not readable
		return "", errors.New("file not readable: " + file)

	} else if err != nil { // Unknown error
		return "", err
	}

	// Read & print
	fileContentByteSlice, err := os.ReadFile(file)

	// Validate reading file
	if err != nil {
		return "", err
	}

	// Return success
	return string(fileContentByteSlice), nil
}
