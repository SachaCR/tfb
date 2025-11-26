package music

import (
	"strings"
)

type Scale struct {
	notes       []Note
	root        Note   // Default to first note if missing
	name        string // Optional
	alterations []string
}

func (scale Scale) Contains(note Note) bool {
	noteFound := false

	for _, scaleNote := range scale.notes {
		if scaleNote == note {
			noteFound = true
		}
	}

	return noteFound
}

func (scale Scale) ResolveEnharmonic(note Note) string {
	possibleNote := strings.Split(note.String(), "/")

	for _, alteration := range scale.alterations {
		if alteration == possibleNote[0] || alteration == possibleNote[1] {
			return alteration
		}
	}

	return note.String()
}

func (scale Scale) Root() Note {
	return scale.root
}

func NewScale(notes []Note, root Note, name string, alterations []string) Scale {
	return Scale{
		notes:       notes,
		root:        root,
		name:        name,
		alterations: alterations,
	}
}

func ParseScale(scaleString string, scaleName string) Scale {
	stringSlice := strings.Split(scaleString, "-")
	var noteSlice []Note
	var alterations []string

	for _, string := range stringSlice {
		note, ok := NoteFromString[string]

		if !ok {
			panic("invalid note")
		}

		if strings.Contains(note.String(), "/") {
			alterations = append(alterations, string)
		}

		noteSlice = append(noteSlice, note)
	}

	var scale Scale
	rootNote := noteSlice[0]

	scale = NewScale(noteSlice, rootNote, scaleName, alterations)

	return scale
}

//
// var scaleRegexp = regexp.MustCompile(`^([A-G](?:#|b)?)(?:-([A-G](?:#|b)?))*(?::([A-Za-z0-9]+))?$`)
//
// func ParseScale(input string) (*Scale, error) {
// 	matches := scaleRegexp.FindStringSubmatch(input)
//
// 	if matches == nil {
// 		return nil, fmt.Errorf("Invalid scale format: %q", input)
// 	}
//
// 	root := NoteFromString[matches[1]]
// 	notes := []Note{NoteFromString[matches[1]]}
//
// 	println("ROOT:", root.String())
// 	fmt.Println("matches[2]", matches)
//
// 	// Collect all captures of the second group (the rest of the notes)
// 	for _, m := range matches[2 : len(matches)-1] {
// 		if m != "" {
// 			notes = append(notes, NoteFromString[m])
// 		}
// 	}
//
// 	name := matches[len(matches)-1] // group 4 (can be empty)
//
// 	return &Scale{
// 		notes: notes,
// 		root:  root,
// 		name:  name,
// 	}, nil
// }
