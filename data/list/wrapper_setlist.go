package list

// SetListWrapper wraps BasicList[T] into field `basicList`,
// and maintains a hash map of seen items and the basicList length
// so that determining duplicates and getting the length take O(1) time.
type SetListWrapper[T comparable, L BasicList[T]] struct {
	basicList  L
	duplicates map[T]struct{}
	length     int
}

// O(1)
func (s *SetListWrapper[T, L]) HasDuplicate(x T) bool {
	_, found := s.duplicates[x]
	return found
}

func (s *SetListWrapper[T, L]) Push(x T) {
	if !s.HasDuplicate(x) {
		s.basicList.Push(x)
		s.duplicates[x] = struct{}{}
		s.length++
	}
}

func (s *SetListWrapper[T, L]) PushSlice(values []T) {
	for _, elem := range values {
		s.Push(elem)
	}
}

func (s *SetListWrapper[T, L]) Pop() *T {
	toPop := s.basicList.Pop()
	s.length--
	return toPop
}

func (s *SetListWrapper[T, L]) Len() int {
	return s.length
}

func (s *SetListWrapper[T, L]) IsEmpty() bool {
	return s.length == 0
}

func WrapSetList[T comparable](underlyingList BasicList[T]) *SetListWrapper[T, BasicList[T]] {
	return &SetListWrapper[T, BasicList[T]]{
		basicList:  underlyingList,
		duplicates: make(map[T]struct{}),
	}
}
