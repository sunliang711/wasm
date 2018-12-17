package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) importSection(sec *Section) error {
	err := checkSection(sec, types.OrderImport)
	if err != nil {
		return err
	}

	rd := bytes.NewReader(sec.Data)
	importSize, err := utils.DecodeUInt32(rd)
	if err != nil {
		return err
	}
	bufKind := make([]byte, 1)
	for i := 0; i < int(importSize); i++ {
		//1. module name
		// 1.1 num char
		// 1.2 chars
		moduleName, err := utils.ReadVarChars(rd)
		if err != nil {
			return err
		}
		//2. export name
		// 2.1 num char
		// 2.2 chars
		exportName, err := utils.ReadVarChars(rd)
		if err != nil {
			return err
		}
		utils.CheckUTF8(moduleName)
		utils.CheckUTF8(exportName)
		logrus.Infof("import section {moduleName: %s,exportName: %s}", moduleName, exportName)

		//3. extern kind(1 byte,native value)
		n, err := rd.Read(bufKind)
		if err != nil {
			return err
		}
		if n != len(bufKind) {
			return fmt.Errorf(utils.ErrInsufficientChar)
		}
		kind := types.ExternKind(bufKind[0])

		//4. switch kind
		switch kind {
		case types.Function:
			//function type index(of type section)
			funcTypeIndex, err := utils.DecodeUInt32(rd)
			if err != nil {
				return err
			}
			if int(funcTypeIndex) >= len(p.Module.Types) {
				return fmt.Errorf(types.ErrFunctionTypeIndexOutOfRange)
			}
			imIndexFuncType := types.ImportIndexedFunctionType{
				Type: types.IndexedFunctionType{Index: uint64(funcTypeIndex)},
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				},
			}
			p.Module.Functions.Imports = append(p.Module.Functions.Imports, imIndexFuncType)
		case types.Table:
			var (
				tableType types.TableType
			)
			// A: ReferenceType (1 byte)
			refType, err := DecodeReferenceType(rd)
			if err != nil {
				return err
			}
			tableType.ElementType = refType
			//// B: flags
			tableType.IsShared, tableType.Size.Min, tableType.Size.Max, err = DecodeFlags(rd)
			if err != nil {
				return err
			}
			p.Module.Tables.Imports = append(p.Module.Tables.Imports, types.ImportTableType{
				Type: tableType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}})

		case types.Memory:
			var (
				memoryType types.MemoryType
			)
			memoryType.IsShared, memoryType.Size.Min, memoryType.Size.Max, err = DecodeFlags(rd)
			if err != nil {
				return err
			}
			p.Module.Memories.Imports = append(p.Module.Memories.Imports, types.ImportMemoryType{
				Type: memoryType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}})

		case types.Global:
			var (
				globalType types.GlobalType
			)
			// A. valueType
			vType, err := types.DecodeValueType(rd)
			if err != nil {
				return err
			}
			globalType.ValType = vType

			//B. isMutable
			isMutable, err := utils.DecodeU1(rd)
			if err != nil {
				return err
			}
			globalType.IsMutable = (isMutable != 0)

			p.Module.Globals.Imports = append(p.Module.Globals.Imports, types.ImportGlobalType{
				Type: globalType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}})

		case types.Exception:
			var (
				exceptionType types.ExceptionType
			)
			err := DecodeTypeTuple(rd, &exceptionType.Params)
			if err != nil {
				return err
			}

			p.Module.ExceptionTypes.Imports = append(p.Module.ExceptionTypes.Imports, types.ImportExceptionType{
				Type: exceptionType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}})
		}

	}

	err = p.validateImport()
	return err
}

func (p *Parser) validateImport() error {
	logrus.Info("TODO: validateImport()")
	return nil
}

func DecodeReferenceType(rd io.Reader) (types.RefType, error) {
	buf := make([]byte, 1)
	_, err := rd.Read(buf)
	if err != nil {
		return types.RTInvalid, err
	}
	switch buf[0] {
	case 0x70:
		return types.RTAnyFunc, nil
	case 0x6f:
		return types.RTAnyRef, nil
	default:
		return types.RTInvalid, fmt.Errorf(types.ErrReferenceTypeByte)
	}
}

func DecodeFlags(rd io.Reader) (bool, uint64, uint64, error) {
	flags, err := utils.DecodeUInt32(rd)
	if err != nil {
		return false, 0, 0, err
	}
	isShared := (flags&0x02 != 0)
	min, err := utils.DecodeUInt32(rd)
	if err != nil {
		return false, 0, 0, err
	}
	var (
		max uint64
	)
	hasMax := (flags&0x01 != 0)
	if hasMax {
		max, err = utils.DecodeUInt64(rd)
		if err != nil {
			return false, 0, 0, err
		}
	} else {
		max = types.UINT64_MAX
	}
	return isShared, uint64(min), max, nil
}
