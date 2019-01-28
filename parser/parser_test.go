package parser

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestParser_NotifyError(t *testing.T) {
	inputfile := "../helloworld.wasm"
	inputfile = "/Users/eagle/Downloads/main.wasm"
	inputfile = "/Users/eagle/Downloads/i32store.wasm"
	inputfile = "/Users/eagle/Downloads/test.wasm"
	bs, err := ioutil.ReadFile(inputfile)
	if err != nil {
		t.Fatal(err)
	}
	parser, err := NewParser(bytes.NewReader(bs))
	if err != nil {
		t.Fatal(err)
	}
	parser.Parse()
}
