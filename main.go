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

}