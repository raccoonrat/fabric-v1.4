package bccsp

import "crypto"

const (
	IBPCLA = "IBPCLA"
)

// CLVerifierOpts contains the options to verify a IBPCLA signature
type CLVerifierOpts struct {
	// KGCPublicKey is the root public-key of the KGC
	KGCPublicKey Key
	// HID is hash(ID||PA)
	HID []byte
	// PA is the vice identity
	PA []byte
	// Hash contain the hash function to be used
	Hash crypto.Hash
}

// HashFunc returns an identifier for the hash function used to produce
// the message passed to Signer.Sign, or else zero to indicate that no
// hashing was done.
func (o *CLVerifierOpts) HashFunc() crypto.Hash {
	return o.Hash
}

// CLKGCPublicKeyImportOpts contains the options for importing of a KGC public key.
type CLKGCPublicKeyImportOpts struct {
	Temporary bool
	// AttributeNames is a list of attributes to ensure the import public key has
	//AttributeNames []string
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (*CLKGCPublicKeyImportOpts) Algorithm() string {
	return IBPCLA
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (o *CLKGCPublicKeyImportOpts) Ephemeral() bool {
	return o.Temporary
}

// CLPrivateKeyImportOpts contains the options for importing of a signing private key.
type CLPrivateKeyImportOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (*CLPrivateKeyImportOpts) Algorithm() string {
	return IBPCLA
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (o *CLPrivateKeyImportOpts) Ephemeral() bool {
	return o.Temporary
}
