package panic

func DoPanic(b bool) {
	if b {
		panic(1)
	}
}
