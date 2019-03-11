package rpc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"gopkg.in/go-playground/validator.v9"
)

type JsonRpcContext struct {
	Request *JsonRpcRequest
}

type JsonRpcRequest struct {
	Version string          `json:"jsonrpc" validate:"eq=2.0"`
	Method  string          `json:"method" validate:"required"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      *string         `json:"id,omitempty"`
}

func (req *JsonRpcRequest) New() *JsonRpcRequest {
	return &JsonRpcRequest{
		Version: "2.0",
	}
}

func (req *JsonRpcRequest) clear() {
	req.Method = ""
	req.Params = json.RawMessage([]byte{})
	req.ID = nil
}

type JsonRpcResponse struct {
	Version string          `json:"jsonrpc" validate:"eq=2.0"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *JsonRpcError   `json:"error,omitempty"`
	ID      *string         `json:"id"`
}

func (resp *JsonRpcResponse) New() *JsonRpcResponse {
	return &JsonRpcResponse{
		Version: "2.0",
	}
}
func (resp *JsonRpcResponse) clear() {
	resp.Result = json.RawMessage([]byte{})
	resp.Error = nil
	resp.ID = nil
}

type JsonRpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type JsonRpc struct {
	handlers map[string]*HandlerSpec

	reqPool  sync.Pool
	respPool sync.Pool

	protoValidator *validator.Validate
}

type HandlerSpec struct {
	v *reflect.Value
}

func NewJsonRpc() *JsonRpc {
	rpc := &JsonRpc{
		handlers:       make(map[string]*HandlerSpec),
		protoValidator: validator.New(),
	}
	rpc.reqPool.New = func() interface{} {
		return &JsonRpcRequest{
			Version: "2.0",
		}
	}
	rpc.respPool.New = func() interface{} {
		return &JsonRpcResponse{
			Version: "2.0",
		}
	}
	return rpc
}

func (rpc *JsonRpc) Register(method string, handler interface{}) error {
	v := reflect.ValueOf(handler)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("handler must be a function")
	}
	if v.Type().NumIn() != 3 || v.Type().NumOut() != 1 {
		return fmt.Errorf("method signature is invalid")
	}

	hs := &HandlerSpec{
		v: &v,
	}
	rpc.handlers[method] = hs
	return nil
}
func (rpc *JsonRpc) allocRequest() *JsonRpcRequest {
	rpcReq := rpc.reqPool.Get().(*JsonRpcRequest)
	rpcReq.clear()
	return rpcReq
}
func (rpc *JsonRpc) recyleRequest(req *JsonRpcRequest) {
	rpc.reqPool.Put(req)
}
func (rpc *JsonRpc) allocResponse() *JsonRpcResponse {
	respResp := rpc.respPool.Get().(*JsonRpcResponse)
	respResp.clear()
	return respResp
}
func (rpc *JsonRpc) recyleResponse(resp *JsonRpcResponse) {
	rpc.respPool.Put(resp)
}
func (rpc *JsonRpc) Handle(req []byte) []byte {
	rpcReq := rpc.allocRequest()
	defer rpc.recyleRequest(rpcReq)

	err := json.Unmarshal(req, rpcReq)
	if err != nil {
		rpcErr := ParseError(err)
		return rpc.handleError(rpcReq, rpcErr)
	}

	err = rpc.protoValidator.Struct(rpcReq)
	if err != nil {
		return rpc.handleError(rpcReq, InvalidParamsError)
	}

	hs, ok := rpc.handlers[rpcReq.Method]
	if !ok {
		return rpc.handleError(rpcReq, MethodNotFoundError)
	}
	ctx := &JsonRpcContext{
		Request: rpcReq,
	}
	vreq := reflect.New(hs.v.Type().In(1).Elem())
	vresp := reflect.New(hs.v.Type().In(2).Elem())
	err = json.Unmarshal(rpcReq.Params, vreq.Interface())
	if err != nil {
		return rpc.handleError(rpcReq, InvalidParamsError)
	}
	args := []reflect.Value{
		reflect.ValueOf(ctx),
		vreq,
		vresp,
	}
	rets := hs.v.Call(args)
	rpcErr := rets[0].Interface().(*JsonRpcError)
	if err == nil {
		return rpc.handleSuccess(rpcReq, vresp.Interface())
	}
	return rpc.handleError(rpcReq, rpcErr)
}

func (rpc *JsonRpc) handleSuccess(req *JsonRpcRequest, data interface{}) []byte {
	rpcResp := rpc.allocResponse()
	defer rpc.recyleResponse(rpcResp)

	rpcResp.ID = req.ID

	payload, _ := json.Marshal(data)
	rpcResp.Result = payload

	result, _ := json.Marshal(rpcResp)
	return result
}

func (rpc *JsonRpc) handleError(req *JsonRpcRequest, rpcErr *JsonRpcError) []byte {
	rpcResp := rpc.allocResponse()
	defer rpc.recyleResponse(rpcResp)

	rpcResp.ID = req.ID
	rpcResp.Error = rpcErr

	result, _ := json.Marshal(rpcResp)
	return result
}
