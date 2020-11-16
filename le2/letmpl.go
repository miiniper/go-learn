package le2

import (
	"errors"
	"html/template"
	"net/http"
)

/*
1。创建模版
2。解析模版
3。渲染模版
*/

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		errors.New("parse file err")
		return
	}
	tmpl.Execute(w, "everyone")
}
