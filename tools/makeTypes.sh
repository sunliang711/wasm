#!/bin/bash
./c2go -do ../types/opcode.go opcode.c
./c2go -do ../types/opcodeAndImm.go opcodeAndImm.c
./c2go -do ../parser/DecodeOpcodeAndImm.go DecodeOpcodeAndImm.c
