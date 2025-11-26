package music

type Scale struct {
	notes []Note
	root  Note   // Default to first note if missing
	name  string // Optional
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

func (scale Scale) Root() Note {
	return scale.root
}

func NewScale(notes []Note, root Note, name string) Scale {
	return Scale{
		notes: notes,
		root:  root,
		name:  name,
	}
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
