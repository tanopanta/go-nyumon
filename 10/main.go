package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/tanopanta/go-nyumon/10/model"
)

// Oracle はお告げです
type Oracle struct {
	Name    string
	Omifuda string
}

var tpl *template.Template

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ハンドラだよ")
}

func omikujiHandeler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(4)

	name := "名無し"
	if n := r.FormValue("name"); n != "" {
		name = n
	}

	o := Oracle{
		Name:    string(name),
		Omifuda: model.Omifuda(num).String(),
	}

	err := tpl.Execute(w, o)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	tpl = template.Must(template.New("omi").
		Parse("<html><body>{{.Name}} さんの運勢は <strong>{{.Omifuda}}</strong> です！</body></html>"))

	http.HandleFunc("/", handler)
	http.HandleFunc("/omikuji", omikujiHandeler)
	http.ListenAndServe(":8080", nil)
}
