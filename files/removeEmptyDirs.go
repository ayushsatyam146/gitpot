package files

import (
	"os"
	"path/filepath"
)

// isDirEmpty checks if a directory is empty.
func isDirEmpty(path string) (bool, error) {
	dir, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	_, err = dir.Readdir(1)
	if err == nil {
		return false, nil
	}

	return true, nil
}


//  recursively removes all the empty directories inside path
func RemoveEmptyDirs(rootPath string) error {
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == rootPath {
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if info.IsDir() {
			isEmpty, err := isDirEmpty(path)
			if err != nil {
				return err
			}

			if isEmpty {
				return os.Remove(path)
			}
		}

		return nil
	})
}