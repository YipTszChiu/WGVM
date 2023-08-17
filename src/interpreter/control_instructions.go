package interpreter

import (
	"WGVM/src/common"
	"WGVM/src/common/op"
)

func unreachable(vm *Vm, _ interface{}) {
	panic(ErrTrap)
}

func nop(vm *Vm, _ interface{}) {
	// do nothing
}

func block(vm *Vm, args interface{}) {
	blockArgs := args.(common.BlockArgs)
	bt := vm.module.GetBlockType(blockArgs.BT)
	vm.enterBlock(op.Block, bt, blockArgs.Instrs)
}

func loop(vm *Vm, args interface{}) {
	blockArgs := args.(common.BlockArgs)
	bt := vm.module.GetBlockType(blockArgs.BT)
	vm.enterBlock(op.Loop, bt, blockArgs.Instrs)
}

func _if(vm *Vm, args interface{}) {
	ifArgs := args.(common.IfArgs)
	bt := vm.module.GetBlockType(ifArgs.BT)
	if vm.PopBool() {
		vm.enterBlock(op.If, bt, ifArgs.Instrs1)
	} else {
		vm.enterBlock(op.If, bt, ifArgs.Instrs2)
	}
}

func br(vm *Vm, args interface{}) {
	labelIdx := int(args.(uint32))
	for i := 0; i < labelIdx; i++ {
		vm.popControlFrame()
	}
	if cf := vm.topControlFrame(); cf.opcode != op.Loop {
		vm.exitBlock()
	} else {
		vm.resetBlock(cf)
		cf.pc = 0
	}
}

func brIf(vm *Vm, args interface{}) {
	if vm.PopBool() {
		br(vm, args)
	}
}

func brTable(vm *Vm, args interface{}) {
	brTableArgs := args.(common.BrTableArgs)
	n := int(vm.PopU32())
	if n < len(brTableArgs.Labels) {
		br(vm, brTableArgs.Labels[n])
	} else {
		br(vm, brTableArgs.Default)
	}
}

func _return(vm *Vm, _ interface{}) {
	_, labelIdx := vm.topCallFrame()
	br(vm, uint32(labelIdx))
}

func call(vm *Vm, args interface{}) {
	f := vm.funcs[args.(uint32)]
	callFunc(vm, f)
}

func callFunc(vm *Vm, f vmFunc) {
	if f._func != nil {
		callExternalFunc(vm, f)
	} else {
		callInternalFunc(vm, f)
	}
}

func callExternalFunc(vm *Vm, f vmFunc) {
	args := popArgs(vm, f._type)
	results, Err := f._func.Call(args...)
	if Err != nil {
		panic(Err)
	}
	pushResults(vm, f._type, results)
}

func popArgs(vm *Vm, ft common.FuncType) []interface{} {
	paramCount := len(ft.ParamTypes)
	args := make([]interface{}, paramCount)
	for i := paramCount - 1; i >= 0; i-- {
		args[i] = wrapU64(ft.ParamTypes[i], vm.PopU64())
	}
	return args
}

func pushResults(vm *Vm, ft common.FuncType, results []interface{}) {
	if len(ft.ResultTypes) != len(results) {
		panic("TODO")
	}
	for _, result := range results {
		vm.PushU64(unwrapU64(ft.ResultTypes[0], result))
	}
}

/*
operand stack:

+~~~~~~~~~~~~~~~+
|               |
+---------------+
|     stack     |
+---------------+
|     locals    |
+---------------+
|     params    |
+---------------+
|  ............ |
*/
func callInternalFunc(vm *Vm, f vmFunc) {
	vm.enterBlock(op.Call, f._type, f.code.Expr)

	// alloc locals
	localCount := int(f.code.GetLocalCount())
	for i := 0; i < localCount; i++ {
		vm.PushU64(0)
	}
}

func callIndirect(vm *Vm, args interface{}) {
	typeIdx := args.(uint32)
	ft := vm.module.TypeSec[typeIdx]

	i := vm.PopU32()
	if i >= vm.table.Size() {
		panic(ErrUndefinedElem)
	}

	f := vm.table.GetElem(i)
	if f.Type().GetSignature() != ft.GetSignature() {
		panic(ErrTypeMismatch)
	}

	// optimize internal func call
	if _f, ok := f.(vmFunc); ok {
		if _f._func == nil && _f.vm == vm {
			callInternalFunc(vm, _f)
			return
		}
	}

	fcArgs := popArgs(vm, ft)
	results, Err := f.Call(fcArgs...)
	if Err != nil {
		panic(Err)
	}
	pushResults(vm, ft, results)
}
