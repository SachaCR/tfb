package music

type Scale struct {
	root  Note
	notes []Note
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

func NewScale(notes []Note, root Note) Scale {
	return Scale{
		notes: notes,
		root:  root,
	}
}
