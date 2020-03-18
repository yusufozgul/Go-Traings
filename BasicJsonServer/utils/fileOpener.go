package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) ([]byte, error) {

	if len(fileName) != 0 {
		bytes, error := ioutil.ReadFile(fileName)

		if error != nil {
			return nil, error
		}

		return bytes, nil
	} else {
		return nil, errors.New("Empty File Name Error")
	}
}
