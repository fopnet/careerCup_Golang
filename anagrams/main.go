package main

import (
	"fmt"
	"strings"
)

/**
The fundamental theorem of arithmetic states:
Every integer either is a prime number itself or can be
represented as the product of prime numbers and that, moreover,
this representation is unique, the order of the factors.

https://hackernoon.com/an-algorithm-for-finding-anagrams-2fe7655de85b
*/
func main() {

	n := 5

	c := make(chan []int)
	go generatePrimeNumbers(n, c)
	primes := <-c
	fmt.Printf("os %d primos são %v ", n, primes)

	primesMap := generatePrimesHash(primes)
	fmt.Println("primesMap", primesMap)

	anagrams := []string{"god", "dog"}
	fmt.Printf("%s isAnagram of %s ? %t \n", anagrams[0], anagrams[1], isAnagram(anagrams[0], anagrams[1], primesMap))

	anagrams = []string{"elvis", "lives"}
	fmt.Printf("%s isAnagram of %s ? %t \n", anagrams[0], anagrams[1], isAnagram(anagrams[0], anagrams[1], primesMap))

	anagrams = []string{"roast beef", "eat for BSE"}
	fmt.Printf("%s isAnagram of %s ? %t \n", anagrams[0], anagrams[1], isAnagram(anagrams[0], anagrams[1], primesMap))
}

func isAnagram(s1, s2 string, primeHash map[string]int) bool {
	return s1 != s2 && calc(s1, primeHash) == calc(s2, primeHash)
}

func calc(s string, primeHash map[string]int) int {
	total := 1
	for _, c := range strings.ToLower(s) {
		// fmt.Println("c", string(c))
		total *= primeHash[string(c)]
	}
	// fmt.Printf("total of %s is %d\n", s, total)

	return total
}

func generatePrimesHash(numbers []int) map[string]int {
	primesHash := make(map[string]int)

	for i := 97; i <= 122; i++ {
		primesHash[string(i)] = i
	}
	primesHash[" "] = 1

	return primesHash
}

func generatePrimeNumbers(n int, c chan []int) {
	numbers := make([]int, 0, n)

	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				numbers = append(numbers, primo)
				// c <- primo
				inicio = primo + 1
				break
			}
		}
	}

	c <- numbers
	close(c)

}

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
