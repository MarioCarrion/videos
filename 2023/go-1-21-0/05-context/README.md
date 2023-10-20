# New `context` functions

Excerpt from official release notes:

> The new `WithoutCancel` function returns a copy of a context that is not canceled when the original context is canceled.

> The new `WithDeadlineCause` and `WithTimeoutCause` functions provide a way to set a context cancellation cause when a deadline or timer expires. The cause may be retrieved with the `Cause` function.

> The new `AfterFunc` function registers a function to run after a context has been cancelled.
