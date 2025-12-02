package ascii

import "strings"

// ValidBannerHashes содержит SHA-256 хеши оригинальных баннеров
var ValidBannerHashes = map[string]string{
	"standard":   "c3ec7584fb7ecfbd739e6b3f6f63fd1fe557d2ae3e24f870730d9cf8b2559e94",
	"shadow":     "78ccd616680eb9068fe1465db1c852ceaffd8c0f318e3aa0414e1635508e85bf",
	"thinkertoy": "e3c7a11f41a473d9b0d3bf2132a8f6dabb754bd16efa3897fa835a432d3b9caa",
}

// GetBannerNameFromPath извлекает имя баннера из пути
func GetBannerNameFromPath(path string) string {
	switch {
	case strings.Contains(path, "standard"):
		return "standard"
	case strings.Contains(path, "shadow"):
		return "shadow"
	case strings.Contains(path, "thinkertoy"):
		return "thinkertoy"
	default:
		return ""
	}
}
