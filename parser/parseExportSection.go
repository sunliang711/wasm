package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/types/IR"
	"wasm/utils"
)

func (p *Parser) exportSection(sec *Section) error {
	err := checkSection(sec, IR.OrderExport)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	//1. num export
	var numExport uint32
	_, err = utils.DecodeVarInt(rd, 32, &numExport)
	if err != nil {
		return err
	}

	buf := make([]byte, 1)
	//2. exports
	for i := 0; i < int(numExport); i++ {
		//name
		_, name, err := utils.ReadVarChars(rd)
		if err != nil {
			return err
		}
		err = utils.CheckUTF8(name)
		if err != nil {
			return err
		}

		//kind
		_, err = rd.Read(buf)
		if err != nil {
			return err
		}
		kind := buf[0]

		//index:pointer to  function index space[import ,then definition]
		var index uint32
		_, err = utils.DecodeVarInt(rd, 32, &index)
		if err != nil {
			return err
		}

		export := IR.Export{Name: string(name), Kind: IR.ExternKind(kind), Index: uint64(index)}
		p.Module.Exports = append(p.Module.Exports, &export)
		logrus.Infof("<export section> export: %v", export)
	}
	//add all exported functions to Module.ExportFunctions
	for i := range p.Module.Exports {
		if p.Module.Exports[i].Kind == IR.Function {
			p.Module.ExportFunctions = append(p.Module.ExportFunctions, p.Module.Exports[i])
		}
	}

	err = p.validateExport()
	return err
}

func (p *Parser) validateExport() error {
	//TODO
	logrus.Info("TODO: validateExport()")
	return nil
}
