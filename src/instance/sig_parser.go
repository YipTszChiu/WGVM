package instance

import (
	"WGVM/src/common"
	"strings"
)

// name(params)->(results)
func parseNameAndSig(nameAndSig string) (string, common.FuncType) {
	idxOfLPar := strings.IndexByte(nameAndSig, '(')
	name := nameAndSig[:idxOfLPar]
	sig := nameAndSig[idxOfLPar:]
	return name, parseSig(sig)
}

func parseSig(sig string) common.FuncType {
	paramsAndResults := strings.SplitN(sig, "->", 2)
	return common.FuncType{
		ParamTypes:  parseValTypes(paramsAndResults[0]),
		ResultTypes: parseValTypes(paramsAndResults[1]),
	}
}

func parseValTypes(list string) []common.ValType {
	list = strings.TrimSpace(list)
	list = list[1 : len(list)-1] // remove ()

	var valTypes []common.ValType
	for _, t := range strings.Split(list, ",") {
		switch strings.TrimSpace(t) {
		case "i32":
			valTypes = append(valTypes, common.ValTypeI32)
		case "i64":
			valTypes = append(valTypes, common.ValTypeI64)
		case "f32":
			valTypes = append(valTypes, common.ValTypeF32)
		case "f64":
			valTypes = append(valTypes, common.ValTypeF64)
		}
	}
	return valTypes
}
