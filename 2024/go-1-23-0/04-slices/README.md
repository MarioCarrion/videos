# `slices` package changes

The `slices` package adds **9 functions** that support iterators:

* [All](https://go.dev/pkg/slices#All) returns an iterator over slice indexes and values.
* [AppendSeq](https://go.dev/pkg/slices#AppendSeq) appends values from an iterator to an existing slice.
* [Backward](https://go.dev/pkg/slices#Backward) returns an iterator that loops over a slice backward.
* [Chunk](https://go.dev/pkg/slices#Chunk) returns an iterator over consecutive sub-slices of up to n elements of a slice.
* [Collect](https://go.dev/pkg/slices#Collect) collects values from an iterator into a new slice.
* [Sorted](https://pkg.go.dev/slices#Sorted) Sorted collects values from seq into a new slice, sorts the slice, and returns it.
* [SortedFunc](https://go.dev/pkg/slices#SortedFunc) is like Sorted but with a comparison function.
* [SortedStableFunc](https://go.dev/pkg/slices#SortedStableFunc) is like SortFunc but uses a stable sort algorithm.
* [Sorted](https://go.dev/pkg/slices#Sorted) collects values from an iterator into a new slice, and then sorts the slice.
* [Values](https://go.dev/pkg/slices#Values) returns an iterator over slice elements.

[pkg.go.dev/slices](https://pkg.go.dev/slices)
