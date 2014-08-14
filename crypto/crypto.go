package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
)

/* Letters type is used to sort a map   */
type Letter struct {
	key   byte
	value byte
}

type Letters []Letter

func (l Letters) Len() int           { return len(l) }
func (l Letters) Less(i, j int) bool { return l[i].value > l[j].value }
func (l Letters) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

/****************************************/

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

func singleByteXor(xs string) (string, byte) {
	s, err := hex.DecodeString(xs)
	if err != nil {
		panic(err)
	}

	// Count the letter occurences
	m := make(map[byte]byte)
	for _, char := range s {
		m[byte(char)]++
	}

	// Move all keys and values in an array
	letters := make(Letters, 0, len(s))
	for key, value := range m {
		letters = append(letters, Letter{key, value})
	}

	// Sort that array by value (reverse)
	sort.Sort(letters)

	// The most used symbol is the space.
	// XOR the most used key with it to find the code
	code := letters[0].key ^ ' '

	// XOR the entire array with the code
	for i, _ := range s {
		s[i] ^= code
	}

	return string(s), code
}

func main() {
	s, code := singleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(s)
	fmt.Printf("%x\n", code)
}
