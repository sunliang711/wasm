package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/core/IR"
	"wasm/utils"
)

func (p *Parser) tableSection(sec *Section) error {
	err := checkSection(sec, IR.OrderTable)
	if err != nil {
		return err
	}

	rd := bytes.NewReader(sec.Data)
	//1. num tab ele
	var numTabEle uint32
	_, err = utils.DecodeVarInt(rd, 32, &numTabEle)
	if err != nil {
		return err
	}

	//2. tab eles
	for i := 0; i < int(numTabEle); i++ {
		tableType, err := DecodeTableType(rd)
		if err != nil {
			return err
		}
		p.Module.Tables.Defs = append(p.Module.Tables.Defs, IR.TableDef{Type: tableType})
	}

	err = p.validateTable()
	return err
}
func (p *Parser) validateTable() error {
	//TODO
	logrus.Info("TODO: validateTable()")
	return nil
}
