package blockChain

import "errors"

var State = map[string]string{}

type Stub struct {
	txId        string
	txTimestamp string
}

func (stub *Stub) GetTxTimestamp() string {
	return stub.txTimestamp
}

func (stub *Stub) SetTxTimestamp(txTimestamp string) {
	stub.txTimestamp = txTimestamp
}

func (stub *Stub) GetTxId() string {
	return stub.txId
}

func (stub *Stub) SetTxId(txId string) {
	stub.txId = txId
}

func (stub *Stub) GetStringState(key string) (string, error) {
	if v, ok := State[key]; ok {
		return v, nil
	}
	return "", errors.New("key exist")
}

func (stub *Stub) SetStringState(key string, value string) error {
	State[key] = value
	return nil
}

//	func (stub *stub) DelState(key string) error {
//		// Function signature
//	}
//
//	func (stub *stub) GetStateByRange(startKey, endKey string) (contractapi.StateQueryIteratorInterface, error) {
//		// Function signature
//	}
//
//	func (stub *stub) GetQueryResult(query string) (contractapi.StateQueryIteratorInterface, error) {
//		// Function signature
//	}
//
//	func (stub *stub) SetEvent(name string, payload []byte) error {
//		// Function signature
//	}
//
//	func (stub *stub) GetBinding() ([]byte, error) {
//		// Function signature
//	}
