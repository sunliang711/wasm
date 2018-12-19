package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) importSection(sec *Section) error {
	err := checkSection(sec, types.OrderImport)
	if err != nil {
		return err
	}

	rd := bytes.NewReader(sec.Data)
	var importSize uint32
	err = utils.DecodeVarInt(rd, 32, &importSize)
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
			var funcTypeIndex uint32
			err := utils.DecodeVarInt(rd, 32, &funcTypeIndex)
			if err != nil {
				return err
			}
			//after type section parsed
			<-p.typeParsed
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
			logrus.Infof("<import section> type: function, function type: %v", imIndexFuncType)
		case types.Table:
			tableType, err := DecodeTableType(rd)
			if err != nil {
				return err
			}
			imTableType := types.ImportTableType{
				Type: tableType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			p.Module.Tables.Imports = append(p.Module.Tables.Imports, imTableType)
			logrus.Infof("<import section> type: table, table type: %v", imTableType)

		case types.Memory:
			var (
				memoryType types.MemoryType
			)
			memoryType.IsShared, memoryType.Size.Min, memoryType.Size.Max, err = DecodeFlags(rd)
			if err != nil {
				return err
			}
			imMemoryType := types.ImportMemoryType{
				Type: memoryType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			logrus.Infof("<import section> type: memory, memory type: %v", imMemoryType)
			p.Module.Memories.Imports = append(p.Module.Memories.Imports, imMemoryType)

		case types.Global:
			globalType, err := DecodeGlobalType(rd)
			if err != nil {
				return err
			}
			imGlobalType := types.ImportGlobalType{
				Type: globalType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			logrus.Infof("<import section> type: global, global type: %v", imGlobalType)
			p.Module.Globals.Imports = append(p.Module.Globals.Imports, imGlobalType)

		case types.Exception:
			var (
				exceptionType types.ExceptionType
			)
			err := DecodeTypeTuple(rd, &exceptionType.Params)
			if err != nil {
				return err
			}

			imExceptionType := types.ImportExceptionType{
				Type: exceptionType,
				ImportCommon: types.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			logrus.Infof("<import section> type: excetion, exception type: %v", imExceptionType)
			p.Module.ExceptionTypes.Imports = append(p.Module.ExceptionTypes.Imports, imExceptionType)
		}

	}

	err = p.validateImport()
	return err
}

func (p *Parser) validateImport() error {
	logrus.Info("TODO: validateImport()")
	return nil
}
