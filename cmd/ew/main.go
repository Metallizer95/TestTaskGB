package main

import (
	"fmt"
	"math/big"
)

func main() {
	str := "e23456a29054d2094"
	c, hz, err := big.ParseFloat(str, 16, 0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Float64())
	fmt.Println(hz)
}
