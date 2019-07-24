package diffiehellman

import (
	"bytes"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"math/big"
)

// DecodePEM decodes the content PEM file that contains Diffie-Hellman
// parameters.
func DecodePEM(b []byte) (*big.Int, *big.Int) {
	block, _ := pem.Decode(b)

	if block == nil || block.Type != "DH PARAMETERS" {
		panic("failed to decode PEM block containing public key")
	}

	buf := bytes.NewBuffer(block.Bytes)

	// fmt.Printf("%x\n", b)
	// printHex("len(b)", len(b))

	// will be 0x30 (ascii "0") to indicate start of package
	// x, _ := buf.ReadByte()
	// printHex("a", x)

	// y, _ := buf.ReadByte()
	// printHex("b", y)

	// t, _ := buf.ReadByte()
	// printHex("t", t)

	buf.Next(3)

	// the size field will be either 1 of 2 bytes
	lenPrime := 0
	var n []byte

	if len(b) < 268 {
		// size is 1 byte
		n = make([]byte, 1, 1)
	} else {
		// size is represented by 2 bytes
		buf.Next(3)
		n = make([]byte, 2, 2)
	}

	// length of data
	if _, err := buf.Read(n); err != nil {
		panic(err)
	}

	if len(n) == 1 {
		// size is in one byte
		lenPrime = int(n[0])
	} else {
		// size is in 2 bytes
		i := binary.BigEndian.Uint16(n)
		lenPrime = int(i)
	}

	// printHex("lenPrime", lenPrime)

	p := make([]byte, lenPrime, lenPrime)
	if _, err := buf.Read(p); err != nil {
		panic("wat")
	}

	// the prime
	prime := new(big.Int)
	prime.SetBytes(p)

	// bytes before the generator
	buf.ReadByte() // 0x02
	buf.ReadByte() // 0x01

	// the generator
	g, _ := buf.ReadByte()
	generator := new(big.Int)
	generator.SetBytes([]byte{g})

	return prime, generator
}

func printHex(label string, v interface{}) {
	fmt.Printf("%s = %v (%#x)\n", label, v, v)
}
