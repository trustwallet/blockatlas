package errors

// Is reports whether err is an *Error of the given Type.
// If err is nil then Is returns false.
func Is(err error, t Type) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.Type != TypeNone {
		return e.Type == t
	}
	if e.Err != nil {
		return Is(e.Err, t)
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

func appendMap(root map[string]interface{}, tmp map[string]interface{}) {
	for k, v := range tmp {
		root[k] = v
	}
}
