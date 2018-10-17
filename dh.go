package diffiehellman

import (
	"math/big"
)

func verify() (*big.Int, *big.Int) {
	// Alice and Bob agree to use a modulus p = 23 and base g = 5 (which is
	// a primitive root modulo 23).
	p := new(big.Int)
	p.SetString("23", 10)

	g := new(big.Int)
	g.SetString("5", 10)

	// Alice chooses a secret integer a = 4, then sends Bob A = g^a mod p
	//     A = 54 mod 23 = 4
	a := new(big.Int)
	a.SetString("4", 10)

	A := powmod(g, a, p)

	// Bob chooses a secret integer b = 3, then sends Alice B = g^b mod p
	//     B = 53 mod 23 = 10
	b := new(big.Int)
	b.SetString("3", 10)

	B := powmod(g, b, p)

	// Alice computes s = B^a mod p
	//     s = 104 mod 23 = 18
	aliceKey := powmod(B, a, p)

	// Bob computes s = A^b mod p
	//     s = 43 mod 23 = 18
	bobKey := powmod(A, b, p)

	// Alice and Bob now share a secret (the number 18).
	return aliceKey, bobKey
}

func verifyValues(p, g, a, b *big.Int) (*big.Int, *big.Int) {
	A := powmod(g, a, p)

	B := powmod(g, b, p)

	return powmod(B, a, p), powmod(A, b, p)
}

func powmod(a, b, p *big.Int) *big.Int {
	// a^b
	e := new(big.Int)
	e.Exp(a, b, nil)

	// e mod p
	B := new(big.Int)
	B.Mod(e, p)

	return B
}

type Pair struct {
	P *big.Int
	G *big.Int
}

func NewPair(p, g *big.Int) Pair {
	return Pair{
		P: p,
		G: g,
	}
}
