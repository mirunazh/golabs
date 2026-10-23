package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	v := scanner.Text()

	x, err := strconv.Atoi(v)

	err = scanner.Err()
	if err != nil {
		fmt.Println("eroare la citire: ", err)
	}

	cn := 0
	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			fmt.Println("neprim")
			cn = 1
			break
		}
	}
	if cn == 0 {
		fmt.Println("prim")
	}
}
