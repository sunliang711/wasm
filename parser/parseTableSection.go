package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) tableSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO tableSection()")

	err = p.validateTable()
	return err
}
func (p *Parser) validateTable() error {
	logrus.Info("TODO: validateTable()")
	return nil
}
