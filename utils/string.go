package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

const (
	ErrInsufficientChar = "Cannot read enough characters from reader"
)

func ReadVarChars(rd io.Reader) ([]byte, error) {
	var numChar uint32
	err := DecodeVarInt(rd, 32, &numChar)
	if err != nil {
		return nil, err
	}

	chars := make([]byte, numChar)
	n, err := rd.Read(chars)
	if err != nil {
		return nil, err
	}
	if n != int(numChar) {
		return nil, fmt.Errorf(ErrInsufficientChar)
	}
	return chars, nil
}

func CheckUTF8(src []byte) error {
	logrus.Info("TODO: CheckUTF8()")
	return nil
}
