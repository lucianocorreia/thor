package thor

import (
	"os"
)

// CreateDirIfNotExists creates the directory is it not exists
func (t *Thor) CreateDirIfNotExists(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}
