# Range over func

Excerpt from official release notes:

> The "range" clause in a "for-range" loop now accepts iterator functions of the following types
>
> func(func() bool)
> func(func(K) bool)
> func(func(K, V) bool)
>
> as range expressions.

[Go Wiki: Rangefunc Experiment](https://go.dev/wiki/RangefuncExperiment)
