package parser

import (
	"testing"
)

func TestParser_NotifyError(t *testing.T) {
	inputfile := "../helloworld.wasm"
	inputfile = "/Users/eagle/Downloads/main.wasm"
	inputfile = "/Users/eagle/Downloads/i32store.wasm"
	inputfile = "/Users/eagle/Downloads/test.wasm"
	parser, err := NewParser(inputfile)
	if err != nil {
		t.Fatal(err)
	}
	parser.Parse()
}
