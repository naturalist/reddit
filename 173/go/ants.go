package main

import (
    "fmt"
    "flag"
)

const (
    white = iota
    black
    red
    green
    blue
)

var (
    sizeX, sizeY int
)

func init() {
    flag.IntVar(&sizeX, "x", 12, "x size of the grid")
    flag.IntVar(&sizeY, "y", 12, "y size of the grid")
    flag.Parse()
}

func main() {
    fmt.Println(sizeX, sizeY)
}
