package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/monnand/dhkx"
)

type peer struct {
	priv  *dhkx.DHKey
	group *dhkx.DHGroup
	pub   *dhkx.DHKey
}

func newPeer(g *dhkx.DHGroup) *peer {
	ret := new(peer)
	// priv, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	ret.priv, _ = g.GeneratePrivateKey(rand.Reader)
	// ret.priv, _ = &dhkx.DHKey{
	// 	PrivateKey: priv,
	// 	Group:      g,
	// 	// group:      g,
	// }
	ret.group = g

	return ret
}

func (p *peer) getPubKey() []byte {
	return p.priv.Bytes()
}

func (p *peer) recvPeerPubKey(pub []byte) {
	pubKey := dhkx.NewPublicKey(pub)
	p.pub = pubKey
}

func (p *peer) getKey() []byte {
	k, err := p.group.ComputeKey(p.pub, p.priv)
	if err != nil {
		return nil
	}

	return k.Bytes()
}

func exchangeKey(p1, p2 *peer) error {
	pub1 := p1.getPubKey()
	pub2 := p2.getPubKey()

	p1.recvPeerPubKey(pub2)
	p2.recvPeerPubKey(pub1)

	key1 := p1.getKey()
	if key1 == nil {
		return fmt.Errorf("p1 has nil key")
	}

	key2 := p2.getKey()

	if key2 == nil {
		return fmt.Errorf("p2 has nil key")
	}

	for i, k := range key1 {
		if key2[i] != k {
			return fmt.Errorf("%vth byte does not same")
		}
	}

	return nil
}

func main() {
	group, err := dhkx.GetGroup(14)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("group : P = 0x%s\n", group.P().Text(16))
	// fmt.Printf("group : G = 0x%s\n", group.G().Text(16))

	p1 := newPeer(group)
	p2 := newPeer(group)

	if err := exchangeKey(p1, p2); err != nil {
		panic(err)
	}

	// fmt.Printf("p1 private key = %#+v\n", p1.priv)
	fmt.Printf("p1 private key = %s\n", p1.priv.String())
	fmt.Printf("p1 pub = %+v\n", *p1.pub)

	fmt.Printf("p2 private key = %s\n", p2.priv.String())

	// r := bytes.NewReader(b)
	// k1, err := rsa.GenerateKey(r, 2048)
	// if err != nil {
	// 	panic(err)
	// }

	message := "what the actual ...?"

	// k1 := rsa.PrivateKey{
	// 	D: new(big.Int).SetBytes(p1.priv.Bytes()),
	// 	PublicKey: rsa.PublicKey{
	// 		N: p1.group.P(),
	// 		E: 2,
	// 	},
	// }

	// k2 := rsa.PrivateKey{
	// 	D: new(big.Int).SetBytes(p2.priv.Bytes()),
	// 	PublicKey: rsa.PublicKey{
	// 		N: p2.group.P(),
	// 		E: 2,
	// 	},
	// }

	k2 := ecdsa.PrivateKey{
		D: new(big.Int).SetBytes(p2.priv.Bytes()),
		PublicKey: ecdsa.PublicKey{
			X: p2.getPubKey().X,
			Y: p2.getPubKey().Y,
		},
	}
	// fmt.Printf("k1 = %+v\n", k1)

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &k2.PublicKey, []byte(message), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ciphertext = %x\n", ciphertext)

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, &k2, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("plaintext = %x\n", plaintext)
}
