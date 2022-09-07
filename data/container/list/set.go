package list

type SetImpl[T comparable] struct {
	haystack   []T
	duplicates map[T]struct{}
	length     int
}

// NewSet[T] returns a new Set[T]. The values in src is iterated over and pushed in to the Set[T]
// according to the set rules.
func NewSet[T comparable](src []T) Set[T] {
	var haystack []T
	duplicates := make(map[T]struct{})
	var length int
	for _, item := range src {
		if _, found := duplicates[item]; !found {
			haystack = append(haystack, item)
			duplicates[item] = struct{}{}
			length++
		}
	}

	return &SetImpl[T]{
		haystack:   haystack,
		duplicates: duplicates,
		length:     length,
	}
}

// O(1)
func (s *SetImpl[T]) HasDuplicate(x T) bool {
	_, found := s.duplicates[x]
	return found
}

func (s *SetImpl[T]) Push(x T) {
	if !s.HasDuplicate(x) {
		s.haystack = append(s.haystack, x)
		s.duplicates[x] = struct{}{}
		s.length++
	}
}

func (self *SetImpl[T]) PushSlice(slice []T) {
	for _, elem := range slice {
		self.Push(elem)
	}
}

func (s *SetImpl[T]) Pop() *T {
	toPop := s.haystack[s.length-1]
	s.length--
	return &toPop
}

func (s *SetImpl[T]) Len() int {
	return s.length
}

func (s *SetImpl[T]) IsEmpty() bool {
	return s.length == 0
}
