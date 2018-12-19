package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) startSection(sec *Section) error {
	err := checkSection(sec, types.OrderStart)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	var startIndex uint32
	_, err = utils.DecodeVarInt(rd, 32, &startIndex)
	if err != nil {
		return err
	}
	p.Module.StartFunctionIndex = int(startIndex)

	logrus.Infof("<start section> start function index: %d", startIndex)

	err = p.validateStart()
	return err
}
func (p *Parser) validateStart() error {
	logrus.Info("TODO: validateStart()")
	return nil
}
