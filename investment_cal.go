// first program go
package main

import (
	"fmt"
	"math"
)
func main() {
	investAmount := 1000
	returnRate := 5.5
	year := 10
	futureValue := float64(investAmount) * math.Pow((1 + returnRate/100), float64(year))
	fmt.Println(futureValue)
}

// go mod init ....
//https://go.dev/blog/using-go-modules
