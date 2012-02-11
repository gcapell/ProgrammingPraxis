package main

import (
	"log"
	"math"
)

func main() {
	log.SetFlags(0) // quieter logging
	log.Print(len(primesLessThan(15485863)))
}

func primesLessThan(n uint64) []uint64 {
	sieve := make([]bool, n)
	var primes []uint64
	sqrtN := uint64(math.Sqrt(float64(n)))
	primes = append(primes, 2)
	for p := uint64(3); p < n; p += 2 {
		if sieve[p] {
			continue
		}
		primes = append(primes, p)
		for j := p * p; j < sqrtN; j += 2 * p {
			sieve[j] = true
		}
	}
	return primes
}
