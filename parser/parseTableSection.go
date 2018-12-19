package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) tableSection(sec *Section) error {
	err := checkSection(sec, types.OrderTable)
	if err != nil {
		return err
	}

	rd := bytes.NewReader(sec.Data)
	//1. num tab ele
	var numTabEle uint32
	err = utils.DecodeVarInt(rd, 32, &numTabEle)
	if err != nil {
		return err
	}

	//2. tab eles
	for i := 0; i < int(numTabEle); i++ {
		tableType, err := DecodeTableType(rd)
		if err != nil {
			return err
		}
		p.Module.Tables.Defs = append(p.Module.Tables.Defs, types.TableDef{Type: tableType})
	}

	err = p.validateTable()
	return err
}
func (p *Parser) validateTable() error {
	logrus.Info("TODO: validateTable()")
	return nil
}
