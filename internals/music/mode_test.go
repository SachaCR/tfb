package music

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIonian(t *testing.T) {
	ionianMode := Ionian()
	cMajorScale := ionianMode.ToScale(C)

	assert.Equal(t, C, cMajorScale.Root(), "Scale root is correct")
	assert.Equal(t, []Note{C, D, E, F, G, A, B}, cMajorScale.notes, "Scale note list is correct")
}

func TestDorian(t *testing.T) {
	dorianMode := Dorian()
	dDorianScale := dorianMode.ToScale(D)

	assert.Equal(t, D, dDorianScale.Root(), "Scale root is correct")
	assert.Equal(t, []Note{D, E, F, G, A, B, C}, dDorianScale.notes, "Scale note list is correct")
}

func TestChromatic(t *testing.T) {
	chromaticMode := Chromatic()
	chromaticScale := chromaticMode.ToScale(C)

	assert.Equal(t, C, chromaticScale.Root(), "Scale root is correct")
	fmt.Println(chromaticScale.notes)
	assert.Equal(t, []Note{C, CSharp_Db, D, DSharp_Eb, E, F, FSharp_Gb, G, GSharp_Ab, A, ASharp_Bb, B}, chromaticScale.notes, "Scale note list is correct")
}
