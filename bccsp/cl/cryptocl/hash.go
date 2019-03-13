package cryptocl

import (
	"hash"

	"github.com/hyperledger/fabric/bccsp"
)

type Hasher struct {
	DoHash func() hash.Hash
}

func (c *Hasher) Hash(msg []byte, opts bccsp.HashOpts) ([]byte, error) {
	h := c.DoHash()
	h.Write(msg)
	return h.Sum(nil), nil
}

func (c *Hasher) GetHash(opts bccsp.HashOpts) (hash.Hash, error) {
	return c.DoHash(), nil
}
