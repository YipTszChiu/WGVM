package interpreter

import (
	"math"
	"math/bits"
)

// const
func i32Const(vm *Vm, args interface{}) {
	vm.PushS32(args.(int32))
}
func i64Const(vm *Vm, args interface{}) {
	vm.PushS64(args.(int64))
}
func f32Const(vm *Vm, args interface{}) {
	vm.PushF32(args.(float32))
}
func f64Const(vm *Vm, args interface{}) {
	vm.PushF64(args.(float64))
}

// i32 test & rel
func i32Eqz(vm *Vm, _ interface{}) {
	vm.PushBool(vm.PopU32() == 0)
}
func i32Eq(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushBool(v1 == v2)
}
func i32Ne(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushBool(v1 != v2)
}
func i32LtS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS32(), vm.PopS32()
	vm.PushBool(v1 < v2)
}
func i32LtU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushBool(v1 < v2)
}
func i32GtS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS32(), vm.PopS32()
	vm.PushBool(v1 > v2)
}
func i32GtU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushBool(v1 > v2)
}
func i32LeS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS32(), vm.PopS32()
	vm.PushBool(v1 <= v2)
}
func i32LeU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushBool(v1 <= v2)
}
func i32GeS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS32(), vm.PopS32()
	vm.PushBool(v1 >= v2)
}
func i32GeU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushBool(v1 >= v2)
}

// i64 test & rel
func i64Eqz(vm *Vm, _ interface{}) {
	vm.PushBool(vm.PopU64() == 0)
}
func i64Eq(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushBool(v1 == v2)
}
func i64Ne(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushBool(v1 != v2)
}
func i64LtS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS64(), vm.PopS64()
	vm.PushBool(v1 < v2)
}
func i64LtU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushBool(v1 < v2)
}
func i64GtS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS64(), vm.PopS64()
	vm.PushBool(v1 > v2)
}
func i64GtU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushBool(v1 > v2)
}
func i64LeS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS64(), vm.PopS64()
	vm.PushBool(v1 <= v2)
}
func i64LeU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushBool(v1 <= v2)
}
func i64GeS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS64(), vm.PopS64()
	vm.PushBool(v1 >= v2)
}
func i64GeU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushBool(v1 >= v2)
}

// f32 rel
func f32Eq(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushBool(v1 == v2)
}
func f32Ne(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushBool(v1 != v2)
}
func f32Lt(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushBool(v1 < v2)
}
func f32Gt(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushBool(v1 > v2)
}
func f32Le(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushBool(v1 <= v2)
}
func f32Ge(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushBool(v1 >= v2)
}

// f64 rel
func f64Eq(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushBool(v1 == v2)
}
func f64Ne(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushBool(v1 != v2)
}
func f64Lt(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushBool(v1 < v2)
}
func f64Gt(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushBool(v1 > v2)
}
func f64Le(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushBool(v1 <= v2)
}
func f64Ge(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushBool(v1 >= v2)
}

// i32 arithmetic & bitwise
func i32Clz(vm *Vm, _ interface{}) {
	vm.PushU32(uint32(bits.LeadingZeros32(vm.PopU32())))
}
func i32Ctz(vm *Vm, _ interface{}) {
	vm.PushU32(uint32(bits.TrailingZeros32(vm.PopU32())))
}
func i32PopCnt(vm *Vm, _ interface{}) {
	vm.PushU32(uint32(bits.OnesCount32(vm.PopU32())))
}
func i32Add(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 + v2)
}
func i32Sub(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 - v2)
}
func i32Mul(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 * v2)
}
func i32DivS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS32(), vm.PopS32()
	if v1 == math.MinInt32 && v2 == -1 {
		panic(ErrIntOverflow)
	}
	vm.PushS32(v1 / v2)
}
func i32DivU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 / v2)
}
func i32RemS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS32(), vm.PopS32()
	vm.PushS32(v1 % v2)
}
func i32RemU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 % v2)
}
func i32And(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 & v2)
}
func i32Or(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 | v2)
}
func i32Xor(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 ^ v2)
}
func i32Shl(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 << (v2 % 32))
}
func i32ShrS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopS32()
	vm.PushS32(v1 >> (v2 % 32))
}
func i32ShrU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(v1 >> (v2 % 32))
}
func i32Rotl(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(bits.RotateLeft32(v1, int(v2)))
}
func i32Rotr(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU32(), vm.PopU32()
	vm.PushU32(bits.RotateLeft32(v1, -int(v2)))
}

