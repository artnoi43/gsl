package list

import "testing"

func testSetListWrapper(t *testing.T) {
	testSetListQueue(t, []float64{1, 2, 4, 2, 4, 1, 3, 5})
	testSetListQueue(t, []string{"foo", "bar", "baz", "bar", "foom"})
	testSetListStack(t, []float64{1, 2, 4, 2, 4, 1, 3, 5})
	testSetListStack(t, []string{"foo", "bar", "baz", "bar", "foom"})
}

func testSetListStack[T comparable](t *testing.T, data []T) {
	set := NewSet(data)

	// Wrap stack with SetListWrapper and then wrap that shit with SafeListWrapper
	stack := NewStack[T]()                    // *StackImpl[T]
	setStack := WrapSetList[T](stack)         // *SetList[T, Stack[T]]
	safeSetStack := WrapSafeList[T](setStack) // *SafeList[T, *SetListWrapper[T]]

	// Wrap anotherStack with SafeListWrapper and then wrap that shit with SetListWrapper
	anotherStack := NewStack[T]()              // *Stack[T]
	safeStack := WrapSafeList[T](anotherStack) // *SafeList[T, Stack[T]]
	setSafeStack := WrapSetList[T](safeStack)  // *SetList[T, *SafeList[T, Stack[T]]]

	// All 6 should implement BasicList[T] and Stack[T]
	_ = []BasicList[T]{
		set,
		stack, setStack, safeSetStack,
		anotherStack, safeStack, setSafeStack,
	}

	// If wrapped lastly be SafeList[T, BasicList[T]] (like safeSetStack), then it's obviously NOT Set[T].
	_ = []Set[T]{
		set,
		setStack,
		// safeSetStack, // Compile error!
		setSafeStack,
	}

	testSets := []Set[T]{set, setStack, setSafeStack}
	for _, set := range testSets {
		testSetPushAndPop[T](t, set, data)
	}
}

func testSetListQueue[T comparable](t *testing.T, data []T) {
	set := NewSet(data)

	queue := NewQueue[T]() // *QueueImpl[T]
	setQueue := WrapSetList[T](queue)
	safeSetQueue := WrapSafeList[T](setQueue)

	anotherQueue := NewQueue[T]()
	safeQueue := WrapSafeList[T](queue)
	setSafeQueue := WrapSetList[T](safeQueue)

	// All 6 should implement BasicList[T] and Queue[T]
	_ = []BasicList[T]{
		set,
		queue, setQueue, safeSetQueue,
		anotherQueue, safeQueue, setSafeQueue,
	}

	// If wrapped lastly be SafeList[T, BasicList[T]] (like safeSetQueue), then it's obviously NOT Set[T].
	_ = []Set[T]{
		set,
		setQueue,
		// safeSetQueue, // Compile error!
		setSafeQueue,
	}

	testSets := []Set[T]{set, setQueue, setSafeQueue}
	for _, set := range testSets {
		testSetPushAndPop[T](t, set, data)
	}
}

func testSetPushAndPop[T comparable](t *testing.T, setList BasicList[T], data []T) {
	for _, item := range data {
		setList.Push(item)
	}

	seenCounts := make(map[T]int)
	for !setList.IsEmpty() {
		popped := setList.Pop()
		if popped == nil {
			t.Fatal("popped nil - should not happen")
		}
		value := *popped
		seenCounts[value]++
	}

	for value, count := range seenCounts {
		if count != 1 {
			t.Fatalf("has duplicates for key %v in set\n", value)
		}
	}
}
