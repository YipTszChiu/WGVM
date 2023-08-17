package interpreter

import (
	"WGVM/src/common"
	"WGVM/src/common/op"
	"WGVM/src/instance"
	"fmt"
)

type WasmVal = instance.WasmVal

type Vm struct {
	operandStack
	controlStack
	module    common.Module
	memory    instance.Memory
	table     instance.Table
	globals   []instance.Global
	funcs     []vmFunc
	local0Idx uint32
}

func New(m common.Module, mm map[string]instance.Module,
) (inst instance.Module, err error) {

	defer func() {
		if _err := recover(); _err != nil {
			switch x := _err.(type) {
			case error:
				err = x
			default:
				panic(err)
			}
		}
	}()

	inst = newVM(m, mm)
	return
}

func newVM(m common.Module, mm map[string]instance.Module) *Vm {
	vm := &Vm{module: m}
	vm.linkImports(mm)
	vm.initFuncs()
	vm.initTable()
	vm.initMem()
	vm.initGlobals()
	vm.execStartFunc()
	return vm
}

/* linking */

func (vm *Vm) linkImports(mm map[string]instance.Module) {
	for _, imp := range vm.module.ImportSec {
		if m := mm[imp.Module]; m == nil {
			panic(fmt.Errorf("module not found: " + imp.Module))
		} else {
			vm.linkImport(m, imp)
		}
	}
}
func (vm *Vm) linkImport(m instance.Module, imp common.Import) {
	exported := m.GetMember(imp.Name)
	if exported == nil {
		panic(fmt.Errorf("unknown import: %s.%s",
			imp.Module, imp.Name))
	}

	typeMatched := false
	switch x := exported.(type) {
	case instance.Function:
		if imp.Desc.Tag == common.ImportTagFunc {
			expectedFT := vm.module.TypeSec[imp.Desc.FuncType]
			typeMatched = isFuncTypeMatch(expectedFT, x.Type())
			vm.funcs = append(vm.funcs, newExternalFunc(expectedFT, x))
		}
	case instance.Table:
		if imp.Desc.Tag == common.ImportTagTable {
			typeMatched = isLimitsMatch(imp.Desc.Table.Limits, x.Type().Limits)
			vm.table = x
		}
	case instance.Memory:
		if imp.Desc.Tag == common.ImportTagMem {
			typeMatched = isLimitsMatch(imp.Desc.Mem, x.Type())
			vm.memory = x
		}
	case instance.Global:
		if imp.Desc.Tag == common.ImportTagGlobal {
			typeMatched = isGlobalTypeMatch(imp.Desc.Global, x.Type())
			vm.globals = append(vm.globals, x)
		}
	}

	if !typeMatched {
		panic(fmt.Errorf("incompatible import type: %s.%s",
			imp.Module, imp.Name))
	}
}

/* init */

func (vm *Vm) initMem() {
	if len(vm.module.MemSec) > 0 {
		vm.memory = newMemory(vm.module.MemSec[0])
	}
	for _, data := range vm.module.DataSec {
		vm.execConstExpr(data.Offset)
		vm.memory.Write(vm.PopU64(), data.Init)
	}
}
func (vm *Vm) initGlobals() {
	for _, global := range vm.module.GlobalSec {
		vm.execConstExpr(global.Init)
		vm.globals = append(vm.globals,
			newGlobal(global.Type, vm.PopU64()))
	}
}
func (vm *Vm) initFuncs() {
	for i, ftIdx := range vm.module.FuncSec {
		ft := vm.module.TypeSec[ftIdx]
		code := vm.module.CodeSec[i]
		vm.funcs = append(vm.funcs, newInternalFunc(vm, ft, code))
	}
}
func (vm *Vm) initTable() {
	if len(vm.module.TableSec) > 0 {
		vm.table = newTable(vm.module.TableSec[0])
	}
	for _, elem := range vm.module.ElemSec {
		vm.execConstExpr(elem.Offset)
		offset := vm.PopU32()
		for i, funcIdx := range elem.Init {
			vm.table.SetElem(offset+uint32(i), vm.funcs[funcIdx])
		}
	}
}

func (vm *Vm) execConstExpr(expr []common.Instruction) {
	for _, instr := range expr {
		vm.execInstr(instr)
	}
}
func (vm *Vm) execStartFunc() {
	if vm.module.StartSec != nil {
		idx := *vm.module.StartSec
		vm.funcs[idx].call(nil)
	}
}

/* block stack */

func (vm *Vm) enterBlock(opcode byte,
	bt common.FuncType, instrs []common.Instruction) {

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
func (vm *Vm) resetBlock(cf *controlFrame) {
	results := vm.popU64s(len(cf.bt.ParamTypes))
	vm.popU64s(vm.stackSize() - cf.bp)
	vm.pushU64s(results)
}

/* loop */

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

/* instance.Module */

func (vm *Vm) GetMember(name string) interface{} {
	for _, exp := range vm.module.ExportSec {
		if exp.Name == name {
			idx := exp.Desc.Idx
			switch exp.Desc.Tag {
			case common.ExportTagFunc:
				return vm.funcs[idx]
			case common.ExportTagTable:
				return vm.table
			case common.ExportTagMem:
				return vm.memory
			case common.ExportTagGlobal:
				return vm.globals[idx]
			}
		}
	}
	return nil
}

func (vm *Vm) InvokeFunc(name string, args ...WasmVal) ([]WasmVal, error) {
	m := vm.GetMember(name)
	if m != nil {
		if f, ok := m.(instance.Function); ok {
			return f.Call(args...)
		}
	}
	return nil, fmt.Errorf("function not found: " + name)
}
func (vm Vm) GetGlobalVal(name string) (WasmVal, error) {
	m := vm.GetMember(name)
	if m != nil {
		if g, ok := m.(instance.Global); ok {
			return g.Get(), nil
		}
	}
	return nil, fmt.Errorf("global not found: " + name)
}
func (vm Vm) SetGlobalVal(name string, val WasmVal) error {
	m := vm.GetMember(name)
	if m != nil {
		if g, ok := m.(instance.Global); ok {
			g.Set(val)
			return nil
		}
	}
	return fmt.Errorf("global not found: " + name)
}

/* helpers */

func isFuncTypeMatch(expected, actual common.FuncType) bool {
	return fmt.Sprintf("%s", expected) == fmt.Sprintf("%s", actual)
}
func isGlobalTypeMatch(expected, actual common.GlobalType) bool {
	return actual.ValType == expected.ValType &&
		actual.Mut == expected.Mut
}
func isLimitsMatch(expected, actual common.Limits) bool {
	return actual.Min >= expected.Min &&
		(expected.Max == 0 || actual.Max > 0 && actual.Max <= expected.Max)
}
