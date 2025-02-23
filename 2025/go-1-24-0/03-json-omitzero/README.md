# `encoding/json` omitzero

Excerpt from [official release notes](https://go.dev/doc/go1.24#encodingjsonpkgencodingjson):

> When marshaling, a struct field with the new `omitzero` option in the struct
> field tag will be omitted if its value is zero. If the field type has an
> `IsZero() bool` method, that will be used to determine whether the value is
> zero. Otherwise, the value is zero if it is the zero value for its type. The
> `omitzero` field tag is clearer and less error-prone than `omitempty` when
> the intent is to omit zero values. In particular, unlike `omitempty`,
> `omitzero` omits zero-valued `time.Time` values, which is a common source of
> friction.
>
> If both `omitempty` and `omitzero` are specified, the field will be omitted
> if the value is either empty or zero (or both).
