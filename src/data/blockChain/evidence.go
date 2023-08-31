package blockChain

import (
	"encoding/json"
	"fmt"
)

type EvidenceInfo struct {
	EvidenceId   string `json:"evidenceId"`
	UploaderSign string `json:"uploaderSign"`
	EvidenceTxId string `json:"evidenceTxId"`
	Content      string `json:"content"`
}

type CreateResult struct {
	EvidenceId  string `json:"evidenceId,omitempty"`
	TxId        string `json:"txId,omitempty"`
	TxTimestamp string `json:"txTimestamp,omitempty"`
}

type QueryResult struct {
	Key      string `json:"key,omitempty"`
	Record   string `json:"record,omitempty"`
	Bookmark string `json:"bookmark,omitempty"`
}

type Error struct {
	message string
	payload string
}

func AddEvidence(context *Context, data string) *CreateResult {
	evidenceinfo := &EvidenceInfo{}
	err := json.Unmarshal([]byte(data), &evidenceinfo)
	stub := context.GetStub()
	key := formatEvidenceKey(evidenceinfo.EvidenceId)
	_, err = stub.GetStringState(key)
	if err != nil {
		panic(Error{
			message: "Evidence " + evidenceinfo.EvidenceId + " already exists",
			payload: "EVIDENCE_ALREADY_EXIST",
		})

	}
	txId := stub.GetTxId()
	timeStamp := stub.GetTxTimestamp()
	evidenceinfo.EvidenceTxId = txId
	evidenceStateBytes, _ := json.Marshal(evidenceinfo)
	stub.SetStringState(key, string(evidenceStateBytes))
	return &CreateResult{
		EvidenceId:  evidenceinfo.EvidenceId,
		TxId:        txId,
		TxTimestamp: timeStamp,
	}
}

func formatEvidenceKey(evidenceId string) string {
	return "Evi_" + evidenceId
}

func CreateEvidenceJson() {
	evidence := EvidenceInfo{
		EvidenceId:   "123",
		UploaderSign: "abc",
		EvidenceTxId: "xyz",
		Content:      "Sample content",
	}

	// 将 EvidenceInfo 转换为 JSON 格式
	evidenceJSON, err := json.Marshal(evidence)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印 JSON 格式的 EvidenceInfo
	fmt.Println(string(evidenceJSON))
}
