package main

import (
	"fmt"

	"github.com/petvajda/fret"
)

func main() {
	bass := []string{"B", "E", "A", "D", "G", "C"}
	for _, s := range bass {
		fmt.Printf("%30s string\n", s)
		fmt.Println(fret.Scramble(s))
		fmt.Println()
	}
}
