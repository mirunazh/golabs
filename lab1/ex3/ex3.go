// 3. Scrieți un program Go care stochează în memorie un șir de numere și care afișează un slice cu
// numerele pare din acel șir.

package main

import "fmt"

func main() {
	tablou := [7]int32{1, 2, 3, 4, 5, 67, 9}
	pare := []int32{}

	for _, val := range tablou {
		if val%2 == 0 {
			pare = append(pare, val)
		}
	}
	fmt.Println(pare)
}
