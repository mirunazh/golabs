// 9. Scrieți un program Go care stochează în memorie un array cu numere naturale. Programul va
// citi de la tastatură un număr s și va afișa în O(n) indicii elementelor care însumate au ca rezultat s. (2p)
// Exemplu:
// 15
// 2 7 11 8 4 9
// Rezultatul va fi:
// 1 3
// 2 4

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	intArr := []int{2, 7, 11, 8, 4, 9}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	v := scanner.Text()
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println("eroare la citire: ", err)
	}
	myMap := map[int]int{}
	for i, val := range intArr {
		dif := s - val
		index, exista := myMap[dif]
		if !exista {
			// cheia nu se afla in map, o adaug
			myMap[val] = i
		} else {
			fmt.Println(index, i)
		}
	}
}
