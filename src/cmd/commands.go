package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func cleanOutputDir(dir string) error {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if err := os.Remove(path); err != nil {
				return err
			}
		}
		return nil
	})

	time.Sleep(2 * time.Second)

	return err
}
