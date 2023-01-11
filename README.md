# gsl: Go Soylib - where I store my boilerplate code

gsl provides frequently used functionalities that I use across my projects.

Since Go is focused on simplicity, some *battery-included* kind of built-ins are not included in Go (and that's a good thing).
This led programmers to manually implement these basic solutions, and this led to a lot of duplicate code, oversighted bugs, etc.

This is where gsl comes in - to standardize and store frequently used generic business logic for soy and non-soy devs alike.

- `gsl` - basic utilities for everyday use

- `soyutils` - utilities for soy devs

- `http` - providing wrapper for Go's `http` package

- `data` - Basic data structure library. Highly unstable, because this is where I learn to code data structures
