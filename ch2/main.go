package main

import (
	tempconv "booklearning/ch2/tempconv"
	"fmt"
)

func main() {

	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	BoilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", BoilingF-tempconv.CToF(tempconv.FreezingC))

	fmt.Printf("BRRRRR! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("BRRRRR! %v\n", tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Printf("it's Kelvin bro %v\n", tempconv.CToK(200))
}
