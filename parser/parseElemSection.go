package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) elemSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO elemSection()")

	err = p.validateElem()
	return err
}

func (p *Parser) validateElem() error {
	logrus.Info("TODO: validateElem()")
	return nil
}
