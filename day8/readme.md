# Day 8

## Number-centric thinking

- 0, 6, and 9 all use 6 segments
  - 6 is the only one that doesn't share both of 1's (or 7's) segments - and can tell us which of 1's segments is which
  - 2, 3, and 5 all use the middle segment, so zero will be the only one that doesn't
- 2, 3, and 5 all use 5 segments
  - 3 is the only one that shares both of 1's segments
  - once we know from comparison of 1 and 6 which of 1's segments is which we can differentiate between 2 and 5
- 1, 4, 7, and 8 use unique number of segments
  - comparison of 1 and 7 can tell us which is the top bar
  - identifying 8 doesn't really help since it doesn't lead to any new information

## Segment-centric thinking

- right-top
  - missing from 6 which is the only 6-segment number that doesn't use both of 1's segments
- right-bottom
  - comparison of 1 and 6
- middle
  - one of the 2 segments 3 adds to 7 (and 3 is the only 5-segment to use both of 1's segments)
  - one of the 2 segments 4 adds to 1
- top
  - the only segment 7 adds to 1
- bottom
- left-top
  - one of the 2 segments 4 adds to 1
- left-bottom

## Order of operations

1. find the four easy numbers (1, 4, 7, 8)
2. this leads to the top segment being identified (the only difference b/w 1 and 7)
3. find 3 - it is the only 5-segment number that shares both of 1's segments
4. find 6 - it is the only 6-segment number that doesn't share both of 1's segments
5. find right-top and right-bottom by figuring out which of 1's segments is shared with 6
6. find middle - it is the only segment that both 3 and 4 add to 1
7. find bottom - it is the only additional segment added by 3 that it doesn't share with 4 or 7
8. find left-top - it is the only additional segment aded by 4 that isn't the middle
9. find left-bottom - it is the only segment left.
10. Now that we have all segments identified, just need to build a tool that can

## Data structures and functions

- probably need an enum for positional segments (top, upper-left, lower-right, bottom, etc) - maybe one that can be bitwise-combined
- probably want a map from known number to segments that compose that number (really only need 1, 4, 7, 3, and 6)
- need a map from letter to the known positional segment that we know that letter represents
- need an algorithm to remove certain elements from an array and easy ways to test if an element exists in an array
- need a function that can translate (map) from a, b, c into known-position-segments
- need a function that can translate from a set of position-segments to an integer
- all of this would probably be easier in a higher-level language than golang - at least to hammer out an algo

