// 7. Scrieți un program Go care citește un slice de string-uri formate numai din litere mici dintr-un
// fișier. Programul trebuie să afișeze de câte ori apare fiecare cuvânt în slice-ul de string-uri. (1p)
// Exemplu de input: mar para mar strugure banana para mar strugure portocala
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cuvinte := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		cuvinte = append(cuvinte, strings.Fields(string(text))...)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("eroare la citire: ", err)
	}

	frecvente := map[string]int{}

	for _, val := range cuvinte {
		frecvente[val]++
	}

	for cuvant, frecventa := range frecvente {
		fmt.Println(cuvant, frecventa)
	}
}
