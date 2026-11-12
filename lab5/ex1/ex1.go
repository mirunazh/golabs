package main

import "fmt"

type Punct struct {
	X, Y float64
}

func calc_media(elem Punct, v []Punct) Punct {
	suma_x := elem.X
	suma_y := elem.Y
	for _, el := range v {
		suma_x += el.X
		suma_y += el.Y
	}
	var rez Punct
	rez.X = suma_x / float64(len(v)+1)
	rez.Y = suma_y / float64(len(v)+1)
	return rez
}

func main() {
	var n int
	fmt.Scan(&n)
	pos := make([]Punct, n)
	for i := range pos {
		fmt.Scan(&pos[i].X, &pos[i].Y)
	} // salveaza citirea in pos
	raza := 16.0
	toler := 0.15

	for runda := 1; runda <= 1000; runda++ {
		// aici trb sa iau fiecare nod si sa verific distanta lui fata de toate celelalte, sa vad care sunt
		// in raza lui si sa elimin pe cele care sunt extraordinar de apropiate, pentru ca atunci spunem ca
		// se confunda, deci sunt considerate un singur punct
		pozitii_noi := make([]Punct, len(pos))
		unite := true
		for i, element := range pos {
			vecini := []Punct{}
			for id, coleg := range pos {
				distanta := (coleg.X-element.X)*(coleg.X-element.X) + (coleg.Y-element.Y)*(coleg.Y-element.Y)
				if distanta > toler*toler {
					unite = false
				}
				if distanta < raza*raza && id != i {
					vecini = append(vecini, coleg)
				}
			}
			pozitii_noi[i] = calc_media(element, vecini)
			// trebuie umva sa adaug un mecanism prin care sa stiu sa opresc rundele.
		}
		if unite == true {
			// inseamna ca toate punctele s-au unit
			break
		}
		pos = pozitii_noi
	}
	fmt.Println(pos)
}
