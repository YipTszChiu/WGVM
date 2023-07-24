package interpreter

import "WGVM/src/common"

type globalVar struct {
	_type common.GlobalType
	val   uint64
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
