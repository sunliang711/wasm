package parser

import (
	"testing"
)

func TestParser_NotifyError(t *testing.T) {
	inputfile := "../helloworld.wasm"
	parser, err := NewParser(inputfile)
	if err != nil {
		t.Fatal(err)
	}
	parser.Parse()
}
