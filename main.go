package main

import (
	"fmt"
	"os"
)

// Добавьте этот импорт если его нет
// import "ascii-art-web/ascii"

func checkFiles() {
	// Проверяем существование необходимых файлов
	files := []string{
		"static/index.html",
		"banners/standard.txt",
		"banners/shadow.txt",
		"banners/thinkertoy.txt",
	}

	fmt.Println("Checking required files:")
	allFilesExist := true
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("❌ MISSING: %s\n", file)
			allFilesExist = false
		} else {
			fmt.Printf("✅ FOUND: %s\n", file)
		}
	}

	if !allFilesExist {
		fmt.Println("Some required files are missing. Please check the file structure.")
	}
}

func main() {
	fmt.Println("Starting ASCII Art Web Server...")
	checkFiles() // Добавляем проверку файлов

	// 	// ВРЕМЕННО: сгенерируйте реальные хеши
	// fmt.Println("Generating REAL banner hashes:")
	// hashes := ascii.GenerateBannerHashes()
	// for banner, hash := range hashes {
	// 	fmt.Printf("%s: %s\n", banner, hash)
	// }

	if err := StartServer(); err != nil {
		fmt.Println("Server error:", err)
	}
}
