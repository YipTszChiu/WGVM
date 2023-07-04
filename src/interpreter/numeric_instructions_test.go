package interpreter

import (
	"WGVM/src/common/op"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConstOps(t *testing.T) {
	vm := &Vm{}
	i32Const(vm, int32(100))
	i64Const(vm, int64(200))
	f32Const(vm, float32(1.5))
	f64Const(vm, 2.5)
	require.Equal(t, 2.5, vm.PopF64())
	require.Equal(t, float32(1.5), vm.PopF32())
	require.Equal(t, int64(200), vm.PopS64())
	require.Equal(t, int32(100), vm.PopS32())
}

func TestI32Cmp(t *testing.T) {
	testI32UnOp(t, op.I32Eqz, 1, 0)
	testI32UnOp(t, op.I32Eqz, 0, 1)
	testI32BinOp(t, op.I32Eq, 1, 1, 1)
	testI32BinOp(t, op.I32Eq, 1, -1, 0)
	testI32BinOp(t, op.I32Ne, 1, 1, 0)
	testI32BinOp(t, op.I32Ne, 1, -1, 1)
	testI32BinOp(t, op.I32LtS, -1, 1, 1)
	testI32BinOp(t, op.I32LtS, 1, -1, 0)
	testI32BinOp(t, op.I32LtU, -1, 1, 0)
	testI32BinOp(t, op.I32LtU, 1, -1, 1)
	testI32BinOp(t, op.I32GtS, -1, 1, 0)
	testI32BinOp(t, op.I32GtS, 1, -1, 1)
	testI32BinOp(t, op.I32GtU, -1, 1, 1)
	testI32BinOp(t, op.I32GtU, 1, -1, 0)
	testI32BinOp(t, op.I32LeS, -1, 1, 1)
	testI32BinOp(t, op.I32LeS, 1, -1, 0)
	testI32BinOp(t, op.I32LeU, -1, 1, 0)
	testI32BinOp(t, op.I32LeU, 1, -1, 1)
	testI32BinOp(t, op.I32GeS, -1, 1, 0)
	testI32BinOp(t, op.I32GeS, 1, -1, 1)
	testI32BinOp(t, op.I32GeU, -1, 1, 1)
	testI32BinOp(t, op.I32GeU, 1, -1, 0)
}

func TestI64Cmp(t *testing.T) {
	testI64UnOp(t, op.I64Eqz, 1, 0)
	testI64UnOp(t, op.I64Eqz, 0, 1)
	testI64BinOp(t, op.I64Eq, 1, 1, 1)
	testI64BinOp(t, op.I64Eq, 1, -1, 0)
	testI64BinOp(t, op.I64Ne, 1, 1, 0)
	testI64BinOp(t, op.I64Ne, 1, -1, 1)
	testI64BinOp(t, op.I64LtS, -1, 1, 1)
	testI64BinOp(t, op.I64LtS, 1, -1, 0)
	testI64BinOp(t, op.I64LtU, -1, 1, 0)
	testI64BinOp(t, op.I64LtU, 1, -1, 1)
	testI64BinOp(t, op.I64GtS, -1, 1, 0)
	testI64BinOp(t, op.I64GtS, 1, -1, 1)
	testI64BinOp(t, op.I64GtU, -1, 1, 1)
	testI64BinOp(t, op.I64GtU, 1, -1, 0)
	testI64BinOp(t, op.I64LeS, -1, 1, 1)
	testI64BinOp(t, op.I64LeS, 1, -1, 0)
	testI64BinOp(t, op.I64LeU, -1, 1, 0)
	testI64BinOp(t, op.I64LeU, 1, -1, 1)
	testI64BinOp(t, op.I64GeS, -1, 1, 0)
	testI64BinOp(t, op.I64GeS, 1, -1, 1)
	testI64BinOp(t, op.I64GeU, -1, 1, 1)
	testI64BinOp(t, op.I64GeU, 1, -1, 0)
}

func TestF32Cmp(t *testing.T) {
	testF32BinCmp(t, op.F32Eq, 1.0, 1.0, 1)
	testF32BinCmp(t, op.F32Eq, 1.0, 2.0, 0)
	testF32BinCmp(t, op.F32Ne, 1.0, 1.0, 0)
	testF32BinCmp(t, op.F32Ne, 1.0, 2.0, 1)
	testF32BinCmp(t, op.F32Lt, 1.0, 2.0, 1)
	testF32BinCmp(t, op.F32Lt, 2.0, 1.0, 0)
	testF32BinCmp(t, op.F32Gt, 1.0, 2.0, 0)
	testF32BinCmp(t, op.F32Gt, 2.0, 1.0, 1)
	testF32BinCmp(t, op.F32Le, 1.0, 2.0, 1)
	testF32BinCmp(t, op.F32Le, 2.0, 1.0, 0)
	testF32BinCmp(t, op.F32Ge, 1.0, 2.0, 0)
	testF32BinCmp(t, op.F32Ge, 2.0, 1.0, 1)
}

func TestF64Cmp(t *testing.T) {
	testF64BinCmp(t, op.F64Eq, 1.0, 1.0, 1)
	testF64BinCmp(t, op.F64Eq, 1.0, 2.0, 0)
	testF64BinCmp(t, op.F64Ne, 1.0, 1.0, 0)
	testF64BinCmp(t, op.F64Ne, 1.0, 2.0, 1)
	testF64BinCmp(t, op.F64Lt, 1.0, 2.0, 1)
	testF64BinCmp(t, op.F64Lt, 2.0, 1.0, 0)
	testF64BinCmp(t, op.F64Gt, 1.0, 2.0, 0)
	testF64BinCmp(t, op.F64Gt, 2.0, 1.0, 1)
	testF64BinCmp(t, op.F64Le, 1.0, 2.0, 1)
	testF64BinCmp(t, op.F64Le, 2.0, 1.0, 0)
	testF64BinCmp(t, op.F64Ge, 1.0, 2.0, 0)
	testF64BinCmp(t, op.F64Ge, 2.0, 1.0, 1)
}

func TestI32Arithmetic(t *testing.T) {
	testI32UnOp(t, op.I32Clz, 0xF0, 24)
	testI32UnOp(t, op.I32Ctz, 0xF0, 4)
	testI32UnOp(t, op.I32PopCnt, 0xF0F0, 8)
	testI32BinOp(t, op.I32Add, 3, 2, 5)
	testI32BinOp(t, op.I32Sub, 3, 2, 1)
	testI32BinOp(t, op.I32Mul, 3, 2, 6)
	testI32BinOp(t, op.I32DivS, -8, 2, -4)
	testI32BinOp(t, op.I32DivU, -8, 2, 0x7FFF_FFFC)
	testI32BinOp(t, op.I32RemS, -5, 2, -1)
	testI32BinOp(t, op.I32RemU, -5, 2, 1)
	testI32BinOp(t, op.I32And, 0x0F0F, 0xF00F, 0x000F)
	testI32BinOp(t, op.I32Or, 0x0F0F, 0xF00F, 0xFF0F)
	testI32BinOp(t, op.I32Xor, 0x0F0F, 0xF00F, 0xFF00)
	testI32BinOp(t, op.I32Shl, -1, 8, -256)
	testI32BinOp(t, op.I32Shl, -1, 200, -256)
	testI32BinOp(t, op.I32ShrS, -1, 8, -1)
	testI32BinOp(t, op.I32ShrS, -1, 200, -1)
	testI32BinOp(t, op.I32ShrU, -1, 8, 0xFF_FFFF)
	testI32BinOp(t, op.I32ShrU, -1, 200, 0xFF_FFFF)
	testI32BinOp(t, op.I32Rotl, 0x1234_5678, 8, 0x3456_7812)
	testI32BinOp(t, op.I32Rotl, 0x1234_5678, 200, 0x3456_7812)
	testI32BinOp(t, op.I32Rotr, 0x1234_5678, 8, 0x7812_3456)
	testI32BinOp(t, op.I32Rotr, 0x1234_5678, 200, 0x7812_3456)
}

func TestI64Arithmetic(t *testing.T) {
	testI64UnOp(t, op.I64Clz, 0xF0, 56)
	testI64UnOp(t, op.I64Ctz, 0xF0, 4)
	testI64UnOp(t, op.I64PopCnt, 0xF0F0, 8)
	testI64BinOp(t, op.I64Add, 3, 2, 5)
	testI64BinOp(t, op.I64Sub, 3, 2, 1)
	testI64BinOp(t, op.I64Mul, 3, 2, 6)
	testI64BinOp(t, op.I64DivS, -8, 2, -4)
	testI64BinOp(t, op.I64DivU, -8, 2, 0x7FFF_FFFF_FFFF_FFFC)
	testI64BinOp(t, op.I64RemS, -5, 2, -1)
	testI64BinOp(t, op.I64RemU, -5, 2, 1)
	testI64BinOp(t, op.I64And, 0x0F0F, 0xF00F, 0x000F)
	testI64BinOp(t, op.I64Or, 0x0F0F, 0xF00F, 0xFF0F)
	testI64BinOp(t, op.I64Xor, 0x0F0F, 0xF00F, 0xFF00)
	testI64BinOp(t, op.I64Shl, -1, 8, -256)
	testI64BinOp(t, op.I64Shl, -1, 200, -256)
	testI64BinOp(t, op.I64ShrS, -1, 8, -1)
	testI64BinOp(t, op.I64ShrS, -1, 200, -1)
	testI64BinOp(t, op.I64ShrU, -1, 8, 0xFF_FFFF_FFFF_FFFF)
	testI64BinOp(t, op.I64ShrU, -1, 200, 0xFF_FFFF_FFFF_FFFF)
	testI64BinOp(t, op.I64Rotl, 0x1234_5678_1234_5678, 8, 0x3456_7812_3456_7812)
	testI64BinOp(t, op.I64Rotl, 0x1234_5678_1234_5678, 200, 0x3456_7812_3456_7812)
	testI64BinOp(t, op.I64Rotr, 0x1234_5678_1234_5678, 8, 0x7812_3456_7812_3456)
	testI64BinOp(t, op.I64Rotr, 0x1234_5678_1234_5678, 200, 0x7812_3456_7812_3456)
}

func TestF32Arithmetic(t *testing.T) {
	testF32UnOp(t, op.F32Abs, -1.5, 1.5)
	testF32UnOp(t, op.F32Neg, 1.5, -1.5)
	testF32UnOp(t, op.F32Ceil, 1.5, 2.0)
	testF32UnOp(t, op.F32Floor, 1.5, 1.0)
	testF32UnOp(t, op.F32Trunc, 1.5, 1.0)
	testF32UnOp(t, op.F32Nearest, 0.5, 0.0)
	testF32UnOp(t, op.F32Nearest, -0.5, 0.0)
	testF32UnOp(t, op.F32Nearest, 1.1, 1.0)
	testF32UnOp(t, op.F32Nearest, 1.5, 2.0)
	testF32UnOp(t, op.F32Nearest, 1.9, 2.0)
	testF32UnOp(t, op.F32Sqrt, 4.0, 2.0)
	testF32BinOp(t, op.F32Add, 3.0, 2.0, 5.0)
	testF32BinOp(t, op.F32Sub, 3.0, 2.0, 1.0)
	testF32BinOp(t, op.F32Mul, 3.0, 2.0, 6.0)
	testF32BinOp(t, op.F32Div, 3.0, 2.0, 1.5)
	testF32BinOp(t, op.F32Min, 3.0, 2.0, 2.0)
	testF32BinOp(t, op.F32Max, 3.0, 2.0, 3.0)
	testF32BinOp(t, op.F32CopySign, 3.0, 2.0, 3.0)
	testF32BinOp(t, op.F32CopySign, 3.0, -2.0, -3.0)
}

func TestF64Arithmetic(t *testing.T) {
	testF64UnOp(t, op.F64Abs, -1.5, 1.5)
	testF64UnOp(t, op.F64Neg, 1.5, -1.5)
	testF64UnOp(t, op.F64Ceil, 1.5, 2.0)
	testF64UnOp(t, op.F64Floor, 1.5, 1.0)
	testF64UnOp(t, op.F64Trunc, 1.5, 1.0)
	testF64UnOp(t, op.F64Nearest, 0.5, 0.0)
	testF64UnOp(t, op.F64Nearest, -0.5, 0.0)
	testF64UnOp(t, op.F64Nearest, 1.1, 1.0)
	testF64UnOp(t, op.F64Nearest, 1.5, 2.0)
	testF64UnOp(t, op.F64Nearest, 1.9, 2.0)
	testF64UnOp(t, op.F64Sqrt, 4.0, 2.0)
	testF64BinOp(t, op.F64Add, 3.0, 2.0, 5.0)
	testF64BinOp(t, op.F64Sub, 3.0, 2.0, 1.0)
	testF64BinOp(t, op.F64Mul, 3.0, 2.0, 6.0)
	testF64BinOp(t, op.F64Div, 3.0, 2.0, 1.5)
	testF64BinOp(t, op.F64Min, 3.0, 2.0, 2.0)
	testF64BinOp(t, op.F64Max, 3.0, 2.0, 3.0)
	testF64BinOp(t, op.F64CopySign, 3.0, 2.0, 3.0)
	testF64BinOp(t, op.F64CopySign, 3.0, -2.0, -3.0)
}

func TestConversions(t *testing.T) {
	testUnOp(t, op.I32WrapI64, int64(0x7F7F_7F7F_7F7F_7F7F), int32(0x7F7F_7F7F))
	testUnOp(t, op.I32TruncF32S, float32(-1.5), int32(-1))
	testUnOp(t, op.I32TruncF32U, float32(1.5), int32(1)) // TODO
	testUnOp(t, op.I32TruncF64S, -1.5, int32(-1))
	testUnOp(t, op.I32TruncF64U, 1.5, int32(1)) // TODO
	testUnOp(t, op.I64ExtendI32S, int32(-1), int64(-1))
	testUnOp(t, op.I64ExtendI32U, int32(-1), int64(0xFFFF_FFFF))
	testUnOp(t, op.I64TruncF32S, float32(-1.5), int64(-1))
	testUnOp(t, op.I64TruncF32U, float32(1.5), int64(1)) // TODO
	testUnOp(t, op.I64TruncF64S, -1.5, int64(-1))
	testUnOp(t, op.I64TruncF64U, 1.5, int64(1)) // TODO
	testUnOp(t, op.F32ConvertI32S, int32(-1), float32(-1.0))
	testUnOp(t, op.F32ConvertI32U, int32(-1), float32(4.2949673e+09))
	testUnOp(t, op.F32ConvertI64S, int64(-1), float32(-1.0))
	testUnOp(t, op.F32ConvertI64U, int64(-1), float32(1.8446744e+19))
	testUnOp(t, op.F32DemoteF64, 1.5, float32(1.5))
	testUnOp(t, op.F64ConvertI32S, int32(-1), -1.0)
	testUnOp(t, op.F64ConvertI32U, int32(-1), 4.294967295e+09)
	testUnOp(t, op.F64ConvertI64S, int64(-1), -1.0)
	testUnOp(t, op.F64ConvertI64U, int64(-1), 1.8446744073709552e+19)
	testUnOp(t, op.F64PromoteF32, float32(1.5), 1.5)
	testUnOp(t, op.I32ReinterpretF32, float32(1.5), int32(0x3FC0_0000))
	testUnOp(t, op.I64ReinterpretF64, 1.5, int64(0x3FF8_0000_0000_0000))
	testUnOp(t, op.F32ReinterpretI32, int32(0x3FC0_0000), float32(1.5))
	testUnOp(t, op.F64ReinterpretI64, int64(0x3FF8_0000_0000_0000), 1.5)
}

func testI32UnOp(t *testing.T, opcode byte, b, c int32) {
	testI32BinOp(t, opcode, 0, b, c)
}
func testI64UnOp(t *testing.T, opcode byte, b, c int64) {
	testI64BinOp(t, opcode, 0, b, c)
}
func testF32UnOp(t *testing.T, opcode byte, b, c float32) {
	testF32BinOp(t, opcode, 0, b, c)
}
func testF64UnOp(t *testing.T, opcode byte, b, c float64) {
	testF64BinOp(t, opcode, 0, b, c)
}

func testI32BinOp(t *testing.T, opcode byte, a, b, c int32) {
	testBinOp(t, opcode, a, b, c)
}
func testI64BinOp(t *testing.T, opcode byte, a, b, c int64) {
	testBinOp(t, opcode, a, b, c)
}
func testF32BinCmp(t *testing.T, opcode byte, a, b float32, c int32) {
	testBinOp(t, opcode, a, b, c)
}
func testF64BinCmp(t *testing.T, opcode byte, a, b float64, c int32) {
	testBinOp(t, opcode, a, b, c)
}
func testF32BinOp(t *testing.T, opcode byte, a, b, c float32) {
	testBinOp(t, opcode, a, b, c)
}
func testF64BinOp(t *testing.T, opcode byte, a, b, c float64) {
	testBinOp(t, opcode, a, b, c)
}

func testUnOp(t *testing.T, opcode byte, b, c interface{}) {
	testBinOp(t, opcode, int32(0), b, c)
}
func testBinOp(t *testing.T, opcode byte, a, b, c interface{}) {
	vm := &Vm{}
	pushVal(vm, a)
	pushVal(vm, b)
	InstrTable[opcode](vm, nil)
	require.Equal(t, c, popVal(vm, c))
}

func pushVal(vm *Vm, val interface{}) {
	switch x := val.(type) {
	case int32:
		vm.PushS32(x)
	case int64:
		vm.PushS64(x)
	case float32:
		vm.PushF32(x)
	case float64:
		vm.PushF64(x)
	default:
		panic(fmt.Errorf("wrong type: %v", val))
	}
}
func popVal(vm *Vm, typeInfo interface{}) interface{} {
	switch typeInfo.(type) {
	case int32:
		return vm.PopS32()
	case int64:
		return vm.PopS64()
	case float32:
		return vm.PopF32()
	case float64:
		return vm.PopF64()
	default:
		panic(fmt.Errorf("wrong type: %v", typeInfo))
	}
}
