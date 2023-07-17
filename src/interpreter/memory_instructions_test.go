package interpreter

import (
	"WGVM/src/common"
	"WGVM/src/common/op"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemSizeAndGrow(t *testing.T) {
	vm := &Vm{memory: newMemory(common.MemType{Min: 2})}
	InstrTable[op.MemorySize](vm, nil)
	require.Equal(t, uint64(2), vm.PopU64())

	vm.PushU32(3)
	InstrTable[op.MemoryGrow](vm, nil)
	require.Equal(t, uint64(2), vm.PopU64())

	InstrTable[op.MemorySize](vm, nil)
	require.Equal(t, uint64(5), vm.PopU64())
}

func TestMemOps(t *testing.T) {
	vm := &Vm{memory: newMemory(common.MemType{Min: 1})}
	testMemOp(t, vm, op.I32Store, op.I32Load, 0x10, 0x01, int32(100))
	testMemOp(t, vm, op.I64Store, op.I64Load, 0x20, 0x02, int64(123))
	testMemOp(t, vm, op.F32Store, op.F32Load, 0x30, 0x03, float32(1.5))
	testMemOp(t, vm, op.F64Store, op.F64Load, 0x40, 0x04, 1.5)
	testMemOp(t, vm, op.I32Store8, op.I32Load8S, 0x50, 0x05, int32(-100))
	testMemOp(t, vm, op.I32Store8, op.I32Load8U, 0x60, 0x06, int32(100))
	testMemOp(t, vm, op.I32Store16, op.I32Load16S, 0x70, 0x07, int32(-10000))
	testMemOp(t, vm, op.I32Store16, op.I32Load16U, 0x80, 0x08, int32(10000))
	testMemOp(t, vm, op.I64Store8, op.I64Load8S, 0x90, 0x09, int32(-100))
	testMemOp(t, vm, op.I64Store8, op.I64Load8U, 0xA0, 0x0A, int32(100))
	testMemOp(t, vm, op.I64Store16, op.I64Load16S, 0xB0, 0x0B, int32(-10000))
	testMemOp(t, vm, op.I64Store16, op.I64Load16U, 0xC0, 0x0C, int32(10000))
	testMemOp(t, vm, op.I64Store32, op.I64Load32S, 0xD0, 0x0D, int32(-1000000))
	testMemOp(t, vm, op.I64Store32, op.I64Load32U, 0xE0, 0x0E, int32(1000000))
}

func testMemOp(t *testing.T, vm *Vm, storeOp, loadOp byte,
	offset, i uint32, val interface{}) {

	memArg := common.MemArg{Offset: offset}

	// store
	vm.PushU32(i)
	pushVal(vm, val)
	InstrTable[storeOp](vm, memArg)

	// load
	vm.PushU32(i)
	InstrTable[loadOp](vm, memArg)

	// check
	require.Equal(t, val, popVal(vm, val))
}
