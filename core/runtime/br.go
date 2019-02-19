package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

func br(vm *VM, frame *Frame, endIndice []int) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	branchImm, ok := ins.Imm.(*IR.BranchImm)
	if !ok {
		vm.panic("opcode: br invalid imm")
	}
	td := branchImm.TargetDepth
	tarInsIndex := endIndice[int(td)]
	frame.advanceTo(tarInsIndex)
	return
}

func br_if(vm *VM, frame *Frame, endIndice []int) (err error) {
	defer utils.CatchError(&err)

	con, err := pop1(vm, frame)
	if err != nil {
		vm.panic(err)
	}

	ins := frame.Instruction[frame.PC]
	v, ok := con.Value().(int32)
	if !ok {
		vm.panic("opcode: br_if parameter invalid")
	}
	if v == 1 {
		branchImm, ok := ins.Imm.(*IR.BranchImm)
		if !ok {
			vm.panic("opcode: br_if invalid imm")
		}
		td := branchImm.TargetDepth
		tarInsIndex := endIndice[int(td)]
		frame.advanceTo(tarInsIndex)
	} else {
		frame.advance(1)
	}
	return
}
func br_table(vm *VM, frame *Frame, endIndice []int) (err error) {
	defer utils.CatchError(&err)
	return
}
