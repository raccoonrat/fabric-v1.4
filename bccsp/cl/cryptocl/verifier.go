/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package cryptocl

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/pkg/errors"
)

var (
	logger = flogging.MustGetLogger("bccsp_cl")
)

type Verifier struct {
}

func (v *Verifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	if k != nil {
		return false, errors.New("invalid key, expected nil")
	}

	verifierOpts, ok := opts.(*bccsp.CLVerifierOpts)
	if !ok {
		return false, errors.New("invalid options, expected *CLVerifierOpts")
	}

	if len(signature) == 0 {
		return false, errors.New("invalid signature, it must not be empty")
	}

	//Recover Pub
	pubkey, err := RecoverPub(verifierOpts)
	if err != nil {
		return false, errors.New("RecoverPub Fail")
	}
	//Verify sig
	valid, err := verifyECDSA(pubkey, signature, digest, opts)
	if err != nil {
		return false, err
	}

	return valid, nil
}

func RecoverPub(opts *bccsp.CLVerifierOpts) (*ecdsa.PublicKey, error) {

	rootPub := opts.KGCPublicKey.(*rootPublicKey).pubKey
	if rootPub == nil {
		return nil, errors.New("can not load root public key")
	}

	HID := opts.HID
	if HID == nil {
		return nil, errors.New("can not load HID from opts")
	}

	PA := opts.PA
	if PA == nil {
		return nil, errors.New("can not load PA from opts")
	}

	c := rootPub.Curve
	if c == nil {
		return nil, errors.New("can not load curve parameters from root public key")
	}

	N := c.Params().N

	PAx := new(big.Int).SetBytes(PA[0:32])
	PAy := new(big.Int).SetBytes(PA[32:64])

	//Pub = e0*PA + e1*rootPub
	x1, y1 := c.ScalarMult(PAx, PAy, HID[0:15])
	x2, y2 := c.ScalarMult(rootPub.X, rootPub.Y, HID[16:31])
	x, y := c.Add(x1, y1, x2, y2)
	x.Mod(x, N)
	y.Mod(y, N)

	return &ecdsa.PublicKey{Curve: c, X: x, Y: y}, nil
}

func verifyECDSA(k *ecdsa.PublicKey, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	r, s, err := utils.UnmarshalECDSASignature(signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}

	lowS, err := utils.IsLowS(k, s)
	if err != nil {
		return false, err
	}

	if !lowS {
		return false, fmt.Errorf("Invalid S. Must be smaller than half the order [%s][%s].", s, utils.GetCurveHalfOrdersAt(k.Curve))
	}

	return ecdsa.Verify(k, digest, r, s), nil
}
