package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) elemSection(sec *Section) error {
	err := checkSection(sec, types.OrderElem)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	//1. num ele
	var numEle uint32
	_, err = utils.DecodeVarInt(rd, 32, &numEle)
	if err != nil {
		return err
	}

	//2. eles
	for i := 0; i < int(numEle); i++ {
		eleSeg := types.ElemSegment{}
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
			eleSeg.BaseOffset = &types.InitializerExpression{}
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
