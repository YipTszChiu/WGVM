package interpreter

func drop(vm *Vm, _ interface{}) {
	vm.PopU64()
}

func _select(vm *Vm, _ interface{}) {
	v3 := vm.PopBool()
	v2 := vm.PopU64()
	v1 := vm.PopU64()

	if v3 {
		vm.PushU64(v1)
	} else {
		vm.PushU64(v2)
	}
}
