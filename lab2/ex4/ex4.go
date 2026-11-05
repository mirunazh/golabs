// 4. Scrieți un program Go în care un goroutine să notifice funcția main() atunci când procesarea sa
// se sfârșește, fără a folosi un wait group. 

package main

import "fmt"

func f(ch chan bool) {
	fmt.Println("gorutina sec")
	
	ch <- true
}

func main() {
	ch := make(chan bool)
	go f(ch)
	

	mesaj := <- ch
	if mesaj == true {
		fmt.Println("gorutina a terminat")
	}
}