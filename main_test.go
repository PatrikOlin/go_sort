package main

import (
	"testing"
	"math/rand"
	"time"

	"github.com/PatrikOlin/go_sort/insertionSort"
	"github.com/PatrikOlin/go_sort/mergeSort"
)


func benchmarkInsertionSort(rn []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		insertionSort.Sort(rn)
	}
}

func BenchmarkInsertionSort1k(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Perm(1000)

	benchmarkInsertionSort(rn, b)
}

func BenchmarkInsertionSort10k(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Perm(10000)

	benchmarkInsertionSort(rn, b)
}

func BenchmarkInsertionSort100k(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Perm(100000)

	benchmarkInsertionSort(rn, b)
}

// func BenchmarkInsertionSort1000k(b *testing.B) {
// 	rand.Seed(time.Now().UnixNano())
// 	rn := rand.Perm(1000000)

// 	benchmarkInsertionSort(rn, b)
// }

func benchmarkMergeSort(rn []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		mergeSort.Sort(rn)
	}
}


func BenchmarkMergeSort1k(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Perm(10000)

	benchmarkMergeSort(rn, b)
}

func BenchmarkMergeSort10k(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Perm(10000)

	benchmarkMergeSort(rn, b)
}


func BenchmarkMergeSort100k(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Perm(100000)

	benchmarkMergeSort(rn, b)
}

// func BenchmarkMergeSort1000k(b *testing.B) {
// 	rand.Seed(time.Now().UnixNano())
// 	rn := rand.Perm(1000000)

// 	benchmarkMergeSort(rn, b)
// }
