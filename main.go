package main

import (
	"fmt"

	"github.com/jayzcrypto/go-learn/formatter"
	"github.com/jayzcrypto/go-learn/math"
)


func main() {
	num := math.Double(2)
	format := print.Format(num)
	fmt.Println(format)
}