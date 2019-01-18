package runtime

import (
	"github.com/sirupsen/logrus"
	"testing"
	"wasm/parser"
)

func TestVM(t *testing.T) {
	inputfile := "../../example/test.wasm"
	//test(t,inputfile,"_Z6dividejj",uint32(2),uint16(3))
	//inputfile = "../../example/Sum.wasm"
	//test(t,inputfile,"_Z3Sumi",int32(2))

	inputfile = "/Users/eagle/Downloads/77.wasm"
	test(t, inputfile, "_Z3maxii", int32(2), int32(3))

}

func TestA(t *testing.T) {
	test(t, "../../example/br_if_memory.wasm", "_Z2ffi", int32(3))
}

func TestMax(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/sum_max.wasm", "_Z3maxii", int32(30), int32(5))
}
func TestSum(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/sum_max.wasm", "_Z3sumi", int32(30))
}
func test(t *testing.T, filename string, funcName string, params ...interface{}) {
	parser, err := parser.NewParser(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = parser.Parse2()
	if err != nil {
		t.Fatal(err)
	}

	vm, err := NewVM(parser.Module)
	if err != nil {
		t.Fatal(err)
	}

	err = vm.Run(funcName, params...)
	if err != nil {
		t.Fatal(err)
	} else {
		if vm.ReturnValue != nil {
			t.Log(vm.ReturnValue.Value())
		} else {
			t.Log("No return value.")
		}

	}
}

func TestPrintAllIns(t *testing.T) {
	wasmFile := "/Users/eagle/Downloads/77.wasm"
	wasmFile = "/Users/eagle/Downloads/add3.wasm"
	wasmFile = "/Users/eagle/Downloads/hello.wasm"
	wasmFile = "/Users/eagle/Downloads/main.wasm"
	wasmFile = "/Users/eagle/Downloads/ret1.wasm"
	wasmFile = "/Users/eagle/Downloads/sqr.wasm"
	wasmFile = "/Users/eagle/Downloads/sum.wasm"
	wasmFile = "/Users/eagle/Downloads/add.wasm"
	wasmFile = "/Users/eagle/Downloads/fibo.wasm"
	wasmFile = "/Users/eagle/Downloads/helloworld.wasm"
	wasmFile = "/Users/eagle/Downloads/max.wasm"
	wasmFile = "/Users/eagle/Downloads/ret93.wasm"
	wasmFile = "/Users/eagle/Downloads/squareSum.wasm"
	wasmFile = "/Users/eagle/Downloads/test.wasm"
	wasmFile = "/Users/eagle/Downloads/add2.wasm"
	wasmFile = "/Users/eagle/Downloads/global.wasm"
	wasmFile = "/Users/eagle/Downloads/i32store.wasm"
	wasmFile = "/Users/eagle/Downloads/pow.wasm"
	wasmFile = "/Users/eagle/Downloads/simple.wasm"
	wasmFile = "/Users/eagle/Downloads/squareSum2.wasm"
	wasmFile = "/Users/eagle/Downloads/useadd.wasm"
	wasmFile = "../../example/test.wasm"
	wasmFile = "../../example/br_if_memory.wasm"
	parser, err := parser.NewParser(wasmFile)
	if err != nil {
		t.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parser.Module.GetAllFuncIns())
}
