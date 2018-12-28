package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/types/IR"
	"wasm/utils"
)

func (p *Parser) memorySection(sec *Section) error {
	err := checkSection(sec, IR.OrderMemory)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	// 1. num memory ele
	var numMemory uint32
	_, err = utils.DecodeVarInt(rd, 32, &numMemory)
	if err != nil {
		return err
	}

	// 2. memory eles
	for i := 0; i < int(numMemory); i++ {
		isShared, min, max, err := DecodeFlags(rd)
		if err != nil {
			return err
		}
		mType := IR.MemoryType{IsShared: isShared, Size: IR.SizeConstraints{min, max}}
		mDef := IR.MemoryDef{Type: mType}
		p.Module.Memories.Defs = append(p.Module.Memories.Defs, mDef)
		logrus.Infof("<memory section> memory def: %v", mDef)
	}
	err = p.validateMemory()
	return err
}

func (p *Parser) validateMemory() error {
	//TODO
	logrus.Info("TODO: validateMemory()")
	return nil
}
