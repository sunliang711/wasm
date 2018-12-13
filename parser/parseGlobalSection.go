package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) globalSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO globalSection()")

	err = p.validateGlobal()
	return err
}

func (p *Parser) validateGlobal() error {
	logrus.Info("TODO: validateGlobal()")
	return nil
}
