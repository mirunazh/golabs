// 2. Scrieți un program Go care să transfere cu success un mesaj (de exemplu „salut”) de la goroutina
// main la alta (diferită de main()) prin intermediul unui channel.
package main

import (
	"fmt"
	"sync"
)

func grt(ch chan string, wg *sync.WaitGroup) {
	var i = <-ch
	fmt.Println(i)
	wg.Done()
}

func main() {
	var ch = make(chan string, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go grt(ch, &wg)
	ch <- "cuvant"
	wg.Wait()
}
