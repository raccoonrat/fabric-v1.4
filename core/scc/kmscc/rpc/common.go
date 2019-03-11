package rpc

var InvalidParamsError *JsonRpcError = &JsonRpcError{
	Code:    -32602,
	Message: "Invalid params",
}
var MethodNotFoundError *JsonRpcError = &JsonRpcError{
	Code:    -32601,
	Message: "Method not found",
}

func ParseError(err error) *JsonRpcError {
	return &JsonRpcError{
		Code:    -32700,
		Message: "Parse error",
		Data:    err.Error(),
	}
}
