#!/bin/bash
./c2go -do ../core/IR/opcode.go opcode.c
./c2go -do ../core/IR/opcodeAndImm.go opcodeAndImm.c
./c2go -do ../parser/DecodeOpcodeAndImm.go DecodeOpcodeAndImm.c
