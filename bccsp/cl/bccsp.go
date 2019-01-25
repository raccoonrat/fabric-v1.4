package cl

import (
	"crypto/sha256"
	"crypto/sha512"
	"reflect"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/cl/cryptocl"
	"github.com/hyperledger/fabric/bccsp/sw"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

type csp struct {
	*sw.CSP
}

type SignatureScheme interface {
	Verify(signature, msg []byte, ID string, PA []byte, KGCPub bccsp.Key) error
}

func New(keyStore bccsp.KeyStore) (*csp, error) {
	base, err := sw.New(keyStore)
	if err != nil {
		return nil, errors.Wrap(err, "failed instantiating base bccsp")
	}

	csp := &csp{CSP: base}

	// signers
	base.AddWrapper(reflect.TypeOf(cryptocl.NewSecretKey()), &cryptocl.Signer{})

	// verifiers
	base.AddWrapper(reflect.TypeOf(&bccsp.CLVerifierOpts{}), &cryptocl.Verifier{})

	//importers
	base.AddWrapper(reflect.TypeOf(&bccsp.CLKGCPublicKeyImportOpts{}), &cryptocl.KGCKeyImporter{})
	base.AddWrapper(reflect.TypeOf(&bccsp.CLPrivateKeyImportOpts{}), &cryptocl.SignerKeyImporter{})

	//hashers
	base.AddWrapper(reflect.TypeOf(&bccsp.SHA256Opts{}), &cryptocl.Hasher{DoHash: sha256.New})
	base.AddWrapper(reflect.TypeOf(&bccsp.SHA384Opts{}), &cryptocl.Hasher{DoHash: sha512.New384})
	base.AddWrapper(reflect.TypeOf(&bccsp.SHA3_256Opts{}), &cryptocl.Hasher{DoHash: sha3.New256})
	base.AddWrapper(reflect.TypeOf(&bccsp.SHA3_384Opts{}), &cryptocl.Hasher{DoHash: sha3.New384})

	return csp, nil
}

// Sign signs digest using key k.
// The opts argument should be appropriate for the primitive used.
//
// Note that when a signature of a hash of a larger message is needed,
// the caller is responsible for hashing the larger message and passing
// the hash (as digest).
// Notice that this is overriding the Sign methods of the sw impl. to avoid the digest check.
func (csp *csp) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) (signature []byte, err error) {
	// Validate arguments
	if k == nil {
		return nil, errors.New("Invalid Key. It must not be nil.")
	}
	// Do not check for digest

	keyType := reflect.TypeOf(k)
	signer, found := csp.Signers[keyType]
	if !found {
		return nil, errors.Errorf("Unsupported 'SignKey' provided [%s]", keyType)
	}

	signature, err = signer.Sign(k, digest, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed signing with opts [%v]", opts)
	}

	return
}

// Verify verifies signature against key k and digest
// Notice that this is overriding the Verify methods of the sw impl.
func (csp *csp) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (valid bool, err error) {
	// Validate arguments
	if k != nil {
		return false, errors.New("Invalid Key. It must be nil.")
	}
	if len(signature) == 0 {
		return false, errors.New("Invalid signature. Cannot be empty.")
	}

	verifier, found := csp.Verifiers[reflect.TypeOf(opts)]
	if !found {
		return false, errors.Errorf("Unsupported CLVerifierOpts provided [%v]", opts)
	}

	valid, err = verifier.Verify(k, signature, digest, opts)
	if err != nil {
		return false, errors.Wrapf(err, "Failed verifing with opts [%v]", opts)
	}

	return
}
