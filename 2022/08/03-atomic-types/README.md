# Core library

## New atomic types

Excerpt from official release notes:

> The sync/atomic package defines new atomic types Bool, Int32, Int64, Uint32, Uint64, Uintptr, and Pointer. These types hide the underlying values so that all accesses are forced to use the atomic APIs. Pointer also avoids the need to convert to unsafe.Pointer at call sites. Int64 and Uint64 are automatically aligned to 64-bit boundaries in structs and allocated data, even on 32-bit systems.
