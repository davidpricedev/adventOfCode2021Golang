# adventOfCode2021

This repository contains my solutions for many of the puzzles in the [advent of code](https://adventofcode.com) 2021.

As I usually do, I've decided to learn a new language while solving these puzzles. This year the new language is ruby.

## My reaction to go

Go in general seems to be 90% C with garbage collection with maybe 10% python sprinkled in
Generally a very verbose, clunky language.
Strong types combined with no type coercion into strings and no easy ways to manipulate lists means you waste a lot of real estate on type conversions.

Go's rigidity is in many ways perfectly antagonistic to all things FP.
It isn't so much that go makes it hard to do FP, go makes it entirely impossible - largely because it doesn't support flexible generics.
This forces you to think procedurally instead of functionally.

### Things I don't like

- imports and conventions for package naming are very verbose and not immediately obvious
- verbosity - there is nothing so compact as python's comprehensions
- fundamentally lacking the expected basic list manipulation functions such as map, filter, reduce
- tabs. Really? tabs in a language written long after tabs should have been dead and gone, they chose tabs
- lacking negative indexers into arrays - they borrowed the array slicing syntax from python but not the time-saving negative indexers?
- missing modern string interpolation - and the automated to-string type coercion that would make it easy to do
- why do we need a different syntax for inferred types? F# and other languages manage just fine without
- pointers - all the other "modern" languages handle pass-by-ref automatically why did we have to continue with C's pointers?

### Things I do like

- capital vs lowercase is an interesting way of handling internal vs exported things - the sort of ethos I'd expect from Python
- core functions seem to generally be immmutable
- supports type aliasing
- automatic initialization of everything with a zero-value, this makes a few things surprisingly nice
