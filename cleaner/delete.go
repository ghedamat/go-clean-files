package cleaner

import (
	"os"
)

func DeleteFiles(files []File) error {
	var err error
	for _, f := range files {
		err = os.Remove(f.Path)
	}
	return err
}
