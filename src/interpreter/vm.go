package interpreter

import (
	"WGVM/src/common"
)

type Vm struct {
	operandStack
	module common.Module
}

func ExecMainFunc(module common.Module) {
	idx := int(*module.StartSec) - len(module.ImportSec)
	vm := &Vm{module: module}
	vm.execCode(idx)
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
