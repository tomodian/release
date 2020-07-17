package files

import (
	"errors"
	"io/ioutil"
	"os"
)

// Read returns string of target path.
func Read(path string) (string, error) {
	if path == "" {
		return "", errors.New("given path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", errors.New("file does not exist")
	}

	byt, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(byt), nil
}
