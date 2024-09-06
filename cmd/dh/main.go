package main

import (
	"fmt"
	"math/big"
)

func main() {
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
	fmt.Printf("alice key : %v\n", aliceKey.Int64())
	fmt.Printf("bob key   : %v\n", bobKey.Int64())
}

func powmod(a, b, p *big.Int) *big.Int {
	return new(big.Int).Exp(a, b, p)
}
