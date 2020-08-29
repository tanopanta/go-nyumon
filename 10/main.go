package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ハンドラだよ")
}

func omikujiHandeler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(4)

	name := "名無し"
	if n := r.FormValue("name"); n != "" {
		name = n
	}

	fmt.Fprintf(w, "%sさんの運勢は %s !\n", name, Omifuda(num).String())
}
func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)
	http.HandleFunc("/omikuji", omikujiHandeler)
	http.ListenAndServe(":8080", nil)
}
