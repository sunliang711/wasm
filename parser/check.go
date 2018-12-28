package parser

import (
	"fmt"
	"io"
	"wasm/types"
	"wasm/types/IR"
)

func checkConstant(rd io.Reader, constant []byte, errMsg string) error {
	numBytes := len(constant)
	buf := make([]byte, numBytes)
	nRead, err := rd.Read(buf)
	if err != nil {
		return err
	}
	if nRead != numBytes {
		return fmt.Errorf(types.ErrReadCount)
	}
	for i := 0; i < numBytes; i++ {
		if buf[i] != constant[i] {
			return fmt.Errorf(errMsg)
		}
	}
	return nil
}

func checkSection(sec *Section, orderSec byte) error {
	if len(sec.Data) != int(sec.NumSectionBytes) {
		return fmt.Errorf(types.ErrSectionNum, IR.OrderSectionString(orderSec))
	}
	if orderSec != sec.Type {
		return fmt.Errorf(types.ErrSectionType, IR.OrderSectionString(orderSec))
	}
	return nil
}
