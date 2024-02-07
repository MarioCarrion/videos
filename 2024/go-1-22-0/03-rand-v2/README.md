# New math/rand/v2 package

Excerpt from official release notes:

> Go 1.22 includes the first “v2” package in the standard library, math/rand/v2. The changes compared to math/rand are
> detailed in proposal #61716. The most important changes are:

> * The Read method, deprecated in math/rand, was not carried forward for math/rand/v2. (It remains available in
>   math/rand.) The vast majority of calls to Read should use crypto/rand’s Read instead. Otherwise a custom Read can be
>   constructed using the Uint64 method.

> * The global generator accessed by top-level functions is unconditionally randomly seeded. Because the API guarantees
>   no fixed sequence of results, optimizations like per-thread random generator states are now possible.

> * The Source interface now has a single Uint64 method; there is no Source64 interface.

> * Many methods now use faster algorithms that were not possible to adopt in math/rand because they changed the output
>   streams.

> * The Intn, Int31, Int31n, Int63, and Int64n top-level functions and methods from math/rand are spelled more
>   idiomatically in math/rand/v2: IntN, Int32, Int32N, Int64, and Int64N. There are also new top-level functions and
>   methods Uint32, Uint32N, Uint64, Uint64N, Uint, and UintN.

> * The new generic function N is like Int64N or Uint64N but works for any integer type. For example a random duration
>   from 0 up to 5 minutes is rand.N(5*time.Minute).

> * The Mitchell & Reeds LFSR generator provided by math/rand’s Source has been replaced by two more modern
>   pseudo-random generator sources: ChaCha8 PCG. ChaCha8 is a new, cryptographically strong random number generator
>   roughly similar to PCG in efficiency. ChaCha8 is the algorithm used for the top-level functions in math/rand/v2. As
>   of Go 1.22, math/rand's top-level functions (when not explicitly seeded) and the Go runtime also use ChaCha8 for
>   randomness.

> We plan to include an API migration tool in a future release, likely Go 1.23.
