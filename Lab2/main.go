package main

import (
	"fmt"

	"example.com/app/mathutil"
)

func main() {
	sum := mathutil.Add(3, 4)
	fmt.Printf("The sum of 3 and 4 is: %d\n", sum)
}
