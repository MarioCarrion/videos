# Wrapping multiple errors

Excerpt from official release notes:

> Go 1.20 expands support for error wrapping to permit an error to wrap multiple other errors.

> An error e can wrap more than one error by providing an Unwrap method that returns a []error.

> The errors.Is and errors.As functions have been updated to inspect multiply wrapped errors.

> The fmt.Errorf function now supports multiple occurrences of the %w format verb, which will cause it to return an error that wraps all of those error operands.

> The new function errors.Join returns an error wrapping a list of errors.
