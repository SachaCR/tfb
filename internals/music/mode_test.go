package music

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIonian(t *testing.T) {
	ionianMode := FindMode("Ionian")
	cMajorScale := ionianMode.ToScale(C)

	assert.Equal(t, C, cMajorScale.Root(), "Scale root is correct")
	assert.Equal(t, []Note{C, D, E, F, G, A, B}, cMajorScale.notes, "Scale note list is correct")
}

func TestDorian(t *testing.T) {
	dorianMode := FindMode("Dorian")
	dDorianScale := dorianMode.ToScale(D)

	assert.Equal(t, D, dDorianScale.Root(), "Scale root is correct")
	assert.Equal(t, []Note{D, E, F, G, A, B, C}, dDorianScale.notes, "Scale note list is correct")
}
