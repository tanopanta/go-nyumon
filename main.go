package main

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

	evenOdd()

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