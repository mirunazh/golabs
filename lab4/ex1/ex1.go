// 1. Implementați algoritmul de Flooding pentru un inel unidirecțional.

package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool
}

// realizeaza operatiile interne nodului
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// fiecare nod trimite initial idul lui
	out <- Msg{Value: id, Done: false}

	// citire pe canalul de input
	for {
		stanga := <-in

		if stanga.Done {
			// mesajul de final a ajuns inapoi la lider
			if *isLeader {
				return
			}
			out <- stanga
			return
		}

		trimite := Msg{}
		trimite.Done = false

		if id > stanga.Value {
			// id ul primit este mai mic si nu il trim
			continue
		} else if id == stanga.Value {
			// id ul s a intors, deci nodul este lider
			*isLeader = true
			trimite.Value = id
			trimite.Done = true
		} else {
			// trim id ul mai mare
			trimite.Value = stanga.Value
		}

		out <- trimite
	}
}

// returneaza id ul liderului
func RingFlooding(ids []int) int {
	wg := sync.WaitGroup{}
	n := len(ids)
	if n == 0 {
		return -1
	}

	channel := make([]chan Msg, n)
	isLeader := make([]bool, n)

	for i := 0; i < n; i++ {
		channel[i] = make(chan Msg, n+1)
	}
	wg.Add(n)

	for i := 0; i < n; i++ {
		in := channel[i]
		out := channel[(i+1)%n]
		go node(ids[i], in, out, &isLeader[i], &wg)
	}

	wg.Wait()
	for i := 0; i < n; i++ {
		if isLeader[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{3, 7, 2, 9, 5}
	lider := RingFlooding(ids)
	fmt.Println("Liderul este:", lider)
}
