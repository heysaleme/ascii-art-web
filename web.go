package main

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/ascii"
)

func StartServer() error {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/generate", handleGenerate)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/404", handleNotFound)

	fmt.Println("Server started at http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}

type PageData struct {
	Text   string
	Banner string
	Result string
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Рендерим пустой результат
	tmpl.Execute(w, PageData{Result: ""})
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	if banner == "" {
		banner = "standard"
	}

	// Генерируем ASCII
	resultText, err := ascii.GenerateASCII(text, "banners/"+banner+".txt")
	if err != nil {
		// Если ошибка, показываем её в поле результата
		resultText = fmt.Sprintf("Error: %v", err)
	}

	tmpl.Execute(w, PageData{
		Text:   text,
		Banner: banner,
		Result: resultText,
	})
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 Not Found: The page '%s' does not exist.\n", r.URL.Path)
}
