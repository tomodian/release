package files

import (
	"errors"
	"io/ioutil"
	"os"
)

// Update overwrites the existing file while retaining the original file permission.
func Update(path, doc string) error {

	if path == "" || doc == "" {
		return errors.New("given input is empty")
	}

	info, err := os.Stat(path)

	if err != nil {
		return nil
	}

	if err := ioutil.WriteFile(path, []byte(doc), info.Mode()); err != nil {
		return err
	}

	return nil
}
