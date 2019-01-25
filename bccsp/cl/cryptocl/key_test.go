package cryptocl

import (
	"crypto/ecdsa"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/bccsp"
)

func Test_rootPublicKey_Bytes(t *testing.T) {
	type fields struct {
		pubKey *ecdsa.PublicKey
	}
	tests := []struct {
		name    string
		fields  fields
		wantRaw []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := rootPublicKey{
				pubKey: tt.fields.pubKey,
			}
			gotRaw, err := k.Bytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("rootPublicKey.Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRaw, tt.wantRaw) {
				t.Errorf("rootPublicKey.Bytes() = %v, want %v", gotRaw, tt.wantRaw)
			}
		})
	}
}

func Test_rootPublicKey_SKI(t *testing.T) {
	type fields struct {
		pubKey *ecdsa.PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := rootPublicKey{
				pubKey: tt.fields.pubKey,
			}
			if got := k.SKI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootPublicKey.SKI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootPublicKey_Symmetric(t *testing.T) {
	type fields struct {
		pubKey *ecdsa.PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := rootPublicKey{
				pubKey: tt.fields.pubKey,
			}
			if got := k.Symmetric(); got != tt.want {
				t.Errorf("rootPublicKey.Symmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootPublicKey_Private(t *testing.T) {
	type fields struct {
		pubKey *ecdsa.PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := rootPublicKey{
				pubKey: tt.fields.pubKey,
			}
			if got := k.Private(); got != tt.want {
				t.Errorf("rootPublicKey.Private() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootPublicKey_PublicKey(t *testing.T) {
	type fields struct {
		pubKey *ecdsa.PublicKey
	}
	tests := []struct {
		name    string
		fields  fields
		want    bccsp.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := rootPublicKey{
				pubKey: tt.fields.pubKey,
			}
			got, err := k.PublicKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("rootPublicKey.PublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootPublicKey.PublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signKey_Bytes(t *testing.T) {
	type fields struct {
		privKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		fields  fields
		wantRaw []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := signKey{
				privKey: tt.fields.privKey,
			}
			gotRaw, err := k.Bytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("signKey.Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRaw, tt.wantRaw) {
				t.Errorf("signKey.Bytes() = %v, want %v", gotRaw, tt.wantRaw)
			}
		})
	}
}

func Test_signKey_SKI(t *testing.T) {
	type fields struct {
		privKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := signKey{
				privKey: tt.fields.privKey,
			}
			if got := k.SKI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signKey.SKI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signKey_Symmetric(t *testing.T) {
	type fields struct {
		privKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := signKey{
				privKey: tt.fields.privKey,
			}
			if got := k.Symmetric(); got != tt.want {
				t.Errorf("signKey.Symmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signKey_Private(t *testing.T) {
	type fields struct {
		privKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := signKey{
				privKey: tt.fields.privKey,
			}
			if got := k.Private(); got != tt.want {
				t.Errorf("signKey.Private() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signKey_PublicKey(t *testing.T) {
	type fields struct {
		privKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		fields  fields
		want    bccsp.Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := signKey{
				privKey: tt.fields.privKey,
			}
			got, err := k.PublicKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("signKey.PublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signKey.PublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSecretKey(t *testing.T) {
	tests := []struct {
		name string
		want *signKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSecretKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSecretKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
