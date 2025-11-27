package music

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScale(t *testing.T) {
	scale, err := NewScale([]Note{C, D, E, F, G, A, B}, C, "Major", []string{})
	assert.Nil(t, err)
	assert.Equal(t, scale.Root(), C, "Scale root is correct")
	assert.Equal(t, scale.notes, []Note{C, D, E, F, G, A, B}, "Scale note list is correct")
}

func TestNewEmptyScale(t *testing.T) {
	_, err := NewScale([]Note{}, C, "Major", []string{})
	assert.EqualError(t, err, "Notes list cannot be empty")
}

func TestParseValidScale(t *testing.T) {

	input := "C-D-E-F-G-A-B"
	var expectedAlterations []string

	scale, err := ParseScale(input, "Major")

	assert.Nil(t, err)
	assert.Equal(t, scale.Root(), C, "Scale root is correct")
	assert.Equal(t, scale.notes, []Note{C, D, E, F, G, A, B}, "Note list is correct")
	assert.Equal(t, scale.alterations, expectedAlterations, "Alteration list is correct")
}

func TestParseValidScaleWithAlterations(t *testing.T) {
	input := "C-D-E-F#-G-A-Bb"
	expectedAlterations := []string{"F#", "Bb"}

	scale, err := ParseScale(input, "Major")
	assert.Nil(t, err)

	assert.Equal(t, scale.Root(), C, "Scale root is correct")
	assert.Equal(t, scale.notes, []Note{C, D, E, FSharp_Gb, G, A, ASharp_Bb}, "Note list is correct")
	assert.Equal(t, scale.alterations, expectedAlterations, "Alteration list is correct")
}

func TestScaleContainsNote(t *testing.T) {
	input := "C-D-E-F#-G-A-Bb"

	scale, err := ParseScale(input, "Major")
	assert.Nil(t, err)

	assert.True(t, scale.Contains(D))
	assert.True(t, scale.Contains(ASharp_Bb))
	assert.True(t, scale.Contains(FSharp_Gb))
	assert.False(t, scale.Contains(F))
}

func TestResolveEnharmonics(t *testing.T) {
	input := "C-D-E-F#-G-A-Bb"

	scale, err := ParseScale(input, "Major")
	assert.Nil(t, err)

	enharmonicFSharp := scale.ResolveEnharmonic(FSharp_Gb)
	enharmonicBFlat := scale.ResolveEnharmonic(ASharp_Bb)
	enharmonicCSharp := scale.ResolveEnharmonic(CSharp_Db)

	assert.Equal(t, enharmonicFSharp, "F#")
	assert.Equal(t, enharmonicBFlat, "Bb")
	assert.Equal(t, enharmonicCSharp, "C#/Db")
}
