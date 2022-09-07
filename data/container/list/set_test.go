package list

import "testing"

func testSetList(t *testing.T) {
	testSetListQueue(t, []float64{1, 2, 4, 2, 4, 1, 3, 5})
	testSetListQueue(t, []string{"foo", "bar", "baz", "bar", "foom"})
	testSetListStack(t, []float64{1, 2, 4, 2, 4, 1, 3, 5})
	testSetListStack(t, []string{"foo", "bar", "baz", "bar", "foom"})
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

func testSetListStack[T comparable](t *testing.T, data []T) {
	stack := NewStack[T]()
	setStack := WrapSetList[T](stack)
	safeSetStack := WrapSafeList[T](setStack)

	testSetPushAndPop[T](t, setStack, data)
	testSetPushAndPop[T](t, safeSetStack, data)
}

func testSetListQueue[T comparable](t *testing.T, data []T) {
	queue := NewQueue[T]()
	setQueue := WrapSetList[T](queue)
	safeSetQueue := WrapSafeList[T](setQueue)

	testSetPushAndPop[T](t, setQueue, data)
	testSetPushAndPop[T](t, safeSetQueue, data)
}
