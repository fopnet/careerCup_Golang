package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fileHandle, err := os.Open("./matrizT/matriz.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)

	scanner.Scan()
	fmt.Println("size", scanner.Text())

	s1 := make([][]string, 0, 4)

	m2 := make([][]string, 4)
	for i := 0; i < len(m2); i++ {
		m2[i] = make([]string, 0, 4)
	}

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(scanner.Text())
		s1 = append(s1, strings.Split(line, " "))

	}
	fmt.Println(s1)

	for _, v := range s1 {
		for j, v2 := range v {
			// fmt.Println("s,v", s,v)
			// sv, _ := strconv.Atoi(v2)
			m2[j] = append(m2[j], v2)
		}
	}

	// fmt.Println(m2)
	for _, l := range m2 {
		for k, v := range l {
			fmt.Print(v)
			if k < 3 {
				fmt.Print(" ")
			}
		}

		fmt.Println("")
	}

}
