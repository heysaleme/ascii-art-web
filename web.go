package main

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/ascii"
)

// Парсим шаблон ОДИН раз при старте приложения
var tmpl = template.Must(
	template.ParseFiles("static/index.html"),
)

type PageData struct {
	Text   string
	Banner string
	Result string
}

func StartServer() error {
	// Статика
	http.Handle("/style/", http.StripPrefix(
		"/style/",
		http.FileServer(http.Dir("style")),
	))

	// Роуты
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/generate", handleGenerate)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/404", handleNotFound)

	fmt.Println("Server started at http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	// Просто рендерим пустую страницу
	err := tmpl.Execute(w, PageData{})
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	if banner == "" {
		banner = "standard"
	}

	// Генерация ASCII-art
	resultText, err := ascii.GenerateASCII(text, "banners/"+banner+".txt")
	if err != nil {
		resultText = fmt.Sprintf("Error: %v", err)
	}

	err = tmpl.Execute(w, PageData{
		Text:   text,
		Banner: banner,
		Result: resultText,
	})
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(
		w,
		"404 Not Found: The page '%s' does not exist.\n",
		r.URL.Path,
	)
}
