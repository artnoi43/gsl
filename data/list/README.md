# Package `list`
Package `list` provides building blocks for other list-like data structures like stacks and queues.

## `BasicList[T]`
The code is built around a base interface `BasicBlock[T]` (like `Stack[T]`, `PriorityQueue[T]`, and `Queue[T]`). All these *list* types are implemented with simplicity in mind, and all have `Push`, `Pop`, `Len`, and `IsEmpty` methods.

## Wrapper types
In addition for the basic list types, this package also provides wrappers for any `BasicList[T]` instances.
### `SafeList[T, L]` wrapper
An example of these wrappers is `SafeList[T any, L BasicList[T]]`, which, as its type name suggests, wraps any `BasicList[T]` into a struct with mutex for safe concurrent operations.
### `SetList[T, L]` wrapper
Another example is `SetList[T comparable, BasicList[T]`], which wraps any `BasicList` into a *set*-like data structures where no duplicates are allowed. It does this by embedding a `BasicList[T]` into a struct with hash map of items and list length, to help with random access time when determining duplicates.