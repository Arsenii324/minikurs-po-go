package main

import (
	"fmt"
	"strings"
)

func f1() {
	// Hello world
	fmt.Println("Hello World")
}

func f2(a int, b int) int {
	// a + b
	return a + b
}

func f3(a int) string {
	// Even/Odd
	if a%2 != 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func f4(a int, b int, c int) int {
	// max of 3
	// MAJ3 vibes
	if a >= b {
		if a >= c {
			return a
		}
		return c
	}
	if b >= c {
		return b
	}
	return c
}

func f5(a int) int64 {
	// factorial
	if a <= 0 {
		return 1
	}
	return int64(a) * f5(a-1)
}

func f6(a rune) bool {
	// is vowel
	vowel := "уеыаоэёяиюeyuioa"
	vowel += strings.ToUpper(vowel)
	for _, c := range vowel {
		if a == c {
			return true
		}
	}
	return false
}

func Use(vals ...interface{}) { _ = vals }

func main() {
	Use(f1, f2, f3, f4, f5, f6)
	for true {
		var a rune
		_, _ = fmt.Scanln(&a)
		fmt.Println(a)
		fmt.Println(f6(a))
	}
}
