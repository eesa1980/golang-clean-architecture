//To create a Go file that recursively iterates over files in a folder and runs a `mockgen` command on them, you can use the `filepath` and `os/exec` packages. Here's an example:

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	root := "./pkg/application/common" // Replace with the root folder path

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // Skip directories
		}

		if !strings.Contains(path, "interface") {
			return nil
		}

		destination := "tests/mocks/interfaces/" + info.Name()

		// Run mockgen command on the file
		cmd := exec.Command("mockgen", "-source", path, "-destination", destination)
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to run mockgen command: %w", err)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

//Make sure to replace `/path/to/folder` with the actual path of the folder you want to iterate over. This code will walk through all the files in the folder (including subfolders) and run the `mockgen` command on each file. The generated mock files will be named with a prefix "mock_" followed by the original file name.
