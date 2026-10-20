package rsa2

import (
	"crypto"
	"fmt"
)

func Main() {
	prvkey, pubkey, err := GenRsaKey(2048)
	if err != nil {
		panic(err)
	}

	data := []byte("foo")

	// Encrypt data.
	cipher, err := RsaEncrypt(pubkey, data)
	if err != nil {
		panic(err)
	}
	if len(cipher) != 2048/8 {
		panic("cipher len not equal to key length")
	}

	// Decrypt data.
	plain, err := RsaDecrypt(prvkey, cipher)
	if err != nil {
		panic(err)
	}
	if string(data) == string(plain) {
		fmt.Printf("rsa encrypt and decrypt succeeded, data:%v plain:%v\n", string(data), string(plain))
	}

	// Using SHA256 to hash msg and then use rsa private key to sign.
	sig, err := RsaSign(prvkey, crypto.SHA256, data)
	if err != nil {
		panic(err)
	}
	if len(sig) != 2048/8 {
		panic("signature len not equal to key length")
	}

	// Using public key to verify signature.
	err = RsaVerifySign(pubkey, crypto.SHA256, data, sig)
	if err != nil {
		panic(err)
	}
	fmt.Println("verify signature succeeded")
}
