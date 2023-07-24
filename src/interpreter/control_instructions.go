package interpreter

import "WGVM/src/common/op"

func call(vm *Vm, args interface{}) {
	// 将立即数强转为函数索引
	idx := int(args.(uint32))
	// 目前只考虑导入函数，不考虑其他类型，后续会补充处理
	// 导入函数的数量
	importedFuncCount := len(vm.module.ImportSec) // TODO
	if idx < importedFuncCount {
		// 外部函数
		callAssertFunc(vm, args) // hack!
	} else {
		// 内部函数
		callInternalFunc(vm, idx-importedFuncCount)
	}
}

func callInternalFunc(vm *Vm, idx int) {
	// 获取函数的类型和指令
	ftIdx := vm.module.FuncSec[idx]
	ft := vm.module.TypeSec[ftIdx]
	code := vm.module.CodeSec[idx]
	// 创建一个新的调用帧，并给局部变量分配好空间
	vm.enterBlock(op.Call, ft, code.Expr)

	localCount := int(code.GetLocalCount())
	for i := 0; i < localCount; i++ {
		vm.PushU64(0)
	}
}

func callAssertFunc(vm *Vm, args interface{}) {
	// TODO
}
