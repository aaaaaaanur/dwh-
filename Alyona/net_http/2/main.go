package main

import (
	"net/http"
	"os"
)

var mePage = `

`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Hexlet"))
	})

	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		// прочитать файл html/me.html
		content, err := os.ReadFile("html/me.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(content)
	})

	http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		// считываем параметр page из запроса
		page := r.URL.Query().Get("page")

		// https: //localhost:80/path?p=1#ff

		// рассчитываем, какую страницу нужно вернуть
		var pageCourses string
		switch page {
		case "":
			pageCourses = "Введите номер курса!"
		case "1":
			pageCourses = "Как написать свой первый \"Hello world\" на go..."
		case "2":
			pageCourses = "Как работает сборщик мусора в go..."
		default:
			pageCourses = "Курс в разработке..."
		}

		// возвращаем страницу курса
		w.Write([]byte(pageCourses))
	})

	http.ListenAndServe(":80", nil)
}
