package files

import (
	"errors"
	"io/ioutil"
)

// Read file content of given path.
func Read(path string) (string, error) {
	if path == "" {
		return "", errors.New("given path is empty")
	}

	byt, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(byt), nil
}
