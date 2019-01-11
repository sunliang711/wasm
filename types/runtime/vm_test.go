package runtime

import (
	"testing"
	"wasm/parser"
	"wasm/types/IR"
)

func TestVM(t *testing.T) {

	inputfile := "/Users/eagle/Downloads/add.wasm"
	parser, err := parser.NewParser(inputfile)
	if err != nil {
		t.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		t.Fatal(err)
	}

	vm, err := NewVM(parser.Module)
	if err != nil {
		t.Fatal(err)
	}
	params := []IR.InterfaceValue{
		&Value{Typ: IR.TypeI32, Val: int32(1)},
		&Value{Typ: IR.TypeI32, Val: int32(2)}}

	err = vm.Run(0, params)
	if err != nil {
		t.Fatal(err)
	}
}
