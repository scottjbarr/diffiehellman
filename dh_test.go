package diffiehellman

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"testing"
)

func TestVerify(t *testing.T) {
	alice, bob := verify()

	if !reflect.DeepEqual(*alice, *bob) {
		t.Errorf("%v and %v should be equal", *alice, *bob)
	}
}

func TestDH(t *testing.T) {
	// Alice and Bob agree to use a modulus p and base g.
	//
	// choose a large prime for p (not too large though!)
	pv := "47055833459"
	// pv := "46810093859"
	p := new(big.Int)
	p.SetString(pv, 10)

	// p := newKey().X

	// choose a large prime for g. Some numbers will result in a zero
	// for A (below)
	gv := "46810093859"
	// gv := "47055833459"
	g := new(big.Int)
	g.SetString(gv, 10)
	// g := newKey().X

	// this pair can be publicly shared
	pair := NewPair(p, g)

	// k := newKey()
	// aliceNumber := k.X

	// av := "4"
	av := "1299553"
	aliceNumber := new(big.Int)
	aliceNumber.SetString(av, 10)

	log.Printf("creating A (this is too slow with large numbers)\n")
	A := powmod(pair.G, aliceNumber, pair.P)
	log.Printf("A = %v\n", A)

	bv := "1000199"
	bobNumber := new(big.Int)
	bobNumber.SetString(bv, 10)

	log.Printf("creating B (this is too slow with large numbers)\n")
	B := powmod(pair.G, bobNumber, pair.P)
	log.Printf("B = %v\n", B)

	// Alice computes s = B^a mod p
	aliceKey := powmod(B, aliceNumber, pair.P)
	fmt.Printf("alice key = %v\n", aliceKey)

	// Bob computes s = A^b mod p
	bobKey := powmod(A, bobNumber, pair.P)
	fmt.Printf("bob key = %v\n", bobKey)

	if !reflect.DeepEqual(*aliceKey, *bobKey) {
		t.Errorf("%v and %v should be equal", *aliceKey, *bobKey)
	}

}

func newKey() *ecdsa.PrivateKey {
	r := rand.Reader
	k, err := ecdsa.GenerateKey(elliptic.P256(), r)
	if err != nil {
		panic(err)
	}

	return k
}
