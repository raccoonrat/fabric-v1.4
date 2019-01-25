/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package cryptocl

import (
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
)

func TestKGCKeyImporter_KeyImport(t *testing.T) {
	type args struct {
		raw  interface{}
		opts bccsp.KeyImportOpts
	}
	tests := []struct {
		name    string
		v       *KGCKeyImporter
		args    args
		wantK   bccsp.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &KGCKeyImporter{}
			gotK, err := v.KeyImport(tt.args.raw, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("KGCKeyImporter.KeyImport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Errorf("KGCKeyImporter.KeyImport() = %v, want %v", gotK, tt.wantK)
			}
		})
	}
}

func TestSignerKeyImporter_KeyImport(t *testing.T) {
	type args struct {
		raw  interface{}
		opts bccsp.KeyImportOpts
	}
	tests := []struct {
		name    string
		v       *SignerKeyImporter
		args    args
		wantK   bccsp.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &SignerKeyImporter{}
			gotK, err := v.KeyImport(tt.args.raw, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignerKeyImporter.KeyImport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Errorf("SignerKeyImporter.KeyImport() = %v, want %v", gotK, tt.wantK)
			}
		})
	}
}
