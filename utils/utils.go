package utils

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func NullIfEmpty(s string) string {
	if s == "" {
		return "null"
	}
	return s
}
