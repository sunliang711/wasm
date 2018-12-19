package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"wasm/types"
)

func Bytes2float32(src []byte, isLittleEndian bool) (float32, error) {
	var result float32
	var err error
	buf := bytes.NewReader(src)
	if isLittleEndian {
		err = binary.Read(buf, binary.LittleEndian, &result)
	} else {
		err = binary.Read(buf, binary.BigEndian, &result)
	}
	if err != nil {
		return 0.0, fmt.Errorf(types.ErrInvalidFloat32Format)
	}
	return result, nil
}

func Bytes2float64(src []byte, isLittleEndian bool) (float64, error) {
	var result float64
	var err error
	buf := bytes.NewReader(src)
	if isLittleEndian {
		err = binary.Read(buf, binary.LittleEndian, &result)
	} else {
		err = binary.Read(buf, binary.BigEndian, &result)
	}
	if err != nil {
		return 0.0, fmt.Errorf(types.ErrInvalidFloat64Format)
	}
	return result, nil
}
