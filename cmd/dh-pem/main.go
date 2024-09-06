package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// From https://datatracker.ietf.org/doc/html/rfc3526#section-3
	//
	// 3.  2048-bit MODP Group
	//
	// This group is assigned id 14.
	//
	// This prime is: 2^2048 - 2^1984 - 1 + 2^64 * { [2^1918 pi] + 124476 }
	//
	// Its hexadecimal value is:
	//
	//    FFFFFFFF FFFFFFFF C90FDAA2 2168C234 C4C6628B 80DC1CD1
	//    29024E08 8A67CC74 020BBEA6 3B139B22 514A0879 8E3404DD
	//    EF9519B3 CD3A431B 302B0A6D F25F1437 4FE1356D 6D51C245
	//    E485B576 625E7EC6 F44C42E9 A637ED6B 0BFF5CB6 F406B7ED
	//    EE386BFB 5A899FA5 AE9F2411 7C4B1FE6 49286651 ECE45B3D
	//    C2007CB8 A163BF05 98DA4836 1C55D39A 69163FA8 FD24CF5F
	//    83655D23 DCA3AD96 1C62F356 208552BB 9ED52907 7096966D
	//    670C354E 4ABC9804 F1746C08 CA18217C 32905E46 2E36CE3B
	//    E39E772C 180E8603 9B2783A2 EC07A28F B5C55DF0 6F4C52C9
	//    DE2BCBF6 95581718 3995497C EA956AE5 15D22618 98FA0510
	//    15728E5A 8AACAA68 FFFFFFFF FFFFFFFF

	base := 36

	// Alice and Bob agree to use group 14 from RFC-3526
	p, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF", 16)
	g := new(big.Int).SetInt64(2)

	fmt.Printf("p = 0x%s\n", p.Text(base))
	fmt.Printf("g = %v\n", g.Text(base))

	// Alice chooses a secret integer, then sends Bob A = g^a mod p
	a, err := newRandomBigInt()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Alice choses secret int\n  %s\n", a.Text(base))

	A := powmod(g, a, p)

	fmt.Printf("Alice sends Bob the result of g^a mod p\n  %s\n", A.Text(base))

	// Bob chooses a secret integer, then sends Alice B = g^b mod p
	b, err := newRandomBigInt()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bob choses secret int\n  %s\n", a.Text(base))

	B := powmod(g, b, p)

	fmt.Printf("Bob sends Alice the result of g^b mod p\n  %s\n", B.Text(base))

	// Alice computes s = B^a mod p
	aliceKey := powmod(B, a, p)

	// Bob computes s = A^b mod p
	bobKey := powmod(A, b, p)

	// Alice and Bob now share a secret
	// fmt.Printf("Alice computes key\n  %v\n", aliceKey.Int64())
	// fmt.Printf("Bob computes key\n  %v\n", bobKey.Int64())

	fmt.Printf("Alice computes key\n  %v\n", aliceKey.Text(base))
	fmt.Printf("Bob computes key\n  %v\n", bobKey.Text(base))

	if aliceKey.Cmp(bobKey) != 0 {
		panic("keys must be equal")
	}
}

func powmod(a, b, p *big.Int) *big.Int {
	// a^b
	// fmt.Printf("exponent 0x%s ^ 0x%s\n", a.Text(base), b.Text(base))
	// e := new(big.Int)
	// e.Exp(a, b, nil)

	// // e mod p
	// fmt.Printf("mod\n")
	// n := new(big.Int)
	// return n.Mod(e, p)

	return new(big.Int).Exp(a, b, p)
}

func newRandomBigInt() (*big.Int, error) {
	// Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(1024), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	return rand.Int(rand.Reader, max)
}
