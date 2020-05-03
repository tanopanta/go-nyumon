package main

import (
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var msg string = "Hello"
	msg2 := "World"
	println(msg, msg2)

	runeA := 'A'
	runeB := 'B'
	println(runeA, runeB)

	const kata string = "katakata"
	println(kata)

	const (
		a = iota
		b
		c
	)
	println(a, b, c)

	// evenOdd()
	omikuji()
}

func evenOdd() {
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			println("偶数 - ", i)
		} else {
			println("奇数 - ", i)
		}
	}
}

func omikuji() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6) + 1
	print(strconv.Itoa(n) + ":")
	switch n {
	case 1:
		println("凶")
	case 2, 3:
		println("吉")
	case 4, 5:
		println("中吉")
	case 6:
		println("大吉") 
	}
}