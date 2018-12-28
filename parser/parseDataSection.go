package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
	"wasm/types/IR"
)

func (p *Parser) dataSection(sec *Section) error {
	err := checkSection(sec, IR.OrderData)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)
	//1. num data
	var numData uint32
	_, err = utils.DecodeVarInt(rd, 32, &numData)
	if err != nil {
		return err
	}
	//2. data items
	for i := 0; i < int(numData); i++ {
		dataSeg := IR.DataSegment{}
		flags, err := utils.ReadByte(rd)
		if err != nil {
			return err
		}
		switch flags {
		case 0:
			dataSeg.IsActive = true
			dataSeg.MemoryIndex = 0
			initExpression, err := DecodeInitializer(rd)
			if err != nil {
				return err
			}
			dataSeg.BaseOffset = &initExpression
		case 1:
			dataSeg.IsActive = false
			dataSeg.MemoryIndex = types.UINT64_MAX
			dataSeg.BaseOffset = &IR.InitializerExpression{}
		case 2:
			dataSeg.IsActive = true
			var memIndex uint32
			_, err = utils.DecodeVarInt(rd, 32, &memIndex)
			if err != nil {
				return err
			}
			dataSeg.MemoryIndex = uint64(memIndex)
			initExpression, err := DecodeInitializer(rd)
			if err != nil {
				return err
			}
			dataSeg.BaseOffset = &initExpression
		default:
			return fmt.Errorf(types.ErrInvalidDataSegFlags)
		}
		p.Module.DataSegments = append(p.Module.DataSegments, dataSeg)
		logrus.Infof("<data section> data segment: %v", dataSeg)
	}
	return nil
}

//decodeXXXImm
