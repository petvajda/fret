package fret

import (
	"fmt"
	"testing"
)

func TestGetChromaticNotes(t *testing.T) {
	notes := GetChromaticNotes("C")
	if len(notes) != 13 {
		t.Errorf("GetChromaticNotes should return 13 notes not %d", len(notes))
	}
	expectedNotes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C"}
	for i := 0; i < len(notes); i++ {
		if notes[i] != expectedNotes[i] {
			t.Errorf("GetChromaticNotes want %s got %s", notes[i], expectedNotes[i])
		}
	}
	e := GetChromaticNotes("R")[0]
	if e != "Error: Wrong note name R" {
		t.Errorf("GetChromaticNotes should return an error got %s", e)
	}
}

func TestBassGuitarStrings(t *testing.T) {
	var tests = []struct {
		s    string
		want []string
	}{
		{"E", []string{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"}},
		{"A", []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}},
		{"D", []string{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"}},
		{"G", []string{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s string", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := GetChromaticNotes(tt.s)
			for i := 0; i < len(tt.want); i++ {
				if tt.want[i] != ans[i] {
					t.Errorf("GetChromaticNotes want %s got %s", tt.want[i], ans[i])
				}
			}
		})
	}
}

func TestScramble(t *testing.T) {

	var tests = []struct {
		s    string
		want string
	}{
		{"E", "10:D   3:G   7:B   2:F#  9:C#  5:A  11:D#  1:F   8:C   4:G# 12:E   6:A# "},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s string", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := Scramble(tt.s)
			if tt.want != ans {
				t.Errorf("Scramble want %s got %s", tt.want, ans)
			}
		})
	}
}
