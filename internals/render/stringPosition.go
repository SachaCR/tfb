package render

type StringPosition int

const (
	Middle StringPosition = iota
	Top
	Bottom
)

func (n StringPosition) String() string {
	return [...]string{
		"Top",
		"Middle",
		"Bottom",
	}[n]
}
