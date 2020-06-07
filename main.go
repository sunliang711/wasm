package main

import (
	"fmt"
	"os"
	"wasm/parser"
)

func main() {
	//inputfile := "helloworld.wasm"
	inputfile := "/Users/eagle/Downloads/add2.wasm"
	//rdr ,err:= parser.LoadBinary(inputfile)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//err = parser.CheckConstant(rdr,parser.MagicNumber,"magic number error")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//err = parser.CheckConstant(rdr,parser.CurrentVersion,"version error")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//f,_ := os.Open("helloworld.wasm")
	//buf := make([]byte,1280)
	//counts:=0
	//for{
	//	n,err := f.Read(buf)
	//	if err != nil{
	//		fmt.Println("occur error: ",err)
	//		break
	//	}
	//	counts += 1
	//	fmt.Println("Read ",n," bytes")
	//}
	//fmt.Println("counts: ",counts)
	rd, err := os.Open(inputfile)
	if err != nil {
		panic(err)
	}
	parser, err := parser.NewParser(rd)
	if err != nil {
		panic(err)
	}
	fmt.Println(parser.Parse())
}
