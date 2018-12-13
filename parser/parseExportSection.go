package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) exportSection(sec *Section) error {
	var (
		err error
	)
	logrus.Info("TODO exportSection()")

	err = p.validateExport()
	return err
}

func (p *Parser) validateExport() error {
	logrus.Info("TODO: validateExport()")
	return nil
}
