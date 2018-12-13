package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) exceptionTypesSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO exceptionTypesSection()")

	err = p.validateExceptionTypes()
	return err
}

func (p *Parser) validateExceptionTypes() error {
	logrus.Info("TODO: validateExceptionTypes()")
	return nil
}
