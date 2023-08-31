package blockChain

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
	state := GetStateInstance()
	return state.GetState(key), nil
}

func (stub *Stub) SetStringState(key string, value string) {
	state := GetStateInstance()
	state.SetState(key, value)
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
