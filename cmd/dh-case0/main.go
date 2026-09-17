package main

func main() {
	panic("commented out")
}

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"fmt"
// 	"io"

// 	"github.com/btcsuite/btcd/btcec"
// 	"github.com/btcsuite/btcd/chaincfg"
// 	"github.com/btcsuite/btcd/wire"
// 	"github.com/btcsuite/btcutil"
// 	"github.com/monnand/dhkx"
// )

// const (
// 	// MainNet represents the magic bytes for BSV mainnet.
// 	MainNet wire.BitcoinNet = 0xe8f3e1e3
// )

// var (
// 	// MainNetParams defines the network parameters for the BSV mainnet.
// 	MainNetParams = chaincfg.MainNetParams
// )

// func main() {
// 	group, err := dhkx.GetGroup(14)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("group : P = 0x%s\n", group.P().Text(16))
// 	fmt.Printf("group : G = 0x%s\n", group.G().Text(16))

// 	p1 := newPeer(group)
// 	p2 := newPeer(group)

// 	if err := exchangeKey(p1, p2); err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("p1 private key = %s\n", p1.priv.String())
// 	fmt.Printf("p2 private key = %s\n", p2.priv.String())

// 	// ecdsa.PrivateKey
// 	priv1, _ := btcec.PrivKeyFromBytes(btcec.S256(), p1.priv.Bytes())
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("len priv.D = %v\n", len(priv1.D.Bytes()))

// 	plaintext := "message in a bottle"

// 	// wif, err := PrivKeyToWIF(priv1, false, &MainNetParams)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// encrypted, err := message.Sign(priv1, []byte("message in a bottle"))
// 	encrypted, err := EncryptAES(priv1.D.Bytes(), []byte(plaintext))
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("encrypted = %s\n", encrypted)

// 	//
// }

// type peer struct {
// 	priv  *dhkx.DHKey
// 	group *dhkx.DHGroup
// 	pub   *dhkx.DHKey
// }

// func newPeer(g *dhkx.DHGroup) *peer {
// 	ret := new(peer)
// 	ret.priv, _ = g.GeneratePrivateKey(nil)
// 	ret.group = g
// 	return ret
// }

// func (self *peer) getPubKey() []byte {
// 	return self.priv.Bytes()
// }

// func (self *peer) recvPeerPubKey(pub []byte) {
// 	pubKey := dhkx.NewPublicKey(pub)
// 	self.pub = pubKey
// }

// func (self *peer) getKey() []byte {
// 	k, err := self.group.ComputeKey(self.pub, self.priv)
// 	if err != nil {
// 		return nil
// 	}

// 	return k.Bytes()
// }

// func exchangeKey(p1, p2 *peer) error {
// 	pub1 := p1.getPubKey()
// 	pub2 := p2.getPubKey()

// 	p1.recvPeerPubKey(pub2)
// 	p2.recvPeerPubKey(pub1)

// 	key1 := p1.getKey()
// 	if key1 == nil {
// 		return fmt.Errorf("p1 has nil key")
// 	}

// 	key2 := p2.getKey()

// 	if key2 == nil {
// 		return fmt.Errorf("p2 has nil key")
// 	}

// 	for i, k := range key1 {
// 		if key2[i] != k {
// 			return fmt.Errorf("%vth byte does not same")
// 		}
// 	}

// 	// if bytes.Compare(key1, key2) != 0 {
// 	// 	return errors.New("keys do not match")
// 	// }

// 	return nil
// }

// func PrivKeyToWIF(privateKey *btcec.PrivateKey, compressed bool, params *chaincfg.Params) (*string, error) {
// 	w, err := btcutil.NewWIF(privateKey, params, compressed)
// 	if err != nil {
// 		return nil, err
// 	}

// 	wif := w.String()

// 	return &wif, nil
// }

// func WIFToPrivKey(s string) (*btcec.PrivateKey, error) {
// 	wif, err := btcutil.DecodeWIF(s)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return wif.PrivKey, nil
// }

// func EncryptAES(key []byte, plaintext []byte) ([]byte, error) {
// 	// create cipher
// 	c, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	gcm, err := cipher.NewGCM(c)
// 	if err != nil {
// 		panic(err)
// 	}

// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		fmt.Println(err)
// 	}

// 	// encrypt
// 	out := gcm.Seal(nonce, nonce, plaintext, nil)

// 	return out, nil
// }

// func DecryptAES(key []byte, ciphertext []byte) ([]byte, error) {
// 	c, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	gcm, err := cipher.NewGCM(c)
// 	if err != nil {
// 		panic(err)
// 	}

// 	nonceSize := gcm.NonceSize()
// 	if len(ciphertext) < nonceSize {
// 		fmt.Println(err)
// 	}

// 	nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize:]
// 	plaintext, err := gcm.Open(nil, nonce, encryptedMessage, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return plaintext, nil
// }
