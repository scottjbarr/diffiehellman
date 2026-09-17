package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//
	// See https://gist.github.com/ekmadsen/99d35898c93da7532e8aafeadd79fce6
	//

	// setup
	fmt.Printf("Testing using big, positive integers.\n")

	g := newRandomBigInt()
	fmt.Printf("g  = %v .  This is the encryption generator.\n", g.Text(32))

	// Large messages are encrypted in multiple blocks.
	// n := big.NewInt(5)
	n := newRandomBigInt()
	fmt.Printf("n  = %v .  This is the block size of the generator.\n", n.Text(32))

	// Send public key from A to B.
	a := newRandomBigInt()
	fmt.Printf("a  = %v .  This is a random integer chosen by Principal A.\n", a.Text(32))

	m1 := modpow(g, a, n)
	fmt.Printf("m1 = %v .  This is a public key transmitted from Principal A to Principal B.\n", m1.Text(32))

	// Send public key from B to A.
	b := newRandomBigInt()
	fmt.Printf("b  = %v .  This is a random integer chosen by Principal B.\n", b.Text(32))

	m2 := modpow(g, b, n)
	fmt.Printf("m2 = %v .  This is a public key transmitted from Principal B to Principal A.\n", m2.Text(32))

	// Compute shared keys.
	ak := modpow(m2, a, n)
	fmt.Printf("ak = %v . Shared key Principal A uses to encrypt and decrypt messages sent to /received from Principal B.\n", ak.Text(32))

	bk := modpow(m1, b, n)
	fmt.Printf("bk = %v .  This is the shared key Principal B uses to encrypt and decrypt messages sent to/received from Principal A.\n", bk.Text(32))

	// if ak == bk {
	if ak.Cmp(bk) == 0 {
		fmt.Printf("Shared keys (for use by Principals A and B using generator g) match.\n")
	} else {
		fmt.Printf("Shared keys (for generator g) do not match.\n")
	}
}

func newRandomBigInt() *big.Int {
	// Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(1024), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return i
}

func modpow(a, b, p *big.Int) *big.Int {
	return new(big.Int).Exp(a, b, p)
}
