package interpreter

func localGet(vm *Vm, args interface{}) {
	idx := args.(uint32)
	val := vm.getOperand(vm.local0Idx + idx)
	vm.PushU64(val)
}
func localSet(vm *Vm, args interface{}) {
	idx := args.(uint32)
	val := vm.PopU64()
	vm.setOperand(vm.local0Idx+idx, val)
}
func localTee(vm *Vm, args interface{}) {
	idx := args.(uint32)
	val := vm.PopU64()
	vm.PushU64(val)
	vm.setOperand(vm.local0Idx+idx, val)
}

func globalGet(vm *Vm, args interface{}) {
	idx := args.(uint32)
	val := vm.globals[idx].GetAsU64()
	vm.PushU64(val)
}
func globalSet(vm *Vm, args interface{}) {
	idx := args.(uint32)
	val := vm.PopU64()
	vm.globals[idx].SetAsU64(val)
}
