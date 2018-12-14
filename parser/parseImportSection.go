package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)


func (p *Parser) importSection(sec *Section) error {
	err := checkSection(sec, types.OrderImport)
	if err != nil {
		return err
	}

	rd := bytes.NewReader(sec.Data)
	importSize, err := utils.DecodeUInt32(rd)
	if err != nil {
		return err
	}
	bufKind := make([]byte, 1)
	for i := 0; i < int(importSize); i++ {
		//1. mode name
		// 1.1 num char
		// 1.2 chars
		modeName, err := utils.ReadVarChars(rd)
		if err != nil {
			return err
		}
		//2. export name
		// 2.1 num char
		// 2.2 chars
		exportName, err := utils.ReadVarChars(rd)
		if err != nil {
			return err
		}
		//TODO modeName,exportName must utf8
		logrus.Infof("modeName: %s,exportName: %s", modeName, exportName)

		//3. extern kind(1 byte,native value)
		n,err := rd.Read(bufKind)
		if err != nil{
			return err
		}
		if n != len(bufKind){
			return fmt.Errorf(utils.ErrInsufficientChar)
		}
		kind := bufKind[0]

		//4. switch kind
		switch kind{

		}

	}

	err = p.validateImport()
	return err
}

func (p *Parser) validateImport() error {
	logrus.Info("TODO: validateImport()")
	return nil
}
