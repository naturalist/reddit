package main

import (
    "fmt"
    "flag"
    "os"
)

var (
    sizeX, sizeY int
    X, Y int
    rules Rules
    grid Grid
)

type Grid []byte

func (g *Grid) init( x, y int) {
    *g = make(Grid, x * y)
}

func (g *Grid) set(x, y int, color byte) error {
    if x < 0 || x >= sizeX || y < 0 || y >= sizeY {
        return fmt.Errorf("Set out of bounderies [%v, %v]", x, y)
    }

    (*g)[x * sizeX + y] = color
    return nil
}

func (g *Grid) get(x, y int) byte {
    if x < 0 || x >= sizeX || y < 0 || y >= sizeY {
        return 0
    }

    return (*g)[x * sizeX + y]
}

func (g *Grid) show() {
    for x := 0; x < sizeX; x++ {
        for y := 0; y < sizeY; y++ {
            var char byte
            switch val := g.get(x, y); {
                case val == 0: char = '.'
                case x == X && y == Y: char = 'X'
                default: char = 'A' - 1 + val
            }
            fmt.Printf("%v ", string(char))
        }
        fmt.Println()
    }
}

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

func init() {
    flag.IntVar(&sizeX, "sizex", 11, "x size of the grid")
    flag.IntVar(&sizeY, "sizey", 11, "y size of the grid")
    flag.IntVar(&X, "x", 5, "starting x position of the ant")
    flag.IntVar(&Y, "y", 5, "starting y position of the ant")
    flag.Parse()

    if ( flag.NArg() < 1 ) {
        fmt.Println("Ants simulation v0.1.\n");
        fmt.Println("Usage: <flags> rules\n");
        fmt.Println("Flags:\n");
        flag.PrintDefaults()
        os.Exit(1)
    }

    err := rules.read(flag.Arg(0))
    if ( err != nil ) {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }

    grid.init(sizeX, sizeY)
}

func main() {
    grid.show()

}

