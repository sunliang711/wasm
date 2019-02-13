package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) elemSection(sec *Section) error {
	err := checkSection(sec, IR.OrderElem)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	//1. num eleSeg
	var numEleSeg uint32
	_, err = utils.DecodeVarInt(rd, 32, &numEleSeg)
	if err != nil {
		return err
	}

	//2. eleSegs
	for i := 0; i < int(numEleSeg); i++ {
		eleSeg := IR.ElemSegment{}
		flags, err := utils.ReadByte(rd)
		if err != nil {
			return err
		}
		switch flags {
		case 0:
			eleSeg.IsActive = true
			eleSeg.TableIndex = 0
			initExpression, err := DecodeInitializer(rd)
			if err != nil {
				return err
			}
			eleSeg.BaseOffset = &initExpression
		case 1:
			eleSeg.IsActive = false
			eleSeg.TableIndex = types.UINT64_MAX
			eleSeg.BaseOffset = &IR.InitializerExpression{}
		case 2:
			eleSeg.IsActive = true
			var tabIndex uint32
			_, err = utils.DecodeVarInt(rd, 32, &tabIndex)
			if err != nil {
				return err
			}
			eleSeg.TableIndex = uint64(tabIndex)
			initExpression, err := DecodeInitializer(rd)
			if err != nil {
				return err
			}
			eleSeg.BaseOffset = &initExpression
		default:
			return fmt.Errorf(types.ErrInvalidElemFlags)
		}
		//num ele
		var numEle uint32
		_, err = utils.DecodeVarInt(rd, 32, &numEle)
		if err != nil {
			return err
		}
		var funcIndex uint32
		for j := 0; j < int(numEle); j++ {

			_, err := utils.DecodeVarInt(rd, 32, &funcIndex)
			if err != nil {
				return err
			}
			eleSeg.Indices = append(eleSeg.Indices, funcIndex)
		}
		//eles
		p.Module.ElemSegments = append(p.Module.ElemSegments, eleSeg)
		logrus.Infof("<elem section> element segment: %v", eleSeg)
	}

	err = p.validateElem()
	return err
}

func (p *Parser) validateElem() error {
	//TODO
	logrus.Info("TODO: validateElem()")
	return nil
}
