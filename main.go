package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/tanopanta/go-nyumon/greeting"
)

func main() {
	var msg string = "Hello"
	msg2 := "World"
	fmt.Println(msg, msg2)

	runeA := 'A'
	runeB := 'B'
	fmt.Println(runeA, runeB)

	const kata string = "katakata"
	fmt.Println(kata)

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	// evenOdd()
	omikuji()
	kataCast()
	sliceDemo()
	fmt.Println(名前付き戻り値(true))
	fmt.Println(名前付き戻り値(false))

	f := func() {
		fmt.Println("My name is 無名関数")
	}
	f()

	n, m := swap(10, 20)
	fmt.Println(n, m)
	swap2(&n, &m)
	fmt.Println(n, m)
	receiver()

	fmt.Println(greeting.Do())

	// fmt.Println(os.Args)

	interfaceSample()

	nn := GameID(222)
	s, error := toStringer(nn)
	if error != nil {
		fmt.Fprint(os.Stderr, "ERROR:", error)
	} else {
		fmt.Println(s)
	}

	canotCast := 222
	s, error = toStringer(canotCast)
	if error != nil {
		fmt.Fprint(os.Stderr, "ERROR:", error)
	} else {
		fmt.Println(s)
	}
}

func evenOdd() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("偶数 - ", i)
		} else {
			fmt.Println("奇数 - ", i)
		}
	}
}

func omikuji() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6) + 1
	print(strconv.Itoa(n) + ":")
	switch n {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}

func kataCast() {
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		fmt.Println("good")
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
	fmt.Println(sum)
}

// Score はゲームスコア用の型です
type Score struct {
	User       string
	GameNumber int
	Point      int
}

func 名前付き戻り値(flag bool) (x int) {
	if flag {
		fmt.Println("型のデフォルト値を返します")
		return
	} else {
		fmt.Println("処理結果を返します")
		x = 1 + 1
		return
	}
}

func swap(x int, y int) (int, int) {
	return y, x
}

func swap2(x *int, y *int) {
	x, y = y, x
}

// MyInt はintに機能を足すための型
type MyInt int

// Inc はインクリメンタル処理
func (n *MyInt) Inc() { *n++ }

func receiver() {
	var n MyInt
	fmt.Println(n)
	n.Inc()
	fmt.Println(n)
}

// Stringer は 文字列屋さんです
type Stringer interface {
	String() string
}

// GameID は ゲームIDです
type GameID int

// UserID は ゲームIDです
type UserID int

// AdminID は ゲームIDです
type AdminID int

func (g GameID) String() string {
	return fmt.Sprintf("g-%d", int(g))
}
func (u UserID) String() string {
	return fmt.Sprintf("u-%d", int(u))
}
func (a AdminID) String() string {
	return fmt.Sprintf("a-%d", int(a))
}

func interfaceSample() {
	var gameID Stringer = GameID(1)
	var userID Stringer = UserID(2)
	var adminID Stringer = AdminID(3)
	printStringer(gameID)
	printStringer(userID)
	printStringer(adminID)

}

func printStringer(s Stringer) {
	switch v := s.(type) {
	case UserID:
		fmt.Println("UserID:", v)
	case GameID:
		fmt.Println("GameID:", v)
	case AdminID:
		fmt.Println("AdminID:", v)
	}

	// fmt.Println(reflect.TypeOf(s).Name(), s.String())
}

// StringerError は Stringerのユーザー定義エラーです
type StringerError string

func (e StringerError) Error() string {
	return string(e)
}

func toStringer(v interface{}) (Stringer, error) {
	s, error := v.(Stringer)

	if !error {
		return nil, StringerError("CastError")
	}

	return s, nil
}
