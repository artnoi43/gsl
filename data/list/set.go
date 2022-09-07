package list

// SetList wraps BasicList[T] field `basicList`,
// and maintains a hash map of seen items and the basicList length
// so that determining duplicates and getting the length take O(1) time.
type SetList[T comparable, L BasicList[T]] struct {
	basicList  L
	duplicates map[T]struct{}
	length     int
}

// O(1)
func (s *SetList[T, L]) HasDuplicate(x T) bool {
	_, found := s.duplicates[x]
	return found
}

func (s *SetList[T, L]) Push(x T) {
	if !s.HasDuplicate(x) {
		s.basicList.Push(x)
		s.duplicates[x] = struct{}{}
		s.length++
	}
}

func (s *SetList[T, L]) Pop() *T {
	toPop := s.basicList.Pop()
	s.length--
	return toPop
}

func (s *SetList[T, L]) Len() int {
	return s.length
}

func (s *SetList[T, L]) IsEmpty() bool {
	return s.length == 0
}

func WrapSetList[T comparable](underlyingList BasicList[T]) *SetList[T, BasicList[T]] {
	return &SetList[T, BasicList[T]]{
		basicList:  underlyingList,
		duplicates: make(map[T]struct{}),
	}
}
