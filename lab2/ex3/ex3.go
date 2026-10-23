// 3. Scrieți un program Go în care funcția main() primește două mesaje pe același canal de comunicare.
package main

import "fmt"

func grt(ch chan string, cuv1 string, cuv2 string) {
	ch <- cuv1
	ch <- cuv2
	close(ch)
}

func main() {
	var ch = make(chan string, 2)
	go grt(ch, "cuvant1", "cuvant2")
	var i = <-ch
	var j = <-ch
	fmt.Println(i, j)
}
