package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}

// StringToBigIntTuple will take the hexadecimal string representing the X and Y of a
// public key and return them in a big Int form separately
func StringToBigIntTuple(s string) (big.Int, big.Int) {
	bx, _ := hex.DecodeString(s[:64])
	by, _ := hex.DecodeString(s[64:])

	var bix big.Int
	var biy big.Int

	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)
	return bix, biy
}

// PublicKeyFromString will call StringToBigIntTuple to get the X and Y values needed for returning
// a new Public Key
func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := StringToBigIntTuple(s)
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
}

func PrivateKeyFromString(s string, pubKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	b, _ := hex.DecodeString(s[:])
	var bi big.Int
	_ = bi.SetBytes(b)
	return &ecdsa.PrivateKey{
		PublicKey: *pubKey,
		D:         &bi,
	}
}
