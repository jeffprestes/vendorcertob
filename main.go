package main

import (
	"fmt"

	"github.com/jeffprestes/vendorcalc"
)

func main() {
	a := 3
	b := 1
	c := 1
	fmt.Println("Quero o Delta também: ", vendorcalc.Delta(a, b, c))
}
