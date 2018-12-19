package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"wasm/types"
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

func ReadByte(rd io.Reader) (byte, error) {
	buf := make([]byte, 1)
	_, err := rd.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func ReadNByte(rd io.Reader, n int) ([]byte, error) {
	if n <= 0 {
		return nil, fmt.Errorf(types.ErrInvalidParameter)
	}
	buf := make([]byte, n)
	nRead, err := rd.Read(buf)
	if err != nil {
		return nil, err
	}
	if nRead != n {
		return nil, fmt.Errorf(types.ErrInsufficientBytes)
	}

	return buf, nil
}
