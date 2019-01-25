//SPDX-License-Identifier: Apache-2.0
package cryptocl

import (
	"hash"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
)

func TestHasher_Hash(t *testing.T) {
	type fields struct {
		DoHash func() hash.Hash
	}
	type args struct {
		msg  []byte
		opts bccsp.HashOpts
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
			c := &Hasher{
				DoHash: tt.fields.DoHash,
			}
			got, err := c.Hash(tt.args.msg, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hasher.Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hasher.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasher_GetHash(t *testing.T) {
	type fields struct {
		DoHash func() hash.Hash
	}
	type args struct {
		opts bccsp.HashOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    hash.Hash
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Hasher{
				DoHash: tt.fields.DoHash,
			}
			got, err := c.GetHash(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hasher.GetHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hasher.GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
