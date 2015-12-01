package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sub(x interface{}, y interface{}) int {
	dif := 0
	intX, okX := x.(int)
	intY, okY := y.(int)
	if okX && okY {
		dif = intX - intY
	}
	return dif
}

func main() {
	const Loop = 10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Loop; i++ {
		fmt.Println("now is", rand.Intn(1000), sub(20, i*5))
	}
}
