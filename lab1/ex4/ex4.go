// 4. Scrieți un program care stochează în memorie un slice de float32 și citește de la tastatură un
// index. Programul trebuie să elimine elementul de la poziția index, păstrând ordinea celorlalte elemente.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := []float32{2.3, 4.6, 9.1}
	slice_cut := []float32{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	v := scanner.Text()
	index, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println("eroare la citire: ", err)
	}

	// slice_cut = append(slice[:index], slice[index+1:]...) // puhnctele fac sa se appenduiasca fiecare elem separat

	for i, val := range slice {
		if i != index {
			slice_cut = append(slice_cut, val)
		}
	}
	fmt.Println(slice_cut)

}
