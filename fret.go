package fret

import "fmt"

// Returns all chromatic notes starting with baseNote
func GetChromaticNotes(baseNote string) []string {
	notes := []string{"C", "D", "E", "F", "G", "A", "B"}
	chromaticNotes := make([]string, 12)
	for i, j := 0, 0; i < len(notes); i++ {
		chromaticNotes[j] = notes[i]
		j++
		switch i {
		case 2, 6:
			break
		default:
			chromaticNotes[j] = notes[i] + "#"
			j++
		}
	}
	for i, n := range chromaticNotes {
		if n == baseNote {
			orderedNotes := append(chromaticNotes[i:], chromaticNotes[:i]...)
			orderedNotes = append(orderedNotes, baseNote)
			return orderedNotes
		}
	}
	return []string{"Error: Wrong note name " + baseNote}
}

// Scramble and returns notes for baseNote
func Scramble(baseNote string) string {
	frets := [...]int{10, 3, 7, 2, 9, 5, 11, 1, 8, 4, 12, 6}
	notes := GetChromaticNotes(baseNote)
	ret := ""
	for _, n := range frets {
		ret += fmt.Sprintf("%2d:%-3s", n, notes[n])
	}
	return ret
}
