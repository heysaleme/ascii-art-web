package ascii

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// ValidateBannerHash проверяет хеш баннера против известных корректных хешей
func ValidateBannerHash(path string, data []byte) error {
	bannerName := GetBannerNameFromPath(path)
	if bannerName == "" {
		return nil // Не проверяем кастомные баннеры
	}

	expectedHash, exists := ValidBannerHashes[bannerName]
	if !exists {
		return nil // Не проверяем неизвестные баннеры
	}

	actualHash := fmt.Sprintf("%x", sha256.Sum256(data))
	if actualHash != expectedHash {
		return fmt.Errorf("banner file has been modified or corrupted")
	}

	return nil
}

// ValidateBannerStructure проверяет общую структуру файла баннера
func ValidateBannerStructure(lines []string) error {
	expectedTotalLines := 9 * (126 - 32 + 1) // 9 строк на символ × 95 символов
	if len(lines) != expectedTotalLines {
		return fmt.Errorf("expected %d lines, got %d", expectedTotalLines, len(lines))
	}

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		return fmt.Errorf("banner file is empty")
	}

	return nil
}

// ValidateCharacterStrict СТРОГО проверяет целостность символа
func ValidateCharacterStrict(char rune, lines []string) error {
	if len(lines) != 9 {
		return fmt.Errorf("expected 9 lines, got %d", len(lines))
	}

	// Для пробела проверяем, что все строки пустые
	if char == 32 {
		for i, line := range lines {
			if strings.TrimSpace(line) != "" {
				return fmt.Errorf("space character should be empty, but line %d has content: %q", i, line)
			}
		}
		return nil
	}

	// Для остальных символов проверяем согласованность
	expectedWidth := -1
	hasContent := false

	for i, line := range lines {
		// Проверяем допустимые символы в ASCII арте
		if !isValidAsciiArtLine(line) {
			return fmt.Errorf("line %d contains invalid characters: %q", i, line)
		}

		if strings.TrimSpace(line) != "" {
			hasContent = true
			if expectedWidth == -1 {
				expectedWidth = len(line)
			} else if len(line) != expectedWidth {
				return fmt.Errorf("inconsistent width at line %d: got %d, expected %d", i, len(line), expectedWidth)
			}
		} else if expectedWidth != -1 {
			// Если мы уже нашли контент, а потом пустая строка - проверяем длину
			if len(line) != expectedWidth {
				return fmt.Errorf("empty line %d has wrong length: got %d, expected %d", i, len(line), expectedWidth)
			}
		}
	}

	if !hasContent {
		return fmt.Errorf("character has no content")
	}

	return nil
}

// isValidAsciiArtLine проверяет, что строка содержит только допустимые символы ASCII арта
func isValidAsciiArtLine(line string) bool {
	for _, r := range line {
		if !isValidAsciiArtChar(r) {
			return false
		}
	}
	return true
}

// isValidAsciiArtChar проверяет допустимость символа в ASCII арте
func isValidAsciiArtChar(r rune) bool {
	// Допустимые символы: пробел и основные символы для рисования ASCII арта
	validChars := " _/\\|()<>[]{}'`!@#$%^&*+-=:;\",.VOo0"
	return r == ' ' || strings.ContainsRune(validChars, r)
}
