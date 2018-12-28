#!/bin/bash
./c2go -do ../types/IR/opcode.go opcode.c
./c2go -do ../types/IR/opcodeAndImm.go opcodeAndImm.c
./c2go -do ../parser/DecodeOpcodeAndImm.go DecodeOpcodeAndImm.c
