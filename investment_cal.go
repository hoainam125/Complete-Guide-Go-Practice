// first program go
package main

import (
	"fmt"
	"math"
)
func main() {
	var investAmount = 1000
	var returnRate = 5.5
	var year = 10
	var futureValue = float64(investAmount) * math.Pow((1 + returnRate/100), float64(year))
	fmt.Println(futureValue)
}

// go mod init ....
//https://go.dev/blog/using-go-modules
