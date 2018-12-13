package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) importSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO importSection()")

	err = p.validateImport()
	return err
}

func (p *Parser) validateImport() error {
	logrus.Info("TODO: validateImport()")
	return nil
}
