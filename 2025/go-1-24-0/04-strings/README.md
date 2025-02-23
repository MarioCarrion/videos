# `strings` package updates

Excerpt from [official release notes](https://go.dev/doc/go1.24#stringspkgstrings):

> The `strings` package adds several functions that work with iterators:
>
> `Lines` returns an iterator over the newline-terminated lines in a string.
> `SplitSeq` returns an iterator over all substrings of a string split around a separator.
> `SplitAfterSeq` returns an iterator over substrings of a string split after each instance of a separator.
> `FieldsSeq` returns an iterator over substrings of a string split around runs of whitespace characters, as defined by `unicode.IsSpace`.
> `FieldsFuncSeq` returns an iterator over substrings of a string split around runs of Unicode code points satisfying a predicate.
