package main

func rotateRunes(s string, shift int) string {
	r := []rune(s)
	n := len(r)
	if n == 0 {
		return ""
	}
	k := shift % n
	if k < 0 {
		k += n
	}

	return string(append(append(make([]rune, 0, n), r[k:]...), r[:k]...))
}
