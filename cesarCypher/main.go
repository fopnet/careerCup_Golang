package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileHandle, err := os.Open("./cesarCypher/os.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println("n", n)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		fmt.Println(cypher(scanner.Text()))
	}
}

func cypher(texto string) string {
	result := ""
	for _, c := range texto {

		switch c {
		case 32:
			result += string(c)
			break
		case 90:
			result += "C"
			break
		default:
			result += string(c + 3)
		}
	}

	return result
}