// i64 arithmetic & bitwise
func i64Clz(vm *Vm, _ interface{}) {
	vm.PushU64(uint64(bits.LeadingZeros64(vm.PopU64())))
}
func i64Ctz(vm *Vm, _ interface{}) {
	vm.PushU64(uint64(bits.TrailingZeros64(vm.PopU64())))
}
func i64PopCnt(vm *Vm, _ interface{}) {
	vm.PushU64(uint64(bits.OnesCount64(vm.PopU64())))
}
func i64Add(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 + v2)
}
func i64Sub(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 - v2)
}
func i64Mul(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 * v2)
}
func i64DivS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS64(), vm.PopS64()
	if v1 == math.MinInt64 && v2 == -1 {
		panic(ErrIntOverflow)
	}
	vm.PushS64(v1 / v2)
}
func i64DivU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 / v2)
}
func i64RemS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopS64(), vm.PopS64()
	vm.PushS64(v1 % v2)
}
func i64RemU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 % v2)
}
func i64And(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 & v2)
}
func i64Or(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 | v2)
}
func i64Xor(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 ^ v2)
}
func i64Shl(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 << (v2 % 64))
}
func i64ShrS(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopS64()
	vm.PushS64(v1 >> (v2 % 64))
}
func i64ShrU(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(v1 >> (v2 % 64))
}
func i64Rotl(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(bits.RotateLeft64(v1, int(v2)))
}
func i64Rotr(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopU64(), vm.PopU64()
	vm.PushU64(bits.RotateLeft64(v1, -int(v2)))
}

// f32 arithmetic
func f32Abs(vm *Vm, _ interface{}) {
	vm.PushF32(float32(math.Abs(float64(vm.PopF32()))))
}
func f32Neg(vm *Vm, _ interface{}) {
	vm.PushF32(-vm.PopF32())
}
func f32Ceil(vm *Vm, _ interface{}) {
	vm.PushF32(float32(math.Ceil(float64(vm.PopF32()))))
}
func f32Floor(vm *Vm, _ interface{}) {
	vm.PushF32(float32(math.Floor(float64(vm.PopF32()))))
}
func f32Trunc(vm *Vm, _ interface{}) {
	vm.PushF32(float32(math.Trunc(float64(vm.PopF32()))))
}
func f32Nearest(vm *Vm, _ interface{}) {
	vm.PushF32(float32(math.RoundToEven(float64(vm.PopF32()))))
}
func f32Sqrt(vm *Vm, _ interface{}) {
	vm.PushF32(float32(math.Sqrt(float64(vm.PopF32()))))
}
func f32Add(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushF32(v1 + v2)
}
func f32Sub(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushF32(v1 - v2)
}
func f32Mul(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushF32(v1 * v2)
}
func f32Div(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushF32(v1 / v2)
}
func f32Min(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	v1NaN := math.IsNaN(float64(v1))
	v2NaN := math.IsNaN(float64(v2))
	if v1NaN && !v2NaN {
		vm.PushF32(v1)
		return
	} else if v2NaN && !v1NaN {
		vm.PushF32(v2)
		return
	}
	vm.PushF32(float32(math.Min(float64(v1), float64(v2))))
}
func f32Max(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	v1NaN := math.IsNaN(float64(v1))
	v2NaN := math.IsNaN(float64(v2))
	if v1NaN && !v2NaN {
		vm.PushF32(v1)
		return
	} else if v2NaN && !v1NaN {
		vm.PushF32(v2)
		return
	}
	vm.PushF32(float32(math.Max(float64(v1), float64(v2))))
}
func f32CopySign(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF32(), vm.PopF32()
	vm.PushF32(float32(math.Copysign(float64(v1), float64(v2))))
}

