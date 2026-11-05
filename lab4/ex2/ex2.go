// 2. Implementați algoritmul de Flooding pentru grafuri tare conexe

package main

import (
	"fmt"
	"sync"
)

// operatii int nod
func node(id int, index int, n int, graf [][]int, canale [][]chan int, rezultate []int, wg *sync.WaitGroup) {
	defer wg.Done()
	max_id := id
	// runde
	for r := 0; r < n-1; r++ {
		// trim maximul catre vec de ies
		for j := 0; j < n; j++ {
			if graf[index][j] == 1 {
				canale[index][j] <- max_id
			}
		}

		// primim valorile de la vecinii de int
		for j := 0; j < n; j++ {
			if graf[j][index] == 1 {
				valoare := <-canale[j][index]
				if valoare > max_id {
					max_id = valoare
				}
			}
		}
	}

	rezultate[index] = max_id
}

// returneaza idul liderului
func GraphFlooding(ids []int, graf [][]int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}
	wg := sync.WaitGroup{}

	canale := make([][]chan int, n)
	rezultate := make([]int, n)

	// canalele dintre noduri
	for i := 0; i < n; i++ {
		canale[i] = make([]chan int, n)
		for j := 0; j < n; j++ {
			if graf[i][j] == 1 {
				canale[i][j] = make(chan int, 1)
			}
		}
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go node(ids[i], i, n, graf, canale, rezultate, &wg)
	}

	wg.Wait()

	// caut nodul de id max
	for i := 0; i < n; i++ {
		if rezultate[i] == ids[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{3, 7, 2, 9}
	graf := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
		{1, 0, 0, 0},
	}
	lider := GraphFlooding(ids, graf)
	fmt.Println(lider)
}
