package main

import (
	"net/http"
)

func main() {
	// обозначаем, что на запрос по пути "/" возвращается строка "Hello World"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// тело ответа — это массив байт
		w.Write([]byte("Hello world!"))
	})

	// запускаем веб-приложение для обработки запросов
	http.ListenAndServe(":80", nil)
}
