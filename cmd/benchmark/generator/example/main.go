package main

import (
	"fmt"
	"github.com/arjunsk/cometkv/cmd/benchmark/generator"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	keygen := generator.Build(generator.SEQUENTIAL, 0, 100)

	for i := 0; i < 100; i++ {
		key := keygen.Next(r)
		fmt.Println(key)
	}
}
