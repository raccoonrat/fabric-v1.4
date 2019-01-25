package cl

import (
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/sw"
)

func TestNew(t *testing.T) {
	type args struct {
		keyStore bccsp.KeyStore
	}
	tests := []struct {
		name    string
		args    args
		want    *csp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.keyStore)
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

func Test_csp_Verify(t *testing.T) {
	type fields struct {
		CSP *sw.CSP
	}
	type args struct {
		k         bccsp.Key
		signature []byte
		digest    []byte
		opts      bccsp.SignerOpts
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValid bool
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csp := &csp{
				CSP: tt.fields.CSP,
			}
			gotValid, err := csp.Verify(tt.args.k, tt.args.signature, tt.args.digest, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("csp.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValid != tt.wantValid {
				t.Errorf("csp.Verify() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
