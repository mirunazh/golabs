// 7. Fie un sistem distribuit format din doi roboți identici. Ei sunt porniți de către un observer,
// reprezentat de funcția main(). Fiecare robot va parcurge un traseu și va trimite periodic printrun 
// canal coordonatele sale 2D către celălalt robot, în timp ce se deplasează. Când un robot
// își termină de parcurs traseul, îl va notifica pe celălalt că s-a oprit. Acesta trebuie să citească
// toate pozițiile primite până la oprire, iar apoi ambii vor raporta observer-ului că și-au terminat
// sarcina. Simulați, folosind un program Go, comportamentul roboților. 2p


package main
import "fmt"
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	
}