package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) startSection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO startSection()")

	err = p.validateStart()
	return err
}
func (p *Parser) validateStart() error {
	logrus.Info("TODO: validateStart()")
	return nil
}
