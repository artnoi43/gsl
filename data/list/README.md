# Package `list`
Package `list` provides building blocks for other list-like data structures like stacks and queues.

## `BasicList[T]`
The code is built around a base interface `BasicList[T]` (like `Stack[T]`, `PriorityQueue[T]`, and `Queue[T]`). All these concrete *list* types are implemented with simplicity in mind, and all have `Push`, `Pop`, `Len`, and `IsEmpty` methods.

## `SetList[T]`
In addition to `BasicList[T]`, this package also provides `SetList[T]`, which is `BasicList[T]` with 1-2 extra methods. This interface is mainly implemented by `*SetListWrapper` and shitty struct `SetListImpl[T]`.

> I don't even know in what scenario will I be forced to use `SetListImpl[T]`. I think other list types wrapped in `*SetListWrapper[T, L]` is much more useful and less confusing to use. I may remove `SetListImpl[T]` because I fucking hate it.

A `SetList[T]` is, as the name suggests, a list of T with set functionality. This is not to be confused with a more traditional `Set[T]` in package `container`, because `SetList[T]` is not a container set where you can freely access data with an index. Instead, `SetListImpl[T]` is just a very small extension of `BasicList[T]`.

## Wrapper types
In addition for the basic list types, this package also provides wrappers for any `BasicList[T]` instances.
### `SafeList[T, L]` wrapper
An example of these wrappers is `SafeList[T any, L BasicList[T]]`, which, as its type name suggests, wraps any `BasicList[T]` into a struct with mutex for safe concurrent operations.
### `SetList[T, L]` wrapper
Another example is `SetList[T comparable, BasicList[T]`], which wraps any `BasicList` into a *set*-like data structures where no duplicates are allowed. It does this by embedding a `BasicList[T]` into a struct with hash map of items and list length, to help with random access time when determining duplicates.