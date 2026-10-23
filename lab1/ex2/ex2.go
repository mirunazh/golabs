// 2. Scrieți un program Go care stochează un array în memorie.
// Programul trebuie să afișeze tabloul sortat crescător folosind o metodă de sortare la alegere.
package main

import "fmt"

func main() {
	tablou := [7]int32{2, 5, 8, 3, 1, 0, -5}

	length := len(tablou)
	for i := range tablou {
		fmt.Println(i)
		for j := i + 1; j < length; j++ {
			if tablou[i] > tablou[j] {
				aux := tablou[i]
				tablou[i] = tablou[j]
				tablou[j] = aux
			}
		}
	}

	fmt.Println(tablou)
}
