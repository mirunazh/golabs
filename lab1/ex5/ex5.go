// 5. Scrieți un program Go care ia elementele unui slice de string-uri și îl modifică astfel încât elementele să
// apară în ordine inversă. Nu folosiți un alt slice pentru această operațiune. (1p)
// Exemplu: slice := []string"Laboratorul", "de", "sisteme", "distribuite".
// Rezultatul va fi: ["distribuite" "sisteme" "de" "Laboratorul"]

package main

import "fmt"

func main() {
	slice := []string{"string1", "string2", "string3", "string4"}

	for i := 0; i <= len(slice)/2; i++ {
		aux := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = aux
	}
	fmt.Println(slice)
}
