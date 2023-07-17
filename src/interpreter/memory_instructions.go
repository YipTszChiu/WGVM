package interpreter

import (
	"WGVM/src/common"
	gobin "encoding/binary"
)

var byteOrder = gobin.LittleEndian

func memorySize(vm *Vm, _ interface{}) {
	vm.PushU32(vm.memory.Size())
}
func memoryGrow(vm *Vm, _ interface{}) {
	oldSize := vm.memory.Grow(vm.PopU32())
	vm.PushU32(oldSize)
}

// load
func i32Load(vm *Vm, memArg interface{}) {
	val := readU32(vm, memArg)
	vm.PushU32(val)
}
func i64Load(vm *Vm, memArg interface{}) {
	val := readU64(vm, memArg)
	vm.PushU64(val)
}
func f32Load(vm *Vm, memArg interface{}) {
	val := readU32(vm, memArg)
	vm.PushU32(val)
}
func f64Load(vm *Vm, memArg interface{}) {
	val := readU64(vm, memArg)
	vm.PushU64(val)
}
func i32Load8S(vm *Vm, memArg interface{}) {
	val := readU8(vm, memArg)
	vm.PushS32(int32(int8(val)))
}
func i32Load8U(vm *Vm, memArg interface{}) {
	val := readU8(vm, memArg)
	vm.PushU32(uint32(val))
}
func i32Load16S(vm *Vm, memArg interface{}) {
	val := readU16(vm, memArg)
	vm.PushS32(int32(int16(val)))
}
func i32Load16U(vm *Vm, memArg interface{}) {
	val := readU16(vm, memArg)
	vm.PushU32(uint32(val))
}
func i64Load8S(vm *Vm, memArg interface{}) {
	val := readU8(vm, memArg)
	vm.PushS64(int64(int8(val)))
}
func i64Load8U(vm *Vm, memArg interface{}) {
	val := readU8(vm, memArg)
	vm.PushU64(uint64(val))
}
func i64Load16S(vm *Vm, memArg interface{}) {
	val := readU16(vm, memArg)
	vm.PushS64(int64(int16(val)))
}
func i64Load16U(vm *Vm, memArg interface{}) {
	val := readU16(vm, memArg)
	vm.PushU64(uint64(val))
}
func i64Load32S(vm *Vm, memArg interface{}) {
	val := readU32(vm, memArg)
	vm.PushS64(int64(int32(val)))
}
func i64Load32U(vm *Vm, memArg interface{}) {
	val := readU32(vm, memArg)
	vm.PushU64(uint64(val))
}

// store
func i32Store(vm *Vm, memArg interface{}) {
	val := vm.PopU32()
	writeU32(vm, memArg, val)
}
func i64Store(vm *Vm, memArg interface{}) {
	val := vm.PopU64()
	writeU64(vm, memArg, val)
}
func f32Store(vm *Vm, memArg interface{}) {
	val := vm.PopU32()
	writeU32(vm, memArg, val)
}
func f64Store(vm *Vm, memArg interface{}) {
	val := vm.PopU64()
	writeU64(vm, memArg, val)
}
func i32Store8(vm *Vm, memArg interface{}) {
	val := vm.PopU32()
	writeU8(vm, memArg, byte(val))
}
func i32Store16(vm *Vm, memArg interface{}) {
	val := vm.PopU32()
	writeU16(vm, memArg, uint16(val))
}
func i64Store8(vm *Vm, memArg interface{}) {
	val := vm.PopU64()
	writeU8(vm, memArg, byte(val))
}
func i64Store16(vm *Vm, memArg interface{}) {
	val := vm.PopU64()
	writeU16(vm, memArg, uint16(val))
}
func i64Store32(vm *Vm, memArg interface{}) {
	val := vm.PopU64()
	writeU32(vm, memArg, uint32(val))
}

func writeU8(vm *Vm, memArg interface{}, n byte) {
	var buf [1]byte
	buf[0] = n
	offset := getOffset(vm, memArg)
	vm.memory.Write(offset, buf[:])
}
func writeU16(vm *Vm, memArg interface{}, n uint16) {
	var buf [2]byte
	byteOrder.PutUint16(buf[:], n)
	offset := getOffset(vm, memArg)
	vm.memory.Write(offset, buf[:])
}
func writeU32(vm *Vm, memArg interface{}, n uint32) {
	var buf [4]byte
	byteOrder.PutUint32(buf[:], n)
	offset := getOffset(vm, memArg)
	vm.memory.Write(offset, buf[:])
}
func writeU64(vm *Vm, memArg interface{}, n uint64) {
	var buf [8]byte
	byteOrder.PutUint64(buf[:], n)
	offset := getOffset(vm, memArg)
	vm.memory.Write(offset, buf[:])
}

func readU8(vm *Vm, memArg interface{}) byte {
	var buf [1]byte
	offset := getOffset(vm, memArg)
	vm.memory.Read(offset, buf[:])
	return buf[0]
}
func readU16(vm *Vm, memArg interface{}) uint16 {
	var buf [2]byte
	offset := getOffset(vm, memArg)
	vm.memory.Read(offset, buf[:])
	return byteOrder.Uint16(buf[:])
}
func readU32(vm *Vm, memArg interface{}) uint32 {
	var buf [4]byte
	offset := getOffset(vm, memArg)
	vm.memory.Read(offset, buf[:])
	return byteOrder.Uint32(buf[:])
}
func readU64(vm *Vm, memArg interface{}) uint64 {
	var buf [8]byte
	offset := getOffset(vm, memArg)
	vm.memory.Read(offset, buf[:])
	return byteOrder.Uint64(buf[:])
}

func getOffset(vm *Vm, memArg interface{}) uint64 {
	offset := memArg.(common.MemArg).Offset
	return uint64(vm.PopU32()) + uint64(offset)
}