// f64 arithmetic
func f64Abs(vm *Vm, _ interface{}) {
	vm.PushF64(math.Abs(vm.PopF64()))
}
func f64Neg(vm *Vm, _ interface{}) {
	vm.PushF64(-vm.PopF64())
}
func f64Ceil(vm *Vm, _ interface{}) {
	vm.PushF64(math.Ceil(vm.PopF64()))
}
func f64Floor(vm *Vm, _ interface{}) {
	vm.PushF64(math.Floor(vm.PopF64()))
}
func f64Trunc(vm *Vm, _ interface{}) {
	vm.PushF64(math.Trunc(vm.PopF64()))
}
func f64Nearest(vm *Vm, _ interface{}) {
	vm.PushF64(math.RoundToEven(vm.PopF64()))
}
func f64Sqrt(vm *Vm, _ interface{}) {
	vm.PushF64(math.Sqrt(vm.PopF64()))
}
func f64Add(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushF64(v1 + v2)
}
func f64Sub(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushF64(v1 - v2)
}
func f64Mul(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushF64(v1 * v2)
}
func f64Div(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushF64(v1 / v2)
}
func f64Min(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	v1NaN := math.IsNaN(v1)
	v2NaN := math.IsNaN(v2)
	if v1NaN && !v2NaN {
		vm.PushF64(v1)
		return
	} else if v2NaN && !v1NaN {
		vm.PushF64(v2)
		return
	}
	vm.PushF64(math.Min(v1, v2))
}
func f64Max(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	v1NaN := math.IsNaN(v1)
	v2NaN := math.IsNaN(v2)
	if v1NaN && !v2NaN {
		vm.PushF64(v1)
		return
	} else if v2NaN && !v1NaN {
		vm.PushF64(v2)
		return
	}
	vm.PushF64(math.Max(v1, v2))
}
func f64CopySign(vm *Vm, _ interface{}) {
	v2, v1 := vm.PopF64(), vm.PopF64()
	vm.PushF64(math.Copysign(v1, v2))
}

// conversions
func i32WrapI64(vm *Vm, _ interface{}) {
	vm.PushU32(uint32(vm.PopU64()))
}
func i32TruncF32S(vm *Vm, _ interface{}) {
	f := math.Trunc(float64(vm.PopF32()))
	if f > math.MaxInt32 || f < math.MinInt32 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushS32(int32(f))
}
func i32TruncF32U(vm *Vm, _ interface{}) {
	f := math.Trunc(float64(vm.PopF32()))
	if f > math.MaxUint32 || f < 0 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushU32(uint32(f))
}
func i32TruncF64S(vm *Vm, _ interface{}) {
	f := math.Trunc(vm.PopF64())
	if f > math.MaxInt32 || f < math.MinInt32 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushS32(int32(f))
}
func i32TruncF64U(vm *Vm, _ interface{}) {
	f := math.Trunc(vm.PopF64())
	if f > math.MaxUint32 || f < 0 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushU32(uint32(f))
}
func i64ExtendI32S(vm *Vm, _ interface{}) {
	vm.PushS64(int64(vm.PopS32()))
}
func i64ExtendI32U(vm *Vm, _ interface{}) {
	vm.PushU64(uint64(vm.PopU32()))
}
func i64TruncF32S(vm *Vm, _ interface{}) {
	f := math.Trunc(float64(vm.PopF32()))
	if f >= math.MaxInt64 || f < math.MinInt64 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushS64(int64(f))
}
func i64TruncF32U(vm *Vm, _ interface{}) {
	f := math.Trunc(float64(vm.PopF32()))
	if f >= math.MaxUint64 || f < 0 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushU64(uint64(f))
}
func i64TruncF64S(vm *Vm, _ interface{}) {
	f := math.Trunc(vm.PopF64())
	if f >= math.MaxInt64 || f < math.MinInt64 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushS64(int64(f))
}
func i64TruncF64U(vm *Vm, _ interface{}) {
	f := math.Trunc(vm.PopF64())
	if f >= math.MaxUint64 || f < 0 {
		panic(ErrIntOverflow)
	}
	if math.IsNaN(f) {
		panic(ErrConvertToInt)
	}
	vm.PushU64(uint64(f))
}
func f32ConvertI32S(vm *Vm, _ interface{}) {
	vm.PushF32(float32(vm.PopS32()))
}
func f32ConvertI32U(vm *Vm, _ interface{}) {
	vm.PushF32(float32(vm.PopU32()))
}
func f32ConvertI64S(vm *Vm, _ interface{}) {
	vm.PushF32(float32(vm.PopS64()))
}
func f32ConvertI64U(vm *Vm, _ interface{}) {
	vm.PushF32(float32(vm.PopU64()))
}
func f32DemoteF64(vm *Vm, _ interface{}) {
	vm.PushF32(float32(vm.PopF64()))
}
func f64ConvertI32S(vm *Vm, _ interface{}) {
	vm.PushF64(float64(vm.PopS32()))
}
func f64ConvertI32U(vm *Vm, _ interface{}) {
	vm.PushF64(float64(vm.PopU32()))
}
func f64ConvertI64S(vm *Vm, _ interface{}) {
	vm.PushF64(float64(vm.PopS64()))
}
func f64ConvertI64U(vm *Vm, _ interface{}) {
	vm.PushF64(float64(vm.PopU64()))
}
func f64PromoteF32(vm *Vm, _ interface{}) {
	vm.PushF64(float64(vm.PopF32()))
}
func i32ReinterpretF32(vm *Vm, _ interface{}) {
	//vm.PushU32(math.Float32bits(vm.PopF32()))
}
func i64ReinterpretF64(vm *Vm, _ interface{}) {
	//vm.PushU64(math.Float64bits(vm.PopF64()))
}
func f32ReinterpretI32(vm *Vm, _ interface{}) {
	//vm.PushF32(math.Float32frombits(vm.PopU32()))
}
func f64ReinterpretI64(vm *Vm, _ interface{}) {
	//vm.PushF64(math.Float64frombits(vm.PopU64()))
}

