package temp

import (
	"fmt"

	"example.com/temp/panic"
)

var shouldBeZero = 0

func t(a int, b int) bool {
	return a+b == 2
}

func foo() {
	panic.DoPanic(true)
	shouldBeZero = 1
	if t(1, 1) {
		shouldBeZero = 2
	}
	fmt.Println(shouldBeZero)
}
