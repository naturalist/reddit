package main

import (
    "fmt"
    "flag"
    "os"
)

const (
    white = iota
    black
    red
    green
    blue
)

// New type to handle the direction rules
type Rules []int8

// Reads the rules from a string
func (r *Rules) read(s string) (err error) {
    result := make(Rules, len(s))
    for i, char := range s {
        if char == 'L' {
            result[i] = -1
        } else if ( char == 'R' ) {
            result[i] = 1
        } else {
            return fmt.Errorf("Unrecognized direction: %c", char)
        }
    }
    *r = result
    return nil
}

var (
    sizeX, sizeY int
    rules Rules
)

func init() {
    flag.IntVar(&sizeX, "x", 12, "x size of the grid")
    flag.IntVar(&sizeY, "y", 12, "y size of the grid")
    flag.Parse()

    if ( flag.NArg() < 1 ) {
        fmt.Println("Ants simulation v0.1.\n");
        fmt.Println("Usage: <flags> rules\n");
        fmt.Println("Flags:\n");
        flag.PrintDefaults()
        os.Exit(1)
    }

}

func main() {
    err := rules.read(flag.Arg(0))
    if ( err != nil ) {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    fmt.Println(rules);

}

