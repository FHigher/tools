package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	/* secp256k1 "github.com/ipsn/go-secp256k1" */
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	//privateKeyBytes := FromECDSA(privateKey)
	// 私钥
	//fmt.Println(privateKeyBytes)

	// 公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := FromECDSAPub(publicKeyECDSA)
	fmt.Println(publicKeyBytes)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
	//fmt.Println(PubkeyToAddress(*publicKeyECDSA))

	fmt.Println(SignAppkey("1608727195", "B218E28A83A9BF8D", "A4BAEC2EAF084C1DB218E28A83A9BF8DA4BAEC2EAF084C1D"))
}

func SignAppkey(timestamp, appid, appkey string) string {
	mac := hmac.New(sha256.New, []byte(appkey))
	mac.Write([]byte(appid + "&" + timestamp + "&" + appkey))

	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}
func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := FromECDSAPub(&p)
	return common.BytesToAddress(Keccak256(pubBytes[1:])[12:])
}

type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := sha3.NewLegacyKeccak256().(KeccakState)
	for _, b := range data {
		d.Write(b)
	}
	d.Read(b)
	return b
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
}

// FromECDSA exports a private key into a binary dump.
func FromECDSA(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	return math.PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}
