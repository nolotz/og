package ptrval

// Value returns the value of pointer p.
// If p is nil, it returns a zero value of T.
func Value[T any](p *T) T {
	if p != nil {
		return *p
	}

	return Zero[T]()
}

// ValueOK returns the value of pointer p and ok set to true.
// If p is nil, it returns a zero value of T and ok is false.
func ValueOK[T any](p *T) (T, bool) {
	if p != nil {
		return *p, true
	}

	return Zero[T](), false
}

// Pointer returns a pointer to v.
func Pointer[T any](v T) *T {
	return &v
}

// Zero returns a new zero value of T.
func Zero[T any]() (v T) {
	return
}
