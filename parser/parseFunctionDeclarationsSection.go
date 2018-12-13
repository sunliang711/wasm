package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) functionDeclarationsSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO functionDeclarationsSection()")

	err = p.validateFunctionDeclarations()
	return err
}
func (p *Parser) validateFunctionDeclarations() error {
	logrus.Info("TODO: validateFunctionDeclarations()")
	return nil
}
