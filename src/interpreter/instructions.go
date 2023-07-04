package interpreter

import (
	"WGVM/src/common/op"
)

type instrFn = func(vm *Vm, args interface{})

var InstrTable = make([]instrFn, 256)

func init() {
	//InstrTable[op.Call] = call // hack!
	InstrTable[op.Drop] = drop
	InstrTable[op.Select] = _select
	InstrTable[op.I32Const] = i32Const
	InstrTable[op.I64Const] = i64Const
	InstrTable[op.F32Const] = f32Const
	InstrTable[op.F64Const] = f64Const
	InstrTable[op.I32Eqz] = i32Eqz
	InstrTable[op.I32Eq] = i32Eq
	InstrTable[op.I32Ne] = i32Ne
	InstrTable[op.I32LtS] = i32LtS
	InstrTable[op.I32LtU] = i32LtU
	InstrTable[op.I32GtS] = i32GtS
	InstrTable[op.I32GtU] = i32GtU
	InstrTable[op.I32LeS] = i32LeS
	InstrTable[op.I32LeU] = i32LeU
	InstrTable[op.I32GeS] = i32GeS
	InstrTable[op.I32GeU] = i32GeU
	InstrTable[op.I64Eqz] = i64Eqz
	InstrTable[op.I64Eq] = i64Eq
	InstrTable[op.I64Ne] = i64Ne
	InstrTable[op.I64LtS] = i64LtS
	InstrTable[op.I64LtU] = i64LtU
	InstrTable[op.I64GtS] = i64GtS
	InstrTable[op.I64GtU] = i64GtU
	InstrTable[op.I64LeS] = i64LeS
	InstrTable[op.I64LeU] = i64LeU
	InstrTable[op.I64GeS] = i64GeS
	InstrTable[op.I64GeU] = i64GeU
	InstrTable[op.F32Eq] = f32Eq
	InstrTable[op.F32Ne] = f32Ne
	InstrTable[op.F32Lt] = f32Lt
	InstrTable[op.F32Gt] = f32Gt
	InstrTable[op.F32Le] = f32Le
	InstrTable[op.F32Ge] = f32Ge
	InstrTable[op.F64Eq] = f64Eq
	InstrTable[op.F64Ne] = f64Ne
	InstrTable[op.F64Lt] = f64Lt
	InstrTable[op.F64Gt] = f64Gt
	InstrTable[op.F64Le] = f64Le
	InstrTable[op.F64Ge] = f64Ge
	InstrTable[op.I32Clz] = i32Clz
	InstrTable[op.I32Ctz] = i32Ctz
	InstrTable[op.I32PopCnt] = i32PopCnt
	InstrTable[op.I32Add] = i32Add
	InstrTable[op.I32Sub] = i32Sub
	InstrTable[op.I32Mul] = i32Mul
	InstrTable[op.I32DivS] = i32DivS
	InstrTable[op.I32DivU] = i32DivU
	InstrTable[op.I32RemS] = i32RemS
	InstrTable[op.I32RemU] = i32RemU
	InstrTable[op.I32And] = i32And
	InstrTable[op.I32Or] = i32Or
	InstrTable[op.I32Xor] = i32Xor
	InstrTable[op.I32Shl] = i32Shl
	InstrTable[op.I32ShrS] = i32ShrS
	InstrTable[op.I32ShrU] = i32ShrU
	InstrTable[op.I32Rotl] = i32Rotl
	InstrTable[op.I32Rotr] = i32Rotr
	InstrTable[op.I64Clz] = i64Clz
	InstrTable[op.I64Ctz] = i64Ctz
	InstrTable[op.I64PopCnt] = i64PopCnt
	InstrTable[op.I64Add] = i64Add
	InstrTable[op.I64Sub] = i64Sub
	InstrTable[op.I64Mul] = i64Mul
	InstrTable[op.I64DivS] = i64DivS
	InstrTable[op.I64DivU] = i64DivU
	InstrTable[op.I64RemS] = i64RemS
	InstrTable[op.I64RemU] = i64RemU
	InstrTable[op.I64And] = i64And
	InstrTable[op.I64Or] = i64Or
	InstrTable[op.I64Xor] = i64Xor
	InstrTable[op.I64Shl] = i64Shl
	InstrTable[op.I64ShrS] = i64ShrS
	InstrTable[op.I64ShrU] = i64ShrU
	InstrTable[op.I64Rotl] = i64Rotl
	InstrTable[op.I64Rotr] = i64Rotr
	InstrTable[op.F32Abs] = f32Abs
	InstrTable[op.F32Neg] = f32Neg
	InstrTable[op.F32Ceil] = f32Ceil
	InstrTable[op.F32Floor] = f32Floor
	InstrTable[op.F32Trunc] = f32Trunc
	InstrTable[op.F32Nearest] = f32Nearest
	InstrTable[op.F32Sqrt] = f32Sqrt
	InstrTable[op.F32Add] = f32Add
	InstrTable[op.F32Sub] = f32Sub
	InstrTable[op.F32Mul] = f32Mul
	InstrTable[op.F32Div] = f32Div
	InstrTable[op.F32Min] = f32Min
	InstrTable[op.F32Max] = f32Max
	InstrTable[op.F32CopySign] = f32CopySign
	InstrTable[op.F64Abs] = f64Abs
	InstrTable[op.F64Neg] = f64Neg
	InstrTable[op.F64Ceil] = f64Ceil
	InstrTable[op.F64Floor] = f64Floor
	InstrTable[op.F64Trunc] = f64Trunc
	InstrTable[op.F64Nearest] = f64Nearest
	InstrTable[op.F64Sqrt] = f64Sqrt
	InstrTable[op.F64Add] = f64Add
	InstrTable[op.F64Sub] = f64Sub
	InstrTable[op.F64Mul] = f64Mul
	InstrTable[op.F64Div] = f64Div
	InstrTable[op.F64Min] = f64Min
	InstrTable[op.F64Max] = f64Max
	InstrTable[op.F64CopySign] = f64CopySign
	InstrTable[op.I32WrapI64] = i32WrapI64
	InstrTable[op.I32TruncF32S] = i32TruncF32S
	InstrTable[op.I32TruncF32U] = i32TruncF32U
	InstrTable[op.I32TruncF64S] = i32TruncF64S
	InstrTable[op.I32TruncF64U] = i32TruncF64U
	InstrTable[op.I64ExtendI32S] = i64ExtendI32S
	InstrTable[op.I64ExtendI32U] = i64ExtendI32U
	InstrTable[op.I64TruncF32S] = i64TruncF32S
	InstrTable[op.I64TruncF32U] = i64TruncF32U
	InstrTable[op.I64TruncF64S] = i64TruncF64S
	InstrTable[op.I64TruncF64U] = i64TruncF64U
	InstrTable[op.F32ConvertI32S] = f32ConvertI32S
	InstrTable[op.F32ConvertI32U] = f32ConvertI32U
	InstrTable[op.F32ConvertI64S] = f32ConvertI64S
	InstrTable[op.F32ConvertI64U] = f32ConvertI64U
	InstrTable[op.F32DemoteF64] = f32DemoteF64
	InstrTable[op.F64ConvertI32S] = f64ConvertI32S
	InstrTable[op.F64ConvertI32U] = f64ConvertI32U
	InstrTable[op.F64ConvertI64S] = f64ConvertI64S
	InstrTable[op.F64ConvertI64U] = f64ConvertI64U
	InstrTable[op.F64PromoteF32] = f64PromoteF32
	InstrTable[op.I32ReinterpretF32] = i32ReinterpretF32
	InstrTable[op.I64ReinterpretF64] = i64ReinterpretF64
	InstrTable[op.F32ReinterpretI32] = f32ReinterpretI32
	InstrTable[op.F64ReinterpretI64] = f64ReinterpretI64
	InstrTable[op.I32Extend8S] = i32Extend8S
	InstrTable[op.I32Extend16S] = i32Extend16S
	InstrTable[op.I64Extend8S] = i64Extend8S
	InstrTable[op.I64Extend16S] = i64Extend16S
	InstrTable[op.I64Extend32S] = i64Extend32S
	InstrTable[op.TruncSat] = truncSat
}