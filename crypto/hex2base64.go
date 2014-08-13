package main

import (
    "fmt"
    "strconv"
    "encoding/base64"
)

func main() {
    hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    s := make([]byte, len(hex) / 2)
    for i := 0; i < len(hex) / 2; i++ {
        val, _ := strconv.ParseInt(hex[i * 2:i * 2 + 2], 16, 8)
        s[i] = byte(val)
    }
    fmt.Println(base64.StdEncoding.EncodeToString(s))
}
