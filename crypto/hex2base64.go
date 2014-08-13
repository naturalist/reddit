package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func hex2base64(h string) (string, error) {
	s, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(s), nil
}

func fixedXor(xs1, xs2 string) (string, error) {
	if len(xs1) != len(xs2) {
		return "", errors.New("Strings with different length")
	}

	s1, err := hex.DecodeString(xs1)
	if err != nil {
		return "", err
	}

	s2, err := hex.DecodeString(xs2)
	if err != nil {
		return "", err
	}

	res := make([]byte, len(s1))
	for i := 0; i < len(s1); i++ {
		res[i] = s1[i] ^ s2[i]
	}

	return hex.EncodeToString(res), nil
}

func main() {
	s, _ := fixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println(s)
}
