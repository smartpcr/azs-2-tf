package utils

import (
	"fmt"
	"os"
)

func EnsureDirectory(folderPath string) error {
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return err
	}

	folderInfo, err := os.Stat(folderPath)
	if err != nil {
		fmt.Println("Error getting folder info:", err)
		return err
	}

	// Verify if the path is a directory
	if !folderInfo.IsDir() {
		fmt.Println("Error: path is not a directory")
		return err
	}

	return nil
}
