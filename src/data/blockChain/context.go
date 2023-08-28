package blockChain

type Context struct {
	//clientIdentity string
	stub Stub
	//transactionId
	//transactionTimestamp
}

func (ctx *Context) GetStub() Stub {
	return ctx.stub
}
