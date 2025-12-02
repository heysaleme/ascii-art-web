package ascii

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

// LoadBanner загружает баннер из файла с СТРОГОЙ проверкой целостности
func LoadBanner(path string) (map[rune][]string, error) {
	// Проверяем существование файла
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("banner not found: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read banner file: %v", err)
	}

	// Проверяем хеш файла для стандартных баннеров
	if err := ValidateBannerHash(path, data); err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	// СТРОГАЯ проверка структуры файла
	if err := ValidateBannerStructure(lines); err != nil {
		return nil, fmt.Errorf("banner file structure invalid: %v", err)
	}

	chars := make(map[rune][]string)
	ascii := rune(32)
	loadedChars := 0

	// Проверяем целостность всех символов
	for i := 0; i < len(lines); i += 9 {
		if ascii > 126 {
			break
		}

		characterLines := lines[i : i+9]

		// СТРОГАЯ проверка каждого символа
		if err := ValidateCharacterStrict(ascii, characterLines); err != nil {
			return nil, fmt.Errorf("banner file corrupted at character %q: %v", string(ascii), err)
		}

		chars[ascii] = characterLines
		ascii++
		loadedChars++
	}

	// Проверяем, что загружены все необходимые символы
	if loadedChars < (126 - 32 + 1) {
		return nil, fmt.Errorf("banner file is incomplete: loaded %d characters, expected %d", loadedChars, 126-32+1)
	}

	return chars, nil
}

// GenerateBannerHashes вспомогательная функция для генерации хешей
func GenerateBannerHashes() map[string]string {
	hashes := make(map[string]string)
	banners := []string{"standard", "shadow", "thinkertoy"}

	for _, banner := range banners {
		path := "banners/" + banner + ".txt"
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", banner, err)
			continue
		}
		hash := fmt.Sprintf("%x", sha256.Sum256(data))
		hashes[banner] = hash
		fmt.Printf("%s: %s\n", banner, hash)
	}

	return hashes
}

// SplitInput разбивает текст пользователя на строки по реальным переносам
func SplitInput(input string) []string {
	cleaned := strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(cleaned, "\n")
}

// RenderLine рендерит одну строку текста в ASCII
func RenderLine(line string, font map[rune][]string) (string, error) {
	var result strings.Builder

	if len(font) == 0 {
		return "", fmt.Errorf("font not loaded or empty")
	}

	// Проверяем только символы, которые есть в баннере
	for _, c := range line {
		if c < 32 || c > 126 {
			return "", fmt.Errorf("character '%c' is not supported", c)
		}
	}

	firstLineEmpty := true
	for _, c := range line {
		if art, ok := font[c]; ok && len(art) > 0 && strings.TrimSpace(art[0]) != "" {
			firstLineEmpty = false
			break
		}
	}

	startIndex := 0
	if firstLineEmpty {
		startIndex = 1
	}

	for i := startIndex; i < 9; i++ {
		lineContent := ""
		for _, c := range line {
			if art, ok := font[c]; ok {
				if i < len(art) {
					lineContent += art[i]
				} else if len(art) > 0 && len(art[0]) > 0 {
					lineContent += strings.Repeat(" ", len(art[0]))
				}
			} else {
				return "", fmt.Errorf("character '%c' not found in font", c)
			}
		}
		result.WriteString(lineContent)
		result.WriteRune('\n')
	}

	return result.String(), nil
}

// GenerateASCII генерирует полный ASCII арт с обработкой ошибок
func GenerateASCII(text, banner string) (string, error) {
	font, err := LoadBanner(banner)
	if err != nil {
		return "", fmt.Errorf("failed to load banner: %v", err)
	}

	lines := SplitInput(text)
	if len(lines) == 1 && lines[0] == "" {
		return "", nil
	}

	var result strings.Builder
	for _, line := range lines {
		if line == "" {
			// пустая строка — просто добавляем перенос
			result.WriteString("\n")
			continue
		}

		rendered, err := RenderLine(line, font)
		if err != nil {
			return "", fmt.Errorf("failed to render line: %v", err)
		}
		result.WriteString(rendered)
	}

	return result.String(), nil
}
