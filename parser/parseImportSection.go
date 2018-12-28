package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/types/IR"
	"wasm/utils"
)

func (p *Parser) importSection(sec *Section) error {
	err := checkSection(sec, IR.OrderImport)
	if err != nil {
		return err
	}

	rd := bytes.NewReader(sec.Data)
	var importSize uint32
	_, err = utils.DecodeVarInt(rd, 32, &importSize)
	if err != nil {
		return err
	}
	bufKind := make([]byte, 1)
	for i := 0; i < int(importSize); i++ {
		//1. module name
		// 1.1 num char
		// 1.2 chars
		_, moduleName, err := utils.ReadVarChars(rd)
		if err != nil {
			return err
		}
		//2. export name
		// 2.1 num char
		// 2.2 chars
		_, exportName, err := utils.ReadVarChars(rd)
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
		kind := IR.ExternKind(bufKind[0])

		//4. switch kind
		switch kind {
		case IR.Function:
			//function type index(of type section)
			var funcTypeIndex uint32
			_, err := utils.DecodeVarInt(rd, 32, &funcTypeIndex)
			if err != nil {
				return err
			}
			//after type section parsed
			<-p.typeParsed
			if int(funcTypeIndex) >= len(p.Module.Types) {
				return fmt.Errorf(types.ErrFunctionTypeIndexOutOfRange)
			}
			imIndexFuncType := IR.ImportIndexedFunctionType{
				Type: IR.IndexedFunctionType{Index: uint64(funcTypeIndex)},
				ImportCommon: IR.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				},
			}
			p.Module.Functions.Imports = append(p.Module.Functions.Imports, imIndexFuncType)
			logrus.Infof("<import section> type: function, function type: %v", imIndexFuncType)
		case IR.Table:
			tableType, err := DecodeTableType(rd)
			if err != nil {
				return err
			}
			imTableType := IR.ImportTableType{
				Type: tableType,
				ImportCommon: IR.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			p.Module.Tables.Imports = append(p.Module.Tables.Imports, imTableType)
			logrus.Infof("<import section> type: table, table type: %v", imTableType)

		case IR.Memory:
			var (
				memoryType IR.MemoryType
			)
			memoryType.IsShared, memoryType.Size.Min, memoryType.Size.Max, err = DecodeFlags(rd)
			if err != nil {
				return err
			}
			imMemoryType := IR.ImportMemoryType{
				Type: memoryType,
				ImportCommon: IR.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			logrus.Infof("<import section> type: memory, memory type: %v", imMemoryType)
			p.Module.Memories.Imports = append(p.Module.Memories.Imports, imMemoryType)

		case IR.Global:
			globalType, err := DecodeGlobalType(rd)
			if err != nil {
				return err
			}
			imGlobalType := IR.ImportGlobalType{
				Type: globalType,
				ImportCommon: IR.ImportCommon{
					ModuleName: string(moduleName),
					ExportName: string(exportName),
				}}
			logrus.Infof("<import section> type: global, global type: %v", imGlobalType)
			p.Module.Globals.Imports = append(p.Module.Globals.Imports, imGlobalType)

		case IR.Exception:
			var (
				exceptionType IR.ExceptionType
			)
			err := DecodeTypeTuple(rd, &exceptionType.Params)
			if err != nil {
				return err
			}

			imExceptionType := IR.ImportExceptionType{
				Type: exceptionType,
				ImportCommon: IR.ImportCommon{
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
	//TODO
	logrus.Info("TODO: validateImport()")
	return nil
}
