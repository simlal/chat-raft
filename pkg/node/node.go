package node

import (
	"fmt"
	"math/rand"
)

func Initiliaze() uint8 {
	fmt.Println("Initializing node...")

	num := rand.Intn(10)

	if num%2 != 0 {
		fmt.Println("odd node!")
	} else {
		fmt.Println("even node!")
	}
	return uint8(num)
}
