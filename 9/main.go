package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ch := input(os.Stdin)
	for {
		msg := <-ch
		if msg == "exit" || msg == "close" {
			break
		}

		fmt.Print(">")
		fmt.Println(msg)
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()

	return ch
}
