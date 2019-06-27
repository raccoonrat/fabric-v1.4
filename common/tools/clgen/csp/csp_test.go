/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package csp

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
)

func TestLoadPrivateKey(t *testing.T) {
	type args struct {
		keystorePath string
	}
	tests := []struct {
		name    string
		args    args
		want    bccsp.Key
		want1   crypto.Signer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := LoadPrivateKey(tt.args.keystorePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadPrivateKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadPrivateKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLoadKGCMasterKey(t *testing.T) {
	type args struct {
		keystorePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PrivateKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadKGCMasterKey(tt.args.keystorePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadKGCMasterKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadKGCMasterKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratePrivateKey(t *testing.T) {
	type args struct {
		keystorePath string
	}
	tests := []struct {
		name    string
		args    args
		want    bccsp.Key
		want1   crypto.Signer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GeneratePrivateKey(tt.args.keystorePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePrivateKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GeneratePrivateKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetECPublicKey(t *testing.T) {
	type args struct {
		priv bccsp.Key
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
			got, err := GetECPublicKey(tt.args.priv)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetECPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetECPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKGCGenerateMasterKey(t *testing.T) {
	type args struct {
		keystorePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PrivateKey
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := KGCGenerateMasterKey(tt.args.keystorePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("KGCGenerateMasterKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KGCGenerateMasterKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("KGCGenerateMasterKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBccspKey2ecdsaKey(t *testing.T) {
	type args struct {
		bkey bccsp.Key
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PrivateKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BccspKey2ecdsaKey(tt.args.bkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("BccspKey2ecdsaKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BccspKey2ecdsaKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKGCGetECPublicKey(t *testing.T) {
	type args struct {
		priv         bccsp.Key
		name         string
		keystorePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PublicKey
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := KGCGetECPublicKey(tt.args.priv, tt.args.name, tt.args.keystorePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("KGCGetECPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KGCGetECPublicKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("KGCGetECPublicKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenFinalKeyPair(t *testing.T) {
	type args struct {
		ID                string
		OU                string
		Role              string
		ClientPrivateKey  *ecdsa.PrivateKey
		PartialPublicKey  []byte
		PartialPrivateKey []byte
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
			got, err := GenFinalKeyPair(tt.args.ID, tt.args.OU, tt.args.Role, tt.args.ClientPrivateKey, tt.args.PartialPublicKey, tt.args.PartialPrivateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenFinalKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenFinalKeyPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateKeyToPEM(t *testing.T) {
	type args struct {
		k *ecdsa.PrivateKey
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
			got, err := PrivateKeyToPEM(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToPEM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrivateKeyToPEM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oidFromNamedCurve(t *testing.T) {
	type args struct {
		curve elliptic.Curve
	}
	tests := []struct {
		name  string
		args  args
		want  asn1.ObjectIdentifier
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := oidFromNamedCurve(tt.args.curve)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oidFromNamedCurve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("oidFromNamedCurve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenSerial(t *testing.T) {
	type args struct {
		za []byte
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
			if got := GenSerial(tt.args.za); got != tt.want {
				t.Errorf("GenSerial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateKey(t *testing.T) {
	type args struct {
		dA   []byte
		P1   ecdsa.PublicKey
		Pa   []byte
		ID   string
		OU   string
		Role string
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
			if err := ValidateKey(tt.args.dA, tt.args.P1, tt.args.Pa, tt.args.ID, tt.args.OU, tt.args.Role); (err != nil) != tt.wantErr {
				t.Errorf("ValidateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrivateKeyToDER(t *testing.T) {
	type args struct {
		d []byte
		c elliptic.Curve
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
			got, err := PrivateKeyToDER(tt.args.d, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToDER() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrivateKeyToDER() = %v, want %v", got, tt.want)
			}
		})
	}
}
