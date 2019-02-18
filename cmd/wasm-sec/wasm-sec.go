package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"wasm/parser"
)

func main() {
	flag.Parse()
	logrus.SetLevel(logrus.PanicLevel)
	for _, wasmFile := range flag.Args() {
		contents, err := ioutil.ReadFile(wasmFile)
		if err != nil {
			panic(err)
		}
		parser, err := parser.NewParser(bytes.NewReader(contents))
		if err != nil {
			panic(err)
		}
		err = parser.Parse()
		if err != nil {
			panic(err)
		}
		fmt.Println(wasmFile + ":")
		fmt.Println(parser.Module.GetAllSections())
	}
}
