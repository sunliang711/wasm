package runtime

import (
	"testing"
	"wasm/parser"
)

func TestVM(t *testing.T) {
	inputfile := "/Users/eagle/Downloads/add.wasm"
	inputfile = "/home/eagle/Downloads/test.wasm"
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
	//params := []IR.InterfaceValue{
	//	&Value{Typ: IR.TypeI32, Val: uint32(20)},
	//	&Value{Typ: IR.TypeI32, Val: uint32(3)},
	//}

	err = vm.Run("_Z6dividejj", uint32(2), uint16(3))
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(vm.ReturnValue.Value())
	}
}

func TestA(t *testing.T) {

}
