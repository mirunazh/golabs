//6. Fie un map de tipul string-int care stochează perechi de tipul nume-înălțime. Să se afișeze numele
// persoanelor în ordinea descrescătoare a înălțimilor. Dacă două persoane au aceeași înălțime, nu
// contează ordinea de afișare.
// Exemplu de input:
// map[string]int{
// "Andrei": 190,
// "Maria": 165,
// "Ion": 182,
// "Elena": 172,
// "George": 178,
// "Ana": 165,
// }

package main

import (
	"fmt"
)

func main() {
	catalog := map[string]int32{"Andrei": 190, "Maria": 165, "Ion": 182, "Elena": 172, "George": 178, "Ana": 165}
	inaltimi := []int32{}
	for _, height := range catalog {
		inaltimi = append(inaltimi, height)
	}
	length := len(inaltimi)
	for i := range inaltimi {
		for j := i; j < length; j++ {
			if inaltimi[i] < inaltimi[j] {
				aux := inaltimi[i]
				inaltimi[i] = inaltimi[j]
				inaltimi[j] = aux
			}
		}
	}

	for _, valoare := range inaltimi {
		for nume, hgh := range catalog {
			if hgh == valoare {
				fmt.Println(nume, hgh)
			}
		}
	}
	// fmt.Println(inaltimi)
}
