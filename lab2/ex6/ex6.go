// 6. Fie un sistem distribuit format din doi roboți identici. Ei sunt porniți de către un observer,
// reprezentat de funcția main(). Cei doi roboți pornesc în același timp și se întrec în linie dreaptă
// până la un finish. Scrieți un program Go în care simulați acest sistem, unde funcția main() va
// afișa care robot a ajuns primul la finish. Hint: folosiți select. 2p 

package main
import (
	"fmt"
	"time"
)

func depl(ch chan string, nume string, durata time.Duration) {
	time.Sleep(durata)
	ch <- nume
}

func main() {
	robot1 := make(chan string)
	robot2 := make(chan string)
	
	go depl(robot1, "rb 1", 160 * time.Millisecond)
	go depl(robot2, "rb 2", 120 * time.Millisecond)

	select {
	case <-robot1:
		fmt.Println("rb1")
	case <-robot2:
		fmt.Println("rb2")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation timed out")
	}

}