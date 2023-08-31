package blockChain

import (
	"fmt"
	"testing"
)

func TestCreateEvidenceJson(t *testing.T) {
	CreateEvidenceJson()
}

func TestAddEvidence(t *testing.T) {
	data := `{"evidenceId":"123","uploaderSign":"abc","evidenceTxId":"xyz","content":"Sample content"}`
	stub := Stub{txId: "12346", txTimestamp: "000001"}
	context := Context{stub: stub}
	AddEvidence(&context, data)
	state := GetStateInstance()
	res := state.GetState("Evi_123")
	fmt.Println(res)
}
