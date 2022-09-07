package list

import "testing"

func testSetList(t *testing.T) {
	testSetListQueue(t, []float64{1, 2, 4, 2, 4, 1, 3, 5})
	testSetListQueue(t, []string{"foo", "bar", "baz", "bar", "foom"})
	testSetListStack(t, []float64{1, 2, 4, 2, 4, 1, 3, 5})
	testSetListStack(t, []string{"foo", "bar", "baz", "bar", "foom"})
}

func testSetListStack[T comparable](t *testing.T, data []T) {
	set := NewSet(data)

	stack := NewStack[T]()
	takeList[T](stack)
	setStack := WrapSetList[T](stack)
	takeList[T](setStack)
	takeSet[T](setStack)
	safeSetStack := WrapSafeList[T](setStack)
	takeList[T](safeSetStack)

	anotherStack := NewStack[T]()
	safeStack := WrapSafeList[T](anotherStack)
	takeList[T](safeStack)
	setSafeStack := WrapSetList[T](safeStack)
	takeSet[T](setSafeStack)
	// takeSet[T](safeSetStack) // Compile error! WrapSafeList[T] retutns *SafeList[T any, L BasicList[T]]

	testSets := []Set[T]{set, setStack, setSafeStack}
	for _, set := range testSets {
		testSetPushAndPop[T](t, set, data)
	}
}

func testSetListQueue[T comparable](t *testing.T, data []T) {
	set := NewSet(data)

	queue := NewQueue[T]()
	setQueue := WrapSetList[T](queue)
	takeSet[T](setQueue)
	safeSetQueue := WrapSafeList[T](setQueue)
	testSetPushAndPop[T](t, safeSetQueue, data)

	testSetPushAndPop[T](t, setQueue, data)
	testSetPushAndPop[T](t, safeSetQueue, data)

	testSets := []Set[T]{set, setQueue}
	for _, set := range testSets {
		testSetPushAndPop[T](t, set, data)
	}
}

func takeSet[T comparable](s Set[T]) {}

func takeList[T comparable](s BasicList[T]) {}

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
