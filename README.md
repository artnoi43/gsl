# gsl: Go Soylib

gsl provides frequently used functionalities that I use across my projects.

It is my personal learning project.

Since Go is focused on simplicity, some *battery-included* kind of built-ins are not included in Go (and that's a good thing).
This led programmers to manually implement these basic solutions, and this led to a lot of duplicate code, oversighted bugs, etc.

This is where gsl comes in - to standardize and store frequently used generic business logic for soy and non-soy devs alike.

> Since 2023, Go 1.21 was released with many "battery-included" standard packages, such as `"slices"`,
> that does many of the things gsl does, so you might want to explore those packages first before gsl.
>
> Go also has released many [official extra packages](https://pkg.go.dev/golang.org/x).

- `gsl` - basic utilities for everyday use

- `soyutils` - utilities for soy devs

- `http` - providing wrapper for Go's `http` package

- `data` - basic data structure library. Highly unstable,
  because this is where I learn to code data structures

- `concurrent` - simple stupid concurrency helper functions

- `sqlquery` - interface-based SQL query builder to complement [squirrel](https://github.com/Masterminds/squirrel)
  on features such as INSERT ALL or upserts
