package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"time"

	"github.com/monnand/dhkx"
)

func main() {
	randomSize := int64(16)

	log.Printf("starting")

	// Alice and Bob agree to use a modulus p = 23 and base g = 5 (which is
	// a primitive root modulo 23).
	// p := new(big.Int)
	// p.SetString("23", 10)

	// g := new(big.Int)
	// g.SetString("5", 10)

	group, err := dhkx.GetGroup(14)
	if err != nil {
		panic(err)
	}

	p := group.P()
	g := group.G()

	// Alice chooses a secret integer a = 4, then sends Bob A = g^a mod p
	//     A = 54 mod 23 = 4
	a, err := newRandomBigInt(randomSize)
	if err != nil {
		panic(err)
	}

	log.Printf("Alice choses secret int 0x%s", a.Text(16))

	A := powmod(g, a, p)

	// Bob chooses a secret integer b = 3, then sends Alice B = g^b mod p
	//     B = 53 mod 23 = 10
	b, err := newRandomBigInt(32)
	if err != nil {
		panic(err)
	}

	log.Printf("Bob choses secret int 0x%s", b.Text(16))

	B := powmod(g, b, p)

	// Alice computes s = B^a mod p
	//     s = 104 mod 23 = 18
	aliceKey := powmod(B, a, p)

	// Bob computes s = A^b mod p
	//     s = 43 mod 23 = 18
	bobKey := powmod(A, b, p)

	// Alice and Bob now share a secret (the number 18).
	log.Printf("alice key : %v", aliceKey.Int64())
	log.Printf("bob key   : %v", bobKey.Int64())
}

func powmod(a, b, p *big.Int) *big.Int {
	defer elapsed(time.Now(), "powmod")

	// a^b
	e := exp(a, b)

	// e mod p
	return mod(e, p)
}

func exp(a, b *big.Int) *big.Int {
	defer elapsed(time.Now(), "exp")
	// fmt.Printf("exp 0x%s ^ 0x%s\n", a.Text(16), b.Text(16))

	e := new(big.Int)
	e.Exp(a, b, nil)

	return e
}

func mod(e, p *big.Int) *big.Int {
	defer elapsed(time.Now(), "mod")
	// fmt.Printf("mod 0x%s ^ 0x%s\n", e.Text(16), p.Text(16))

	n := new(big.Int)
	n.Mod(e, p)

	return n
}

func newRandomBigInt(size int64) (*big.Int, error) {
	// Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(size), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	return rand.Int(rand.Reader, max)
}

func elapsed(start time.Time, name string) {
	go func() {
		// get elapsed time in milliseconds
		elapsed := float64(time.Since(start).Nanoseconds()) / float64(1000000)

		log.Printf("%s elapsed=%v", name, elapsed)
	}()
}
