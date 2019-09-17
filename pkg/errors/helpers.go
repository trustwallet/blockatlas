package errors

// Is reports whether err is an *Error of the given Type.
// If err is nil then Is returns false.
func Is(kind Type, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.Type != TypeNone {
		return e.Type == kind
	}
	if e.Err != nil {
		return Is(kind, e.Err)
	}
	return false
}

func Equal(err1, err2 error) bool {
	e1, ok := err1.(*Error)
	if !ok {
		return false
	}
	e2, ok := err2.(*Error)
	if !ok {
		return false
	}
	if e1.Err != nil && e2.Err != e1.Err {
		return false
	}
	if e1.meta != nil && e2.meta != e1.meta {
		return false
	}
	if e1.Type != TypeNone && e2.Type != e1.Type {
		return false
	}
	if e1.Err != nil {
		if _, ok := e1.Err.(*Error); ok {
			return Equal(e1.Err, e2.Err)
		}
		if e2.Err == nil || e2.Err.Error() != e1.Err.Error() {
			return false
		}
	}
	return true
}
