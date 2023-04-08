package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Splitn the PATH env vairible a slice of dir
	dirs := filepath.SplitList(os.Getenv("PATH"))

	// Get the command=line arguments as slice of lowercase strings
	args := make([]string, len(os.Args)-1)
	for i := 1; i <  len(os.Args); i++ {
		args[i-1] = strings.ToLower(os.Args[i])
	}

	// Search for the aguments in each dir of the PATH
	for _, dir := range dirs{
		files, err := os.ReadDir(dir)
		if err != nil {
			continue // Skip dir that cannot be read
		}
		for _, file := range files {
			name := strings.ToLower(file.Name())
			for _, arg := range args {
				if strings.Contains(name, arg) {
					fmt.Printf("%s/%s\n", dir, file.Name())
				}
			}
		}
	}

}