/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"time"

	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/token"
	tk "github.com/hyperledger/fabric/token"
)

//go:generate counterfeiter -o mock/prover.go -fake-name Prover . Prover

type Prover interface {

	// RequestImport allows the client to submit an issue request to a prover peer service;
	// the function takes as parameters tokensToIssue and the signing identity of the client;
	// it returns a response in bytes and an error message in the case the request fails.
	// The response corresponds to a serialized TokenTransaction protobuf message.
	RequestImport(tokensToIssue []*token.TokenToIssue, signingIdentity tk.SigningIdentity) ([]byte, error)

	// RequestTransfer allows the client to submit a transfer request to a prover peer service;
	// the function takes as parameters a fabtoken application credential, the identifiers of the tokens
	// to be transfererd and the shares describing how they are going to be distributed
	// among recipients; it returns a response in bytes and an error message in the case the
	// request fails
	RequestTransfer(tokenIDs [][]byte, shares []*token.RecipientTransferShare, signingIdentity tk.SigningIdentity) ([]byte, error)

	// ListTokens allows the client to submit a list request to a prover peer service;
	// it returns a list of TokenOutput and an error message in the case the request fails
	ListTokens(signingIdentity tk.SigningIdentity) ([]*token.TokenOutput, error)
}

//go:generate counterfeiter -o mock/fabric_tx_submitter.go -fake-name FabricTxSubmitter . FabricTxSubmitter

type FabricTxSubmitter interface {

	// Submit allows the client to submit a fabric token transaction.
	// It takes as input a transaction envelope and a timeout to wait for the transaction to be committed.
	// The function returns the orderer response status, committed boolean, and error in case of failure.
	// The 'waitTimeout' parameter defines the time to wait for the transaction to be committed.
	// If it is 0, the function will return immediately after receiving a response from the orderer
	// without waiting for transaction commit response.
	// If it is greater than 0, the function will wait until the transaction commit response is received or timeout, whichever is earlier.
	Submit(txEnvelope *common.Envelope, waitTimeout time.Duration) (*common.Status, bool, error)

	// CreateTxEnvelope creates a transaction envelope from the serialized TokenTransaction.
	// It returns the transaction envelope, transaction id, and error.
	CreateTxEnvelope(tokenTx []byte) (*common.Envelope, string, error)
}

// Client represents the client struct that calls Prover and TxSubmitter
type Client struct {
	Config          *ClientConfig
	SigningIdentity tk.SigningIdentity
	Prover          Prover
	TxSubmitter     FabricTxSubmitter
}

// NewClient creates a new Client from token client config
// It initializes msp crypto, gets SigningIdenity and creates a TxSubmitter
func NewClient(config ClientConfig, signingIdentity tk.SigningIdentity) (*Client, error) {
	err := ValidateClientConfig(config)
	if err != nil {
		return nil, err
	}

	prover, err := NewProverPeer(&config)
	if err != nil {
		return nil, err
	}

	txSubmitter, err := NewTxSubmitter(&config, signingIdentity)
	if err != nil {
		return nil, err
	}

	return &Client{
		Config:          &config,
		SigningIdentity: signingIdentity,
		Prover:          prover,
		TxSubmitter:     txSubmitter,
	}, nil
}

// Issue is the function that the client calls to introduce tokens into the system.
// It takes as parameter an array of token.TokenToIssue that defines what tokens are going to be issued.
// The 'waitTimeout' parameter defines the time to wait for the transaction to be committed.
// If it is 0, the function will return right after receiving a response from the orderer.
// If it is greater than 0, the function will wait until receiving the transaction event or timed out, whichever is earlier.
// This API sends the transaction to the orderer and returns the envelope, transaction id, orderer status, committed boolean, and error.
// When an error is returned, analyze the orderer status and error message to understand how to fix the problem.
// If the status is SUCCESS (200), it means that the transaction has been successfully submitted regardless of the error.
// In this case, check the transaction status to know if the transaction is committed or invalidated.
// If the transaction is invalidated, the application can fix the error and call the function again.
func (c *Client) Issue(tokensToIssue []*token.TokenToIssue, waitTimeout time.Duration) (*common.Envelope, string, *common.Status, bool, error) {
	serializedTokenTx, err := c.Prover.RequestImport(tokensToIssue, c.SigningIdentity)
	if err != nil {
		return nil, "", nil, false, err
	}

	txEnvelope, txid, err := c.TxSubmitter.CreateTxEnvelope(serializedTokenTx)
	if err != nil {
		return nil, "", nil, false, err
	}

	ordererStatus, committed, err := c.TxSubmitter.Submit(txEnvelope, waitTimeout)
	return txEnvelope, txid, ordererStatus, committed, err
}

// Transfer is the function that the client calls to transfer his tokens.
// Transfer takes as parameter an array of token.RecipientTransferShare that
// identifies who receives the tokens and describes how the tokens are distributed.
// The 'waitTimeout' parameter defines the time to wait for the transaction to be committed.
// If it is 0, the function will return right after receiving a response from the orderer.
// If it is greater than 0, the function will wait until receiving the transaction event or timed out, whichever is earlier.
// This API sends the transaction to the orderer and returns the envelope, transaction id, orderer status, committed boolean, and error.
// When an error is returned, analyze the orderer status and error message to understand how to fix the problem.
// If the status is SUCCESS (200), it means that the transaction has been successfully submitted regardless of the error.
// In this case, check the transaction status to know if the transaction is committed or invalidated.
// If the transaction is invalidated, the application can fix the error and call the function again.
func (c *Client) Transfer(tokenIDs [][]byte, shares []*token.RecipientTransferShare, waitTimeout time.Duration) (*common.Envelope, string, *common.Status, bool, error) {
	serializedTokenTx, err := c.Prover.RequestTransfer(tokenIDs, shares, c.SigningIdentity)
	if err != nil {
		return nil, "", nil, false, err
	}

	txEnvelope, txid, err := c.TxSubmitter.CreateTxEnvelope(serializedTokenTx)
	if err != nil {
		return nil, "", nil, false, err
	}

	ordererStatus, committed, err := c.TxSubmitter.Submit(txEnvelope, waitTimeout)
	return txEnvelope, txid, ordererStatus, committed, err
}

// ListTokens allows the client to submit a list request to a prover peer service;
// it returns a list of TokenOutput and an error in the case the request fails
func (c *Client) ListTokens() ([]*token.TokenOutput, error) {
	return c.Prover.ListTokens(c.SigningIdentity)
}
