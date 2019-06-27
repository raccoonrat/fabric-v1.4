/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package cryptocl

import (
	"crypto/ecdsa"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
)

func TestVerifier_Verify(t *testing.T) {
	type args struct {
		k         bccsp.Key
		signature []byte
		digest    []byte
		opts      bccsp.SignerOpts
	}
	tests := []struct {
		name    string
		v       *Verifier
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Verifier{}
			got, err := v.Verify(tt.args.k, tt.args.signature, tt.args.digest, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verifier.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verifier.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecoverPub(t *testing.T) {
	type args struct {
		opts *bccsp.CLVerifierOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PublicKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecoverPub(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecoverPub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecoverPub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyECDSA(t *testing.T) {
	type args struct {
		k         *ecdsa.PublicKey
		signature []byte
		digest    []byte
		opts      bccsp.SignerOpts
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := verifyECDSA(tt.args.k, tt.args.signature, tt.args.digest, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyECDSA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("verifyECDSA() = %v, want %v", got, tt.want)
			}
		})
	}
}
