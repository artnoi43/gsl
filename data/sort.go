package data

type ArraySortDirection uint8

const (
	Ascending ArraySortDirection = iota
	Descending
)

func (d ArraySortDirection) IsValid() bool {
	switch d {
	case Ascending, Descending:
		return true
	}

	return false
}
