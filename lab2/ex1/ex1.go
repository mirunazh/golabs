// 1. Scrieți un program Go care să transfere cu succes un mesaj (de exemplu „salut”) de la o goroutină
// (diferită de main()) la goroutina main() prin intermediul unui channel.

package main

import "fmt"

func grt(ch chan string, cuv string) {
	ch <- cuv
}

func main() {
	var ch = make(chan string)
	go grt(ch, "cuvant")
	var i = <-ch
	fmt.Println(i)
}
