package runtime

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
	"wasm/parser"
)

func TestVM(t *testing.T) {
	inputfile := "../../example/test.wasm"
	//test(t,inputfile,"_Z6dividejj",uint32(2),uint16(3))
	//inputfile = "../../example/Sum.wasm"
	//test(t,inputfile,"_Z3Sumi",int32(2))

	inputfile = "/Users/eagle/Downloads/77.wasm"
	test(t, inputfile, 1000, "_Z3maxii", int32(2), int32(3))

}

func TestA(t *testing.T) {
	test(t, "../../example/br_if_memory.wasm", 100, "_Z2ffi", int32(3))
}

func TestMax(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/sum_max.wasm", 100, "_Z3maxii", int32(30), int32(5))
}

func TestArea(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/area.wasm", 10, "_Z4aread", float64(10))
}
func TestSum(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/sum_max.wasm", 900, "_Z3sumi", int32(30))
}
func test(t *testing.T, filename string, gas uint64, funcName string, params ...interface{}) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	parser, err := parser.NewParser(bytes.NewReader(bs))
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

	usedGas, err := vm.Run(funcName, gas, params...)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("used gas: ", usedGas)
		if vm.ReturnValue != nil {
			t.Log("return value:", vm.ReturnValue.Value())
		} else {
			t.Log("No return value.")
		}

	}
}

func TestPrintAllIns(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	wasmFile := "../../example/test.wasm"
	wasmFile = "../../example/br_if_memory.wasm"
	bs, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}
	parser, err := parser.NewParser(bytes.NewReader(bs))
	if err != nil {
		t.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parser.Module.GetAllFuncIns(true))
}

func TestLoadStoreI32(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/i32LoadStore.wasm", 100, "_Z1fv")
}

func TestLoadStoreI64(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/i64LoadStore.wasm", 100, "_Z1fv")
}

func TestI32Load8_s(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/i32Load8_s.wasm", 100, "_Z1fv")
}

func TestLoadStoreF32(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/f32LoadStore.wasm", 100, "_Z1fv")
}

func TestLoadStoreF64(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/f64LoadStore.wasm", 100, "_Z1fv")
}

func TestWrap64To32(t *testing.T) {
	var i64 int64 = 0x1234567811223344
	var i32 int32 = int32(i64)
	t.Logf("%#x", i32)
}

func TestAddGet(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	test(t, "../../example/add-get.wasm", 100, "_Z3getv")
	test(t, "../../example/add-get.wasm", 100, "_Z3addi", int32(2))
	test(t, "../../example/add-get.wasm", 100, "_Z3getv")
}
