package parser

import (
	"bytes"
	"io"
	"io/ioutil"
)

func LoadBinary(filename string) (io.Reader, error) {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(fileContents), nil
}
