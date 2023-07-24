package interpreter

import (
	"WGVM/src/common"
	"WGVM/src/common/op"
)

type Vm struct {
	operandStack
	controlStack
	module    common.Module
	memory    *memory
	globals   []*globalVar
	local0Idx uint32
}

func ExecMainFunc(module common.Module) {
	vm := &Vm{module: module}
	vm.initMem()
	vm.initGlobals()
	call(vm, *module.StartSec)
	vm.loop()
}

func (vm *Vm) initMem() {
	if len(vm.module.MemSec) > 0 {
		vm.memory = newMemory(vm.module.MemSec[0])
	}
	for _, data := range vm.module.DataSec {
		for _, instr := range data.Offset {
			vm.execInstr(instr)
		}
		vm.memory.Write(vm.PopU64(), data.Init)
	}
}

func (vm *Vm) loop() {
	depth := vm.controlDepth()
	for vm.controlDepth() >= depth {
		cf := vm.topControlFrame()
		if cf.pc == len(cf.instrs) {
			vm.exitBlock()
		} else {
			instr := cf.instrs[cf.pc]
			cf.pc++
			vm.execInstr(instr)
		}
	}
}

func (vm *Vm) execInstr(instr common.Instruction) {
	//fmt.Printf("%s %v\n", instr.GetOpname(), instr.Args)
	InstrTable[instr.Opcode](vm, instr.Args)
}

func (vm *Vm) enterBlock(opcode byte,
	bt common.FuncType, instrs []common.Instruction) {
	// 在当前操作数栈的高度上减去参数数量，起到了调用帧重叠的效果
	bp := vm.stackSize() - len(bt.ParamTypes)
	cf := newControlFrame(opcode, bt, instrs, bp)
	vm.pushControlFrame(cf)
	if opcode == op.Call {
		vm.local0Idx = uint32(bp)
	}
}

func (vm *Vm) exitBlock() {
	cf := vm.popControlFrame()
	vm.clearBlock(cf)
}

func (vm *Vm) clearBlock(cf *controlFrame) {
	results := vm.popU64s(len(cf.bt.ResultTypes))
	vm.popU64s(vm.stackSize() - cf.bp)
	vm.pushU64s(results)
	if cf.opcode == op.Call && vm.controlDepth() > 0 {
		lastCallFrame, _ := vm.topCallFrame()
		vm.local0Idx = uint32(lastCallFrame.bp)
	}
}

func (vm *Vm) initGlobals() {
	for _, global := range vm.module.GlobalSec {
		for _, instr := range global.Init {
			vm.execInstr(instr)
		}
		vm.globals = append(vm.globals,
			newGlobal(global.Type, vm.PopU64()))
	}
}
