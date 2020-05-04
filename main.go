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
	kataCast()
	sliceDemo()
	println(名前付き戻り値(true))
	println(名前付き戻り値(false))

}

func evenOdd() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
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

func kataCast() {
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		println("good")
	}
}

func sliceDemo() {
	// n1 := 19
	// n2 := 86
	// n3 := 1
	// n4 := 12
	// sum := n1 + n2 + n3 + n4

	num := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	println(sum)
}

// Score はゲームスコア用の型です
type Score struct {
	User       string
	GameNumber int
	Point      int
}

func 名前付き戻り値(flag bool) (x int) {
	if flag {
		print("型のデフォルト値を返します")
		return
	} else {
		print("処理結果を返します")
		x = 1 + 1
		return
	}
}
