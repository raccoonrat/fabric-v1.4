/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package signer

import (
	"crypto"
	"io"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
)

func TestNew(t *testing.T) {
	type args struct {
		csp bccsp.BCCSP
		key bccsp.Key
	}
	tests := []struct {
		name    string
		args    args
		want    crypto.Signer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.csp, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clCryptoSigner_Public(t *testing.T) {
	type fields struct {
		csp bccsp.BCCSP
		key bccsp.Key
		pk  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   crypto.PublicKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &clCryptoSigner{
				csp: tt.fields.csp,
				key: tt.fields.key,
				pk:  tt.fields.pk,
			}
			if got := s.Public(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clCryptoSigner.Public() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clCryptoSigner_Sign(t *testing.T) {
	type fields struct {
		csp bccsp.BCCSP
		key bccsp.Key
		pk  interface{}
	}
	type args struct {
		rand   io.Reader
		digest []byte
		opts   crypto.SignerOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &clCryptoSigner{
				csp: tt.fields.csp,
				key: tt.fields.key,
				pk:  tt.fields.pk,
			}
			got, err := s.Sign(tt.args.rand, tt.args.digest, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("clCryptoSigner.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clCryptoSigner.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
