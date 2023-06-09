package op

var Opnames = make([]string, 256)

func init() {
	Opnames[Unreachable] = "unreachable"
	Opnames[Nop] = "nop"
	Opnames[Block] = "block"
	Opnames[Loop] = "loop"
	Opnames[If] = "if"
	Opnames[Else_] = "else"
	Opnames[End_] = "end"
	Opnames[Br] = "br"
	Opnames[BrIf] = "br_if"
	Opnames[BrTable] = "br_table"
	Opnames[Return] = "return"
	Opnames[Call] = "call"
	Opnames[CallIndirect] = "call_indirect"
	Opnames[Drop] = "drop"
	Opnames[Select] = "select"
	Opnames[LocalGet] = "local.get"
	Opnames[LocalSet] = "local.set"
	Opnames[LocalTee] = "local.tee"
	Opnames[GlobalGet] = "global.get"
	Opnames[GlobalSet] = "global.set"
	Opnames[I32Load] = "i32.load"
	Opnames[I64Load] = "i64.load"
	Opnames[F32Load] = "f32.load"
	Opnames[F64Load] = "f64.load"
	Opnames[I32Load8S] = "i32.load8_s"
	Opnames[I32Load8U] = "i32.load8_u"
	Opnames[I32Load16S] = "i32.load16_s"
	Opnames[I32Load16U] = "i32.load16_u"
	Opnames[I64Load8S] = "i64.load8_s"
	Opnames[I64Load8U] = "i64.load8_u"
	Opnames[I64Load16S] = "i64.load16_s"
	Opnames[I64Load16U] = "i64.load16_u"
	Opnames[I64Load32S] = "i64.load32_s"
	Opnames[I64Load32U] = "i64.load32_u"
	Opnames[I32Store] = "i32.store"
	Opnames[I64Store] = "i64.store"
	Opnames[F32Store] = "f32.store"
	Opnames[F64Store] = "f64.store"
	Opnames[I32Store8] = "i32.store8"
	Opnames[I32Store16] = "i32.store16"
	Opnames[I64Store8] = "i64.store8"
	Opnames[I64Store16] = "i64.store16"
	Opnames[I64Store32] = "i64.store32"
	Opnames[MemorySize] = "memory.size"
	Opnames[MemoryGrow] = "memory.grow"
	Opnames[I32Const] = "i32.const"
	Opnames[I64Const] = "i64.const"
	Opnames[F32Const] = "f32.const"
	Opnames[F64Const] = "f64.const"
	Opnames[I32Eqz] = "i32.eqz"
	Opnames[I32Eq] = "i32.eq"
	Opnames[I32Ne] = "i32.ne"
	Opnames[I32LtS] = "i32.lt_s"
	Opnames[I32LtU] = "i32.lt_u"
	Opnames[I32GtS] = "i32.gt_s"
	Opnames[I32GtU] = "i32.gt_u"
	Opnames[I32LeS] = "i32.le_s"
	Opnames[I32LeU] = "i32.le_u"
	Opnames[I32GeS] = "i32.ge_s"
	Opnames[I32GeU] = "i32.ge_u"
	Opnames[I64Eqz] = "i64.eqz"
	Opnames[I64Eq] = "i64.eq"
	Opnames[I64Ne] = "i64.ne"
	Opnames[I64LtS] = "i64.lt_s"
	Opnames[I64LtU] = "i64.lt_u"
	Opnames[I64GtS] = "i64.gt_s"
	Opnames[I64GtU] = "i64.gt_u"
	Opnames[I64LeS] = "i64.le_s"
	Opnames[I64LeU] = "i64.le_u"
	Opnames[I64GeS] = "i64.ge_s"
	Opnames[I64GeU] = "i64.ge_u"
	Opnames[F32Eq] = "f32.eq"
	Opnames[F32Ne] = "f32.ne"
	Opnames[F32Lt] = "f32.lt"
	Opnames[F32Gt] = "f32.gt"
	Opnames[F32Le] = "f32.le"
	Opnames[F32Ge] = "f32.ge"
	Opnames[F64Eq] = "f64.eq"
	Opnames[F64Ne] = "f64.ne"
	Opnames[F64Lt] = "f64.lt"
	Opnames[F64Gt] = "f64.gt"
	Opnames[F64Le] = "f64.le"
	Opnames[F64Ge] = "f64.ge"
	Opnames[I32Clz] = "i32.clz"
	Opnames[I32Ctz] = "i32.ctz"
	Opnames[I32PopCnt] = "i32.popcnt"
	Opnames[I32Add] = "i32.add"
	Opnames[I32Sub] = "i32.sub"
	Opnames[I32Mul] = "i32.mul"
	Opnames[I32DivS] = "i32.div_s"
	Opnames[I32DivU] = "i32.div_u"
	Opnames[I32RemS] = "i32.rem_s"
	Opnames[I32RemU] = "i32.rem_u"
	Opnames[I32And] = "i32.and"
	Opnames[I32Or] = "i32.or"
	Opnames[I32Xor] = "i32.xor"
	Opnames[I32Shl] = "i32.shl"
	Opnames[I32ShrS] = "i32.shr_s"
	Opnames[I32ShrU] = "i32.shr_u"
	Opnames[I32Rotl] = "i32.rotl"
	Opnames[I32Rotr] = "i32.rotr"
	Opnames[I64Clz] = "i64.clz"
	Opnames[I64Ctz] = "i64.ctz"
	Opnames[I64PopCnt] = "i64.popcnt"
	Opnames[I64Add] = "i64.add"
	Opnames[I64Sub] = "i64.sub"
	Opnames[I64Mul] = "i64.mul"
	Opnames[I64DivS] = "i64.div_s"
	Opnames[I64DivU] = "i64.div_u"
	Opnames[I64RemS] = "i64.rem_s"
	Opnames[I64RemU] = "i64.rem_u"
	Opnames[I64And] = "i64.and"
	Opnames[I64Or] = "i64.or"
	Opnames[I64Xor] = "i64.xor"
	Opnames[I64Shl] = "i64.shl"
	Opnames[I64ShrS] = "i64.shr_s"
	Opnames[I64ShrU] = "i64.shr_u"
	Opnames[I64Rotl] = "i64.rotl"
	Opnames[I64Rotr] = "i64.rotr"
	Opnames[F32Abs] = "f32.abs"
	Opnames[F32Neg] = "f32.neg"
	Opnames[F32Ceil] = "f32.ceil"
	Opnames[F32Floor] = "f32.floor"
	Opnames[F32Trunc] = "f32.trunc"
	Opnames[F32Nearest] = "f32.nearest"
	Opnames[F32Sqrt] = "f32.sqrt"
	Opnames[F32Add] = "f32.add"
	Opnames[F32Sub] = "f32.sub"
	Opnames[F32Mul] = "f32.mul"
	Opnames[F32Div] = "f32.div"
	Opnames[F32Min] = "f32.min"
	Opnames[F32Max] = "f32.max"
	Opnames[F32CopySign] = "f32.copysign"
	Opnames[F64Abs] = "f64.abs"
	Opnames[F64Neg] = "f64.neg"
	Opnames[F64Ceil] = "f64.ceil"
	Opnames[F64Floor] = "f64.floor"
	Opnames[F64Trunc] = "f64.trunc"
	Opnames[F64Nearest] = "f64.nearest"
	Opnames[F64Sqrt] = "f64.sqrt"
	Opnames[F64Add] = "f64.add"
	Opnames[F64Sub] = "f64.sub"
	Opnames[F64Mul] = "f64.mul"
	Opnames[F64Div] = "f64.div"
	Opnames[F64Min] = "f64.min"
	Opnames[F64Max] = "f64.max"
	Opnames[F64CopySign] = "f64.copysign"
	Opnames[I32WrapI64] = "i32.wrap_i64"
	Opnames[I32TruncF32S] = "i32.trunc_f32_s"
	Opnames[I32TruncF32U] = "i32.trunc_f32_u"
	Opnames[I32TruncF64S] = "i32.trunc_f64_s"
	Opnames[I32TruncF64U] = "i32.trunc_f64_u"
	Opnames[I64ExtendI32S] = "i64.extend_i32_s"
	Opnames[I64ExtendI32U] = "i64.extend_i32_u"
	Opnames[I64TruncF32S] = "i64.trunc_f32_s"
	Opnames[I64TruncF32U] = "i64.trunc_f32_u"
	Opnames[I64TruncF64S] = "i64.trunc_f64_s"
	Opnames[I64TruncF64U] = "i64.trunc_f64_u"
	Opnames[F32ConvertI32S] = "f32.convert_i32_s"
	Opnames[F32ConvertI32U] = "f32.convert_i32_u"
	Opnames[F32ConvertI64S] = "f32.convert_i64_s"
	Opnames[F32ConvertI64U] = "f32.convert_i64_u"
	Opnames[F32DemoteF64] = "f32.demote_f64"
	Opnames[F64ConvertI32S] = "f64.convert_i32_s"
	Opnames[F64ConvertI32U] = "f64.convert_i32_u"
	Opnames[F64ConvertI64S] = "f64.convert_i64_s"
	Opnames[F64ConvertI64U] = "f64.convert_i64_u"
	Opnames[F64PromoteF32] = "f64.promote_f32"
	Opnames[I32ReinterpretF32] = "i32.reinterpret_f32"
	Opnames[I64ReinterpretF64] = "i64.reinterpret_f64"
	Opnames[F32ReinterpretI32] = "f32.reinterpret_i32"
	Opnames[F64ReinterpretI64] = "f64.reinterpret_i64"
	Opnames[I32Extend8S] = "i32.extend8_s"
	Opnames[I32Extend16S] = "i32.extend16_s"
	Opnames[I64Extend8S] = "i64.extend8_s"
	Opnames[I64Extend16S] = "i64.extend16_s"
	Opnames[I64Extend32S] = "i64.extend32_s"
	Opnames[TruncSat] = "trunc_sat"
}
