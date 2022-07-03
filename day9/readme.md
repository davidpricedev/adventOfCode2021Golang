# Day 9

For part 2 I really want to use a set to de-duplicate the collection of points that will compose a basin.
Golang, of course, doesn't provide a set.
I was able to find a library with a nice implementation of a Set for go: https://github.com/deckarep/golang-set
And I learned how to import an external dependency.
First on the command line: `go get github.com/deckarep/golang-set/v2`, then add the import `mapset "github.com/deckarep/golang-set/v2"`.