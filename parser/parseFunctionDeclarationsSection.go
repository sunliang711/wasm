package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/types/IR"
	"wasm/utils"
)

func (p *Parser) functionDeclarationsSection(sec *Section) error {
	err := checkSection(sec, IR.OrderFunctionDeclarations)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)
	//1. num functions
	var numFun uint32
	_, err = utils.DecodeVarInt(rd, 32, &numFun)
	if err != nil {
		return err
	}

	<-p.typeParsed
	var funcTypeIndex uint32
	//2. function type index
	for i := 0; i < int(numFun); i++ {
		_, err = utils.DecodeVarInt(rd, 32, &funcTypeIndex)
		if err != nil {
			return err
		}
		if int(funcTypeIndex) > len(p.Module.Types) {
			return fmt.Errorf(types.ErrFunctionTypeIndexOutOfRange)
		}
		funcDef := IR.FunctionDef{
			Type:         IR.IndexedFunctionType{uint64(funcTypeIndex)},
			FunctionType: p.Module.Types[funcTypeIndex],
		}
		p.Module.Functions.Defs = append(p.Module.Functions.Defs, funcDef)
		logrus.Infof("<function Declaration section> function def: %v", funcDef)
	}
	p.funcDeclarationParsed <- struct{}{}

	err = p.validateFunctionDeclarations()
	return err
}
func (p *Parser) validateFunctionDeclarations() error {
	//TODO
	logrus.Info("TODO: validateFunctionDeclarations()")
	return nil
}
