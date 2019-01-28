package runtime

import (
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

const (
	CMP_LT byte = iota
	CMP_LE
	CMP_EQ
	CMP_NE
	CMP_GT
	CMP_GE
)

func i32_compare(vm *VM, frame *Frame, cmpType byte, isSigned bool) (err error) {
	defer utils.CatchError(&err)
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	b, _ := frame.Stack.Pop()
	a, _ := frame.Stack.Pop()
	if a.Type() != b.Type() {
		vm.panic("cmp type not match")
	}

	if isSigned {
		lh := a.Value().(int32)
		rh := b.Value().(int32)
		switch cmpType {
		case CMP_LT:
			if lh < rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_LE:
			if lh <= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_EQ:
			if lh == rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_NE:
			if lh != rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GT:
			if lh > rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GE:
			if lh >= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		}
	} else {
		lh := a.Value().(uint32)
		rh := b.Value().(uint32)
		switch cmpType {
		case CMP_LT:
			if lh < rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_LE:
			if lh <= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_EQ:
			if lh == rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_NE:
			if lh != rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GT:
			if lh > rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GE:
			if lh >= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		}
	}
	frame.advance(1)
	return
}

func i64_compare(vm *VM, frame *Frame, cmpType byte, isSigned bool) (err error) {
	defer utils.CatchError(&err)
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	b, _ := frame.Stack.Pop()
	a, _ := frame.Stack.Pop()
	if a.Type() != b.Type() {
		vm.panic("cmp type not match")
	}

	if isSigned {
		lh := a.Value().(int64)
		rh := b.Value().(int64)
		switch cmpType {
		case CMP_LT:
			if lh < rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_LE:
			if lh <= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_EQ:
			if lh == rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_NE:
			if lh != rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GT:
			if lh > rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GE:
			if lh >= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		}
	} else {
		lh := a.Value().(uint64)
		rh := b.Value().(uint64)
		switch cmpType {
		case CMP_LT:
			if lh < rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_LE:
			if lh <= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_EQ:
			if lh == rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_NE:
			if lh != rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GT:
			if lh > rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		case CMP_GE:
			if lh >= rh {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
		}
	}
	frame.advance(1)
	return
}

func f32_compare(vm *VM, frame *Frame, cmpType byte) (err error) {
	defer utils.CatchError(&err)
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	b, _ := frame.Stack.Pop()
	a, _ := frame.Stack.Pop()
	if a.Type() != b.Type() {
		vm.panic("cmp type not match")
	}

	lh := a.Value().(float32)
	rh := b.Value().(float32)
	switch cmpType {
	case CMP_LT:
		if lh < rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_LE:
		if lh <= rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_EQ:
		if lh == rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_NE:
		if lh != rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_GT:
		if lh > rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_GE:
		if lh >= rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	}
	frame.advance(1)
	return
}

func f64_compare(vm *VM, frame *Frame, cmpType byte) (err error) {
	defer utils.CatchError(&err)
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	b, _ := frame.Stack.Pop()
	a, _ := frame.Stack.Pop()
	if a.Type() != b.Type() {
		vm.panic("cmp type not match")
	}

	lh := a.Value().(float64)
	rh := b.Value().(float64)
	switch cmpType {
	case CMP_LT:
		if lh < rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_LE:
		if lh <= rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_EQ:
		if lh == rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_NE:
		if lh != rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_GT:
		if lh > rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	case CMP_GE:
		if lh >= rh {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		}
	}
	frame.advance(1)
	return
}
