/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package kgc

import (
	"crypto/ecdsa"
	"math/big"
	"reflect"
	"testing"
)

func TestNewKGC(t *testing.T) {
	type args struct {
		baseDir string
		org     string
		name    string
	}
	tests := []struct {
		name    string
		args    args
		want    *KGC
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKGC(tt.args.baseDir, tt.args.org, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKGC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKGC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKGC_KGCGenPartialKey(t *testing.T) {
	type fields struct {
		Name         string
		MasterKey    *ecdsa.PrivateKey
		MasterPub    *ecdsa.PublicKey
		RawPub       []byte
		Organization string
	}
	type args struct {
		baseDir string
		ID      string
		XA      *ecdsa.PublicKey
		s       *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PartialKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kgc := &KGC{
				Name:         tt.fields.Name,
				MasterKey:    tt.fields.MasterKey,
				MasterPub:    tt.fields.MasterPub,
				RawPub:       tt.fields.RawPub,
				Organization: tt.fields.Organization,
			}
			got, err := kgc.KGCGenPartialKey(tt.args.baseDir, tt.args.ID, tt.args.XA, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("KGC.KGCGenPartialKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KGC.KGCGenPartialKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKGCGenPartialKeyInternal(t *testing.T) {
	type args struct {
		ID string
		XA *ecdsa.PublicKey
		s  *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		want1   *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := KGCGenPartialKeyInternal(tt.args.ID, tt.args.XA, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("KGCGenPartialKeyInternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KGCGenPartialKeyInternal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("KGCGenPartialKeyInternal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLoadKGCPublicKey(t *testing.T) {
	type args struct {
		certPath string
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
			got, got1, err := LoadKGCPublicKey(tt.args.certPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadKGCPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadKGCPublicKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadKGCPublicKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
