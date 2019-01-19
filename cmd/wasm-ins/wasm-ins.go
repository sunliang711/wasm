package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"wasm/parser"
)

func main() {
	detail := flag.Bool("d", false, "show detail info")
	flag.Parse()
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("%s: List all instructions of all functions of wasm file\n", os.Args[0])
		fmt.Printf("Usage: %s wasmfile1 [wasmfile2 ...]\n", os.Args[0])
		return
	}
	logrus.SetLevel(logrus.PanicLevel)
	for _, wasmFile := range flag.Args() {
		parser, err := parser.NewParser(wasmFile)
		if err != nil {
			panic(err)
		}
		err = parser.Parse()
		if err != nil {
			panic(err)
		}
		fmt.Println(wasmFile)
		fmt.Println(parser.Module.GetAllFuncIns(*detail))
	}
}
