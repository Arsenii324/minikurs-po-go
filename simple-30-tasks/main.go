package main

import (
	"cmp"
	"fmt"
	"hash/fnv"
	"sort"
	"strconv"
	"strings"
)

func Use(vals ...interface{}) { _ = vals }

func f1() {
	// hello world
	fmt.Println("Hello World")
}

func f2(a int, b int) int {
	// a + b
	return a + b
}

func f3() {
	// even or odd
	var a int
	_, _ = fmt.Scan(&a)
	if a%2 != 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func f4(a int, b int, c int) int {
	// max of 3
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func f5() {
	// factorial
	//
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(f5Rec1(a))
}

func f5Rec1(a int) int64 {
	// factorial helper function
	// hopefully int
	if a <= 0 {
		return 1
	}
	return int64(a) * f5Rec1(a-1)
}

func f6(a string) bool {
	// is vowel
	vowel := "aeuioy"
	vowel += strings.ToUpper(vowel)
	return len(a) == 1 && strings.Contains(vowel, a)
}

func f7() {
	// all primes under a
	var a int
	_, _ = fmt.Scan(&a)
	if a <= 1 {
		fmt.Println()
		return
	}
	var tb = make([]bool, a+1)

	for i := range tb {
		tb[i] = true
	}
	tb[0], tb[1] = false, false

	for i := 2; i <= a; i++ {
		if tb[i] {
			for j := i; j <= a; j += i {
				tb[j] = false
			}
			fmt.Print(i, ' ')
		}
	}
	fmt.Println()
}

func f8(a string) string {
	// reverse string
	var b []rune
	p := 0
	for _, r := range a {
		b = append(b, r)
		p += 1
	}
	b = b[:p]
	// Reverse
	for i := 0; i < p/2; i++ {
		b[i], b[p-1-i] = b[p-1-i], b[i]
	}

	return string(b)
}

func f9(a []int) int {
	// sum array
	b := 0
	for _, x := range a {
		b += x
	}
	return b
}

type Rectangle struct {
	Width, Height int
}

func (a *Rectangle) Surface() int {
	// rectangle
	return a.Width * a.Height
}

// ===

func f11(a float32) float32 {
	// celsium to fahrenheit
	return a*1.8 + 32
}

func f12() {
	// count from N to 1
	var N int
	_, _ = fmt.Scan(&N)
	for i := N; i >= 1; i-- {
		fmt.Println(i)
	}
}

func f13(a string) int {
	// string size
	size := 0
	for range a {
		size++
	}
	return size
}

func f14[T cmp.Ordered](a []T, elem T) bool {
	// check if element in array
	for _, v := range a {
		if elem == v {
			return true
		}
	}
	return false
}

func f15(a []int) float32 {
	// find average
	s := 0
	for _, v := range a {
		s += v
	}
	return float32(s) / float32(len(a))
}

func f16() {
	// multiplication table
	var N int
	if _, err := fmt.Scan(&N); err != nil {
		panic(err)
	}
	if N <= 0 {
		fmt.Println("empty")
		return
	}
	lg := len(strconv.Itoa(N * N))
	fmt.Print(strings.Repeat(" ", lg))
	for i := 1; i <= N; i++ {
		fmt.Print(strings.Repeat(" ", 1+lg-len(strconv.Itoa(i))))
		fmt.Print(i)
	}
	fmt.Println()
	for j := 1; j <= N; j++ {
		for i := 1; i <= N; i++ {
			fmt.Print(strings.Repeat(" ", 1+lg-len(strconv.Itoa(i*j))))
			fmt.Print(i * j)
		}
		fmt.Println()
	}

}

func f17(a string) bool {
	// is palindrome
	for i := range len(a) / 2 {
		if a[len(a)-i-1] != a[i] {
			return false
		}
	}
	return true
}

func f18(a []int) (int, int) {
	// find min, max
	if len(a) == 0 {
		panic("Empty array")
	}
	minmax := [2]int{a[0], a[0]}
	for _, v := range a {
		if v < minmax[0] {
			minmax[0] = v
		}
		if v > minmax[1] {
			minmax[1] = v
		}
	}
	return minmax[0], minmax[1]
}

func f19(a []int, i int) []int {
	// remove ith element
	if (i < 0) || (i >= len(a)) {
		panic("Index out of range")
	}
	return append(a[:i], a[i+1:]...)
}

func f20(a []int, elem int) int {
	// find index
	for i, v := range a {
		if elem == v {
			return i
		}
	}
	return -1
}

func f21(a []int) []int {
	// deduplicate
	var m map[int]bool
	var b []int
	for _, v := range a {
		if !m[v] {
			b = append(b, v)
			m[v] = true
		}
	}
	return b
}

func f22(a []int) []int {
	// bubble sort
	for i := range len(a) {
		for j := i - 1; j >= 0; j-- {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func f23(N int) []int {
	// fibonacci sequence
	if N <= 0 {
		return make([]int, 0)
	}
	if N == 1 {
		a := []int{1}
		return a
	}
	a := make([]int, N)
	a[0] = 1
	a[1] = 1
	for i := 2; i < N; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a
}

func f24[T cmp.Ordered](a []T, elem T) int {
	// count occurrences
	cnt := 0
	for _, v := range a {
		if elem == v {
			cnt += 1
		}
	}
	return cnt
}

func f25(a, b []int) []int {
	// array intersection
	aC := make([]int, len(a))
	copy(a, aC)
	sort.Ints(aC)
	bC := make([]int, len(b))
	copy(b, bC)
	sort.Ints(bC)
	aCIter, bCIter := 0, 0
	var ans []int
	for aCIter < len(a) && bCIter < len(b) {
		if aC[aCIter] == bC[bCIter] {
			ans = append(ans, aC[aCIter])
			aCIter++
			bCIter++
		} else {
			if aC[aCIter] > bC[bCIter] {
				bCIter++
			} else {
				aCIter++
			}
		}
	}
	return ans
}

func f26(a, b string) bool {
	// is anagram
	if len(a) != len(b) {
		return false
	}
	aVec, bVec := []rune(a), []rune(b)
	sort.Slice(aVec, func(i, j int) bool {
		return aVec[i] < aVec[j]
	})
	sort.Slice(bVec, func(i, j int) bool {
		return bVec[i] < bVec[j]
	})
	for i, v := range aVec {
		if v != bVec[i] {
			return false
		}
	}
	return true
}

func f27(a, b []int) []int {
	// merge sorted arrays
	// hopefully ascending order
	i, j := 0, 0
	c := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	return append(append(c, a[i:]...), b[j:]...)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

type MyMap[T any] struct {
	KVPairs [][]KVPair[T]
	count   int
}

type KVPair[T any] struct {
	key   string
	value T
}

func (a *MyMap[T]) get(key string) (T, bool) {
	h := hash(key) % uint32(len(a.KVPairs))
	for _, p := range a.KVPairs[h] {
		if p.key == key {
			return p.value, true
		}
	}
	var t T
	return t, false
}

func (a *MyMap[T]) rehash(size int) {
	NewKVPairs := make([][]KVPair[T], size)
	for _, v := range a.KVPairs {
		for _, p := range v {
			h := hash(p.key)
			NewKVPairs[h%uint32(len(NewKVPairs))] = append(NewKVPairs[h], p)
		}
	}
	a.KVPairs = NewKVPairs
}

func (a *MyMap[T]) set(key string, value T) {
	h := hash(key) % uint32(len(a.KVPairs))
	for i, p := range a.KVPairs[h] {
		if p.key == key {
			a.KVPairs[h][i].value = value
			return
		}
	}
	a.KVPairs[h] = append(a.KVPairs[h], KVPair[T]{key: key, value: value})
	a.count += 1
	if a.count*2 >= len(a.KVPairs) {
		a.rehash(a.count * 2)
	}
}

func NewMyMap[T any]() MyMap[T] {
	a := MyMap[T]{}
	a.rehash(32)
	return a
}

func f29(a []int, value int) int {
	// binary search
	l, r := -1, len(a)
	for l+1 < r {
		m := (l + r) / 2
		if a[m] > value {
			r = m
		} else {
			l = m
		}
	}
	if l == -1 || a[l] != value {
		return -1
	}
	return l
}

type MyQueue[T any] struct {
	stack1, stack2 []T
}

func (a *MyQueue[T]) Push(v T) {
	a.stack1 = append(a.stack1, v)
}

func (a *MyQueue[T]) Pop() (T, bool) {
	if len(a.stack2) == 0 {
		for len(a.stack1) != 0 {
			a.stack2 = append(a.stack2, a.stack1[len(a.stack1)-1])
			a.stack1 = a.stack1[:len(a.stack1)-1]
		}
	}
	if len(a.stack2) == 0 {
		var t T
		return t, false
	}
	t := a.stack2[len(a.stack2)-1]
	a.stack2 = a.stack2[:len(a.stack2)-1]
	return t, true
}

func NewMyQueue[T any]() MyQueue[T] {
	return MyQueue[T]{}
}

func main() {
	Use(f1, f2, f3, f4, f5, f6, f7, f8, f9, Rectangle{})
	Use(f11, f12, f13, f14[int8], f15, f16, f17, f18, f19, f20)
	Use(f21, f22, f23, f24[int8], f25, f26, f27, NewMyMap[int8], f29, NewMyQueue[int8])
}
