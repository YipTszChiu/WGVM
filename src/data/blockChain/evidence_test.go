package blockChain

import (
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
}
