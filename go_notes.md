# Go Thumb Rules / Patterns — Revision Notes

## 1. Constructor pattern: `NewXxx(...)`

When a struct's fields need setup beyond Go's zero values (e.g. allocating a
slice with `make`, computing a derived field), write a plain function
`NewXxx(...) *Xxx` that builds and returns it — instead of writing the struct
literal by hand everywhere it's needed.

```go
func NewSparseSet(n int) *SparseSet {
    return &SparseSet{
        dense:  make([]int, n),
        sparse: make([]int, n),
        n:      n,
    }
}
```

- If a struct's fields are fine with Go's zero values (`0`, `""`, `nil`, `false`),
  you don't need a constructor — `SparseSet{}` directly is enough.
- Constructor centralizes setup in one place: add/change a field once, every
  caller stays correct automatically.
- `NewXxx` has **no receiver** — there's no existing instance yet, it's creating one.

## 2. Pointer receiver vs value receiver

```go
func (c Counter) IncrementWrong()  { c.value++ }  // modifies a COPY — no-op on the original
func (c *Counter) IncrementRight() { c.value++ }  // modifies the ORIGINAL
```

- Value receiver `(c Counter)` → Go passes a copy. Changes inside the method vanish.
- Pointer receiver `(c *Counter)` → Go passes the address. Changes stick.
- Rule of thumb: if the method needs to **mutate** the struct's fields, use a
  pointer receiver. Read-only methods can use either, but pointer receivers
  are the common default once *any* method on the type needs one (consistency).
