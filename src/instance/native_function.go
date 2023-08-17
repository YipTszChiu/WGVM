package instance

import "WGVM/src/common"

var _ Function = (*nativeFunction)(nil)

type GoFunc = func(args []WasmVal) ([]WasmVal, error)

type nativeFunction struct {
	t common.FuncType
	f GoFunc
}

func (nf nativeFunction) Type() common.FuncType {
	return nf.t
}
func (nf nativeFunction) Call(args ...WasmVal) ([]WasmVal, error) {
	return nf.f(args)
}
