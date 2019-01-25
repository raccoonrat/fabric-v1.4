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

func TestSigner_Sign(t *testing.T) {
	type args struct {
		k      bccsp.Key
		digest []byte
		opts   bccsp.SignerOpts
	}
	tests := []struct {
		name    string
		v       *Signer
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Signer{}
			got, err := v.Sign(tt.args.k, tt.args.digest, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signer.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Signer.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signECDSA(t *testing.T) {
	type args struct {
		k      *ecdsa.PrivateKey
		digest []byte
		opts   bccsp.SignerOpts
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := signECDSA(tt.args.k, tt.args.digest, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("signECDSA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signECDSA() = %v, want %v", got, tt.want)
			}
		})
	}
}
