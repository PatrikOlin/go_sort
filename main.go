package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PatrikOlin/go_sort/insertionSort"
	"github.com/PatrikOlin/go_sort/mergeSort"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Perm(5000000)

	defer elapsed("InsertionSort")()
	go insertionSort.Sort(n)

				
	defer elapsed("MergeSort")()
	go mergeSort.Sort(n)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}


				
