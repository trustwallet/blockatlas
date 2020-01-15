package blockatlas

func GetValidParameter(first, second string) string {
	if len(first) > 0 {
		return first
	}
	return second
}
