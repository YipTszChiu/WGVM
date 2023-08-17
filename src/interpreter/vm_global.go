package interpreter

import (
	"WGVM/src/common"
	"WGVM/src/instance"
)

type globalVar struct {
	_type common.GlobalType
	val   uint64
}

func NewGlobal(vt common.ValType, mut bool, val uint64) instance.Global {
	gt := common.GlobalType{ValType: vt}
	if mut {
		gt.Mut = 1
	}
	return newGlobal(gt, val)
}

func newGlobal(gt common.GlobalType, val uint64) *globalVar {
	return &globalVar{_type: gt, val: val}
}

func (g *globalVar) Type() common.GlobalType {
	return g._type
}

func (g *globalVar) GetAsU64() uint64 {
	return g.val
}
func (g *globalVar) SetAsU64(val uint64) {
	if g._type.Mut != 1 {
		panic(ErrImmutableGlobal)
	}
	g.val = val
}

func (g *globalVar) Get() instance.WasmVal {
	return wrapU64(g._type.ValType, g.val)
}
func (g *globalVar) Set(val instance.WasmVal) {
	g.val = unwrapU64(g._type.ValType, val)
}
