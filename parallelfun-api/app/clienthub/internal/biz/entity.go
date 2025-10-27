package biz

type ClientInfo struct {
	ClientId string
	Target   string
	Secret   string
}

type RpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}
