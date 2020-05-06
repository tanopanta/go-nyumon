package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n         bool
	lineCount = 1
)

func init() {
	flag.BoolVar(&n, "n", false, "行番号の表示")
}

func main() {
	flag.Parse()

	// コマンド引数のファイル名を取得
	var fileNames []string
	for _, v := range flag.Args() {
		fileNames = append(fileNames, v)
	}

	for _, fileName := range fileNames {
		catFile(fileName)
	}

}

func catFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	// 関数終了時に閉じる
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		if n {
			fmt.Fprint(os.Stdout, lineCount, ": ")
			lineCount++
		}
		fmt.Fprintln(os.Stdout, scanner.Text())
	}
	// まとめてエラー処理をする
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return err
}
