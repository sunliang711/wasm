package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) userSection(sec *Section) error {
	err := checkSection(sec, types.OrderUser)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)
	userSec := types.UserSection{}
	//1. name
	nameUsedBytes, nameBytes, err := utils.ReadVarChars(rd)
	if err != nil {
		return err
	}
	err = utils.CheckUTF8(nameBytes)
	if err != nil {
		return err
	}
	userSec.Name = string(nameBytes)

	//2. data
	dataLen := int(sec.NumSectionBytes) - nameUsedBytes

	data, err := utils.ReadNByte(rd, dataLen)
	if err != nil {
		return err
	}
	userSec.Data = data

	p.Module.UserSections = append(p.Module.UserSections, userSec)
	logrus.Infof("<user section> user section: %v", userSec)

	return nil
}