func i32Extend8S(vm *Vm, _ interface{}) {
	vm.PushS32(int32(int8(vm.PopS32())))
}
func i32Extend16S(vm *Vm, _ interface{}) {
	vm.PushS32(int32(int16(vm.PopS32())))
}
func i64Extend8S(vm *Vm, _ interface{}) {
	vm.PushS64(int64(int8(vm.PopS64())))
}
func i64Extend16S(vm *Vm, _ interface{}) {
	vm.PushS64(int64(int16(vm.PopS64())))
}
func i64Extend32S(vm *Vm, _ interface{}) {
	vm.PushS64(int64(int32(vm.PopS64())))
}

func truncSat(vm *Vm, args interface{}) {
	switch args.(byte) {
	case 0: // i32.trunc_sat_f32_s
		v := truncSatS(float64(vm.PopF32()), 32)
		vm.PushS32(int32(v))
	case 1: // i32.trunc_sat_f32_u
		v := truncSatU(float64(vm.PopF32()), 32)
		vm.PushU32(uint32(v))
	case 2: // i32.trunc_sat_f64_s
		v := truncSatS(vm.PopF64(), 32)
		vm.PushS32(int32(v))
	case 3: // i32.trunc_sat_f64_u
		v := truncSatU(vm.PopF64(), 32)
		vm.PushU32(uint32(v))
	case 4: // i64.trunc_sat_f32_s
		v := truncSatS(float64(vm.PopF32()), 64)
		vm.PushS64(v)
	case 5: // i64.trunc_sat_f32_u
		v := truncSatU(float64(vm.PopF32()), 64)
		vm.PushU64(v)
	case 6: // i64.trunc_sat_f64_s
		v := truncSatS(vm.PopF64(), 64)
		vm.PushS64(v)
	case 7: // i64.trunc_sat_f64_u
		v := truncSatU(vm.PopF64(), 64)
		vm.PushU64(v)
	default:
		panic("unreachable")
	}
}

func truncSatU(z float64, n int) uint64 {
	if math.IsNaN(z) {
		return 0
	}
	if math.IsInf(z, -1) {
		return 0
	}
	max := (uint64(1) << n) - 1
	if math.IsInf(z, 1) {
		return max
	}
	if x := math.Trunc(z); x < 0 {
		return 0
	} else if x >= float64(max) {
		return max
	} else {
		return uint64(x)
	}
}
func truncSatS(z float64, n int) int64 {
	if math.IsNaN(z) {
		return 0
	}
	min := -(int64(1) << (n - 1))
	max := (int64(1) << (n - 1)) - 1
	if math.IsInf(z, -1) {
		return min
	}
	if math.IsInf(z, 1) {
		return max
	}
	if x := math.Trunc(z); x < float64(min) {
		return min
	} else if x >= float64(max) {
		return max
	} else {
		return int64(x)
	}
}
