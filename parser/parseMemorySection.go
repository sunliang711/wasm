package parser

import (
	"github.com/sirupsen/logrus"
)

func (p *Parser) memorySection(sec *Section) error {
	var(
		err error
	)
	logrus.Info("TODO memorySection()")

	err = p.validateMemory()
	return err
}

func (p *Parser) validateMemory() error {
	logrus.Info("TODO: validateMemory()")
	return nil
}
