package utils

import (
    "os"
    "path/filepath"
	"fmt"
)

func FindProjectRoot() (string, error) {
    dir, err := os.Getwd()
    if err != nil {
        return "", err
    }

    for {
        if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
            return dir, nil
        }
        parent := filepath.Dir(dir)
        if parent == dir {
            return "", fmt.Errorf("could not find go.mod file")
        }
        dir = parent
    }
}