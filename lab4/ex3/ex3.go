// 3. Implementați algoritmul de Flooding de medie pentru vot majoritar.

package main

import (
	"fmt"
	"sync"
)

// realizeaza operatiile interne nodului
func node(id int, valoare float64, stanga chan float64, dreapta chan float64, inStanga chan float64, inDreapta chan float64, rezultate []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 30; i++ {
		// trim valoarea catre vecini
		stanga <- valoare
		dreapta <- valoare
		// prim valorile vecinilor
		v1 := <-inStanga
		v2 := <-inDreapta

		valoare = (valoare + v1 + v2) / 3
	}
	rezultate[id] = valoare
}

// media cu flooding
func AverageFlooding(voturi []float64) float64 {
	n := len(voturi)
	if n < 3 {
		return -1
	}
	wg := sync.WaitGroup{}

	dreapta := make([]chan float64, n)
	stanga := make([]chan float64, n)
	rezultate := make([]float64, n)

	for i := 0; i < n; i++ {
		dreapta[i] = make(chan float64, 1)
		stanga[i] = make(chan float64, 1)
	}

	wg.Add(n)

	for i := 0; i < n; i++ {
		go node(
			i,
			voturi[i],
			stanga[i],
			dreapta[i],
			dreapta[(i-1+n)%n],
			stanga[(i+1)%n],
			rezultate,
			&wg,
		)
	}
	wg.Wait()
	// medie rez
	suma := 0.0
	for i := 0; i < n; i++ {
		suma += rezultate[i]
	}
	return suma / float64(n)
}

func main() {
	voturi := []float64{1, 0, 1, 1, 0}
	medie := AverageFlooding(voturi)
	fmt.Println("medie ", medie)
	if medie > 0.5 {
		fmt.Println("vot majoritar da")
	} else if medie < 0.5 {
		fmt.Println("vot majoritar nu")
	} else {
		fmt.Println("egal")
	}
}
