package neck

import (
	"github.com/SachaCR/tfb/internals/music"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomNeck(t *testing.T) {
	neck, err := NewCustom("Guitare 7", "C-D-E-F-G-A-B")
	assert.Nil(t, err, "No error")

	assert.Equal(t, neck.Instrument(), "Guitare 7", "Instrument name is correct")
	assert.Equal(t, neck.Tuning(), []music.Note{music.C, music.D, music.E, music.F, music.G, music.A, music.B}, "Tuning notes are correct")
}
func TestNewInvalidCustomNeck(t *testing.T) {
	_, err := NewCustom("Guitare 7", "C-D-E-K-G-A-B")
	assert.Error(t, err, "It returns an error because of the invalid note")
}
