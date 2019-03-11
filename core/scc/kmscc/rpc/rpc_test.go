package rpc

import (
	"log"
	"testing"
)

type TestEchoRequest struct {
	Msg string `json:"msg"`
}

type TestEchoResponse struct {
	Msg string `json:"msg"`
}

func echo(ctx *JsonRpcContext, req *TestEchoRequest, resp *TestEchoResponse) *JsonRpcError {
	resp.Msg = req.Msg
	return nil
}

func TestHandle(t *testing.T) {
	rpc := NewJsonRpc()
	rpc.Register("echo", echo)

	req := `{"jsonrpc":"2.0", "method":"echo", "params":{"msg":"hello"}}`
	resp := `{"jsonrpc":"2.0","result":{"msg":"hello"},"id":null}`
	result := rpc.Handle([]byte(req))
	if string(result) != resp {
		log.Printf("resp: %v", string(result))
		t.FailNow()
	}
}

func TestParamError(t *testing.T) {
	rpc := NewJsonRpc()
	rpc.Register("echo", echo)

	req := `{"jsonrpc":"1.0", "method":"echo", "params":{"msg":"hello"}}`
	resp := `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params"},"id":null}`
	result := rpc.Handle([]byte(req))
	if string(result) != resp {
		log.Printf("resp: %v", string(result))
		t.FailNow()
	}
}

func BenchmarkHandle(b *testing.B) {
	rpc := NewJsonRpc()
	rpc.Register("echo", echo)

	req := []byte(`{"jsonrpc":"1.0", "method":"echo", "params":{"msg":"hello"}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rpc.Handle(req)
	}
}
