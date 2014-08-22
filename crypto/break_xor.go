package main

import (
	"encoding/base64"
	//"encoding/hex"
	"fmt"
	"io/ioutil"
	"sort"
)

type Distance struct {
	keyLen   int
	distance float32
}

type DistanceList []Distance
func (d DistanceList) Len() int           { return len(d) }
func (d DistanceList) Less(i, j int) bool { return d[i].distance < d[j].distance }
func (d DistanceList) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

type Freq struct {
	char byte
	freq uint
}

type FreqList []Freq
func (d FreqList) Len() int           { return len(d) }
func (d FreqList) Less(i, j int) bool { return d[i].freq > d[j].freq }
func (d FreqList) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }


type Buffer []byte

// Read base64 encoded file into a Buffer
func NewBufferBase64(filename string) (b Buffer, err error) {
	s, _ := ioutil.ReadFile(filename)
	b, err = base64.StdEncoding.DecodeString(string(s))
	return b, err
}

// Get the normalized Hamming distance in bits for slices a and b
func GetEditDistance(a, b Buffer) (float32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Both buffers must have the same length")
	}

	var diff int
	ln := len(a)

	for i := 0; i < ln; i++ {
		xor := a[i] ^ b[i]
		for xor > 0 {
			diff++
			xor &= xor - 1
		}
	}

	return float32(diff) / float32(ln), nil
}

// Get all Hamming distance for keys sizes from min to max
// Returns a sorted list of Distance types
func (a Buffer) FindDistances(min, max int) (result DistanceList) {
	result = make(DistanceList, (max-min)+1)
	for i := min; i <= max; i++ {
		distance, _ := GetEditDistance(a[:i], a[i:i*2])
		result[i-min] = Distance{i, distance}
	}

	sort.Sort(result)

	return
}

// Splits the buffer into several transposed buffers, depending on the
// keyLen
func (a Buffer) Transpose(keyLen int) (blocks []Buffer) {
	blocks = make([]Buffer, keyLen)
	for i := 0; i < keyLen; i++ {
		blocks[i] = make(Buffer, 0, len(a)/keyLen+1)
		for j := i; j < len(a); j += keyLen {
			blocks[i] = append(blocks[i], a[j])
		}

	}
	return
}

func (a Buffer) LetterFreqs() (result FreqList) {
	m := make(map[byte]uint)
	for _, char := range a {
		m[char]++
	}
	result = make(FreqList, 0, len(m))
	for key, val := range m {
		result = append(result, Freq{key, val})
	}

	sort.Sort(result)
	return
}

func (a Buffer) SingleByteXor(top byte) byte {

	// Find the most used byte
	var (
		max  uint
		char byte
	)
	m := make(map[byte]uint)
	for _, b := range a {
		if b == 0 {
			continue
		}
		m[b]++
		if m[b] > max {
			max = m[b]
			char = b
		}
	}

	// XOR with the top char (usually 0x20)
	return char ^ top
}

func (a Buffer) RepeatingXor(key string) (result Buffer) {
	if len(key) == 0 {
		panic("Key length can not be zero")
	}
	result = make(Buffer, 0, len(a))
	j := 0
	for i := 0; i < len(a); i++ {
		result = append(result, a[i]^key[j])
		j++
		j %= len(key)
	}

	return
}

func (a Buffer) String() string {
	return string(a)
}

func (a Buffer) Hex() string {
	return fmt.Sprintf("%X", a)
}

func main() {
	buf, err := NewBufferBase64("6.txt")
	if err != nil {
		panic(err)
	}

	// Get a list of Distance structs, sorted by smallest distance to largest
	distances := buf.FindDistances(2, 40)
	for i := 0; i < 1; i++ {
		keyLen := distances[i].keyLen
		fmt.Printf("Trying key length: %v\n", keyLen)

		// Get keyLen number of transposed blocks
		blocks := buf.Transpose(keyLen)

		freq := blocks[4].LetterFreqs()
		fmt.Println(freq)
	}
}
