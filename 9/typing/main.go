package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	bc := context.Background()
	t := 30 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	typing(ctx)
}

func typing(ctx context.Context) {
	fmt.Println("タイピングゲームスタート！")

	inputCh := input(os.Stdin)
	var word string
	total := 0
	for {
		word = randomWord()
		fmt.Println(word)
		select {
		case <-ctx.Done():
			fmt.Println("タイムアップ!")
			fmt.Printf("正解数は [ %d ] !!\n", total)
			return
		case line := <-inputCh:
			if line == word {
				fmt.Println("正解!")
				fmt.Println()
				total++
			} else {
				fmt.Println("不正解")
				fmt.Println()
			}
		}
	}

}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}

func randomWord() string {
	word := []string{
		"apple",
		"orange",
		"banana",
	}

	return word[rand.Intn(len(word))]
}
