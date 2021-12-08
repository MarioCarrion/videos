# Learning Go - Benchmarking

## Important arguments

* `-benchtime=N` it accepts time or amount, for example `20s` (20 seconds) or `20x` (20 times).
* `-count=N` indicates how many times the benchmarks should be run.
* `-benchmem` includes memory allocation statistics in the result.

## Benchstat

Benchstat computes and compares statistics about benchmarks.

```
go install golang.org/x/perf/cmd/benchstat@latest
```

For example, running a benchmark using a:

```
go test -count=10 -bench=BenchmarkConcat | tee old.txt
```

Then modifying the original code with the new implementation and running:

```
go test -count=10 -bench=BenchmarkConcat | tee new.txt
```

Finally, you can compare both using `benchstat`:

```
benchstat old.txt new.txt
```
