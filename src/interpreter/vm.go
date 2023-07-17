package interpreter

import (
	"WGVM/src/common"
)

type Vm struct {
	operandStack
	module common.Module
	memory *memory
}

func ExecMainFunc(module common.Module) {
	idx := int(*module.StartSec) - len(module.ImportSec)
	vm := &Vm{module: module}
	vm.initMem()
	vm.execCode(idx)
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

func (vm *Vm) execCode(idx int) {
	code := vm.module.CodeSec[idx]
	for _, instr := range code.Expr {
		vm.execInstr(instr)
	}
}

func (vm *Vm) execInstr(instr common.Instruction) {
	//fmt.Printf("%s %v\n", instr.GetOpname(), instr.Args)
	InstrTable[instr.Opcode](vm, instr.Args)
}
