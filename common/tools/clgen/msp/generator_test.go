/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package msp

import (
	"crypto/x509"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/common/tools/clgen/ca"
	"github.com/hyperledger/fabric/common/tools/clgen/kgc"
)

func TestGenerateLocalMSP(t *testing.T) {
	type args struct {
		baseDir  string
		name     string
		sans     []string
		signKGC  *kgc.KGC
		tlsCA    *ca.CA
		nodeType int
		nodeOUs  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateLocalMSP(tt.args.baseDir, tt.args.name, tt.args.sans, tt.args.signKGC, tt.args.tlsCA, tt.args.nodeType, tt.args.nodeOUs); (err != nil) != tt.wantErr {
				t.Errorf("GenerateLocalMSP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateVerifyingMSP(t *testing.T) {
	type args struct {
		baseDir string
		signKGC *kgc.KGC
		tlsCA   *ca.CA
		nodeOUs bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateVerifyingMSP(tt.args.baseDir, tt.args.signKGC, tt.args.tlsCA, tt.args.nodeOUs); (err != nil) != tt.wantErr {
				t.Errorf("GenerateVerifyingMSP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createFolderStructure(t *testing.T) {
	type args struct {
		rootDir string
		local   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createFolderStructure(tt.args.rootDir, tt.args.local); (err != nil) != tt.wantErr {
				t.Errorf("createFolderStructure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_x509Filename(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := x509Filename(tt.args.name); got != tt.want {
				t.Errorf("x509Filename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_x509Export(t *testing.T) {
	type args struct {
		path string
		cert *x509.Certificate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := x509Export(tt.args.path, tt.args.cert); (err != nil) != tt.wantErr {
				t.Errorf("x509Export() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMasterPubFilename(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MasterPubFilename(tt.args.name); got != tt.want {
				t.Errorf("MasterPubFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasterPubExport(t *testing.T) {
	type args struct {
		path string
		raw  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MasterPubExport(tt.args.path, tt.args.raw); (err != nil) != tt.wantErr {
				t.Errorf("MasterPubExport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_keyExport(t *testing.T) {
	type args struct {
		keystore string
		output   string
		key      bccsp.Key
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := keyExport(tt.args.keystore, tt.args.output, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("keyExport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pemExport(t *testing.T) {
	type args struct {
		path    string
		pemType string
		bytes   []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pemExport(tt.args.path, tt.args.pemType, tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("pemExport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exportConfig(t *testing.T) {
	type args struct {
		mspDir  string
		kgcFile string
		enable  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exportConfig(tt.args.mspDir, tt.args.kgcFile, tt.args.enable); (err != nil) != tt.wantErr {
				t.Errorf("exportConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exportConfigID(t *testing.T) {
	type args struct {
		mspDir  string
		kgcFile string
		enable  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exportConfigID(tt.args.mspDir, tt.args.kgcFile, tt.args.enable); (err != nil) != tt.wantErr {
				t.Errorf("exportConfigID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
