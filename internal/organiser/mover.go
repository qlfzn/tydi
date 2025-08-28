package organiser

import (
	"log"
	"os"
	"path/filepath"
)

func (f *File) MoveFiles(baseDir string, fileGroup map[string][]os.DirEntry) error {
	for group, files := range fileGroup {
		destDir := filepath.Join(baseDir, group)

		err := os.MkdirAll(destDir, 0755)
		if err != nil {
			log.Printf("ERROR - create folder failed: %s", err)
			return err
		}

		log.Printf("INFO - checked folder: %s", destDir)

		for _, file := range files {
			srcPath := filepath.Join(baseDir, file.Name())
			destPath := filepath.Join(destDir, file.Name())

			if _, err := os.Stat(destPath); err == nil {
				log.Printf("WARN - skipped (exists): %s -> %s", srcPath, destPath)
				continue
			}

			if err := os.Rename(srcPath, destPath); err != nil {
				log.Printf("ERROR - move failed: %s -> %s (%v)", srcPath, destPath, err)
				continue
			}

			log.Printf("INFO - moved: %s -> %s", srcPath, destPath)
		}
	}

	return nil
}
