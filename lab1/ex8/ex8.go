// 8. Angajații unei firme sunt stocați într-un dicționar map[string]int cu perechi de forma nume - id_departament.
// Creați un nou map[int][]string care să stocheze perechi de tipul id_departamentslice cu numele angajaților
// din departament. (1p)
// Exemplu:
// map[string]int{
// "Ana Popescu": 1,
// "Mihai Ionescu": 2,
// "Bogdan Radu": 1,
// "Elena Dumitrescu": 3,
// "Cristian Vasile": 2,
// "Ioana Enache": 1,
// 9
// }
// Rezultatul va fi:
// Departamentul 1: [Ana Popescu, Bogdan Radu, Ioana Enache]
// Departamentul 2: [Mihai Ionescu, Cristian Vasile]
// Departamentul 3: [Elena Dumitrescu]

package main

import (
	"fmt"
	"strings"
)

func main() {
	angajati := map[string]int{"Ana Popescu": 1, "Mihai Ionescu": 2, "Bogdan Radu": 1, "Elena Dumitrescu": 3, "Cristian Vasile": 2, "Ioana Enache": 1}
	id_departamentslice := map[int][]string{}
	for nume, id := range angajati {
		id_departamentslice[id] = append(id_departamentslice[id], nume)
	}
	for dep, ang := range id_departamentslice {
		fmt.Printf("Departamentul %d: %s\n", dep, strings.Join(ang, ", "))
	}
}
