package kmscc

import "github.com/hyperledger/fabric/core/scc/kmscc/rpc"

type CreateCommitteeRequest struct {
	N      int    `json:"n"`      //委员会最大成员数
	K      int    `json:"k"`      //恢复密钥依赖的最少成员数
	Policy []byte `json:"policy"` //委员会的成员条件
}

func CreateCommittee(ctx *rpc.JsonRpcContext,
	req *CreateCommitteeRequest, resp *Response) *rpc.JsonRpcError {
	return nil
}
