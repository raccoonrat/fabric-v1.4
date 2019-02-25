package cl

import (
	"bytes"
	"math/big"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/sw"
	"github.com/pkg/errors"
)

type csp struct {
	*sw.CSP
}

type PA struct {
	raw []byte
}

func New(keyStore bccsp.KeyStore) (*csp, error) {
	base, err := sw.New(keyStore)
	if err != nil {
		return nil, errors.Wrap(err, "failed instantiating base bccsp")
	}

	csp := &csp{CSP: base}
	return csp, nil
}

func RecoverPub(rootPub bccsp.Key, PA []byte, ID string, hashOpt bccsp.HashOpts) (bccsp.Key, error) {

	//
	var buffer bytes.Buffer
	c := rootPub.pubKey.Curve
	N := c.Params().N

	buffer.Write([]byte(ID))
	buffer.Write(PA)
	e := csp.Hash(buffer.Bytes(), hashOpt)
	e0 := new(big.Int).SetBytes(e[0:15])
	e1 := new(big.Int).SetBytes(e[16:32])
	//Pub = e0*PA + e1*rootPub
	PAInt := new(big.Int).SetBytes(PA)
	e0.Mul(PAInt, e0)
	e0.Mod(e0, N)

	PInt := rootPub.(*ecdsaPublicKey).pubKey.X
	e1.Mul(PInt, e1)
	e1.Mod(e1, N)

	pubInt := new(big.Int).Add(e0, e1)
	pubInt.Mod(pubInt, N)

	pubKey, err := csp.KeyImport(pubInt.Bytes(), &bccsp.ECDSAPublicKeyImportOpts{Temporary: true})
	if err != nil {
		return nil, errors.WithMessage(err, "RecoverPubs error: Failed to import EC public key")
	}
	return pubKey, nil
}
