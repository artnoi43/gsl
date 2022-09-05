# gsl: G Soylib - where I store my boilerplate code
gsl was conceived simply because I found myself writing too much duplicate boilerplate code.

Code like handling errors concurrently, or priority queues, are duplicated all over my projects. Even with Go generics (>=1.18), there're still duplicate code. This is bad, because I'll have to change code in many places. And I ususally change code, since I'm a beginner who only learned to code like < 2 years.

So this package is meant to fix that - by providing a single-source library for all my other code to import from. These are some subpackages, which may not be in sync with the current repo state:

1. `http` - providing wrapper for Go's `http` package

2. `strutils` - utilities for `~string` types

2. `mathutils` - basic math utilities

4. `data` - interesting data structures - is further divided into 2 subpackages:

- `data/list` - list data structures like stacks, queues, and priority queues

- `data/graph` - graph data structures