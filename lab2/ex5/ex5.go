// 5. Scrieți un program Go în care un goroutine(send) trimite un mesaj unui alt goroutine(receive)
// printr-un canal. Receive va returna mesajul primit, iar funcția main() îl va afișa pe ecran. 1p

package main

import "fmt"

func receive (ch <-chan string) string {
	return <-ch
}

func send (ch chan string) {
	ch <- "mesaj trimi de send"
}

func main() {
	ch := make(chan string)
	ch2 := make(chan string)
	go send(ch)
	go func() {
		mesaj := receive(ch)
		ch2 <- mesaj
	}() // aici se transmite argumentul functiei anonime
	mesaj := <- ch2
	fmt.Println(mesaj)
}