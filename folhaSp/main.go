package main

import (
	"fmt"
)

func main() {
	planets := []string{"SIRIUS", "LALANDE", "PROCION", "ALPHA CENTAURI", "BARNARD"}

	engineers := 75
	targetStart := targetStart(engineers, planets)
	fmt.Printf("O target planet is %s\n", targetStart)

}

func targetStart(engineers int, planets []string) string {

	fibonacci := fiboacciMap([]rune{'A', 'E', 'I', 'O', 'U'})

	var targetStar string
	for _, p := range planets {
		total := 1
		for _, c := range p {
			if fibonacci[c] != 0 {
				total *= fibonacci[c]
			}
		}
		if total == engineers {
			targetStar = p
			// fmt.Printf("O target planet is %s\n", p)
		}
	}

	return targetStar
}

func fibo(n int) int {

	if n <= 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}

}

func fiboacciMap(vowels []rune) map[rune]int {
	fibonacci := map[rune]int{}

	for i := 0; i < len(vowels); i++ {
		fmt.Printf("fibo of %s is %d\n", string(vowels[i]), fibo(i+2))
		fibonacci[vowels[i]] = fibo(i + 2)
	}

	fmt.Println("fibonacci 2", fibonacci)

	return fibonacci
}
