package urlshortcutter

import (
	"math/rand"
	"time"
)

const symbols = "ABCDEFGHIJKLMNO5PQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

var urlToShort = make(map[string]string)
var urlToLong = make(map[string]string)

func randomizer() (shortURL string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		shortURL += string(symbols[r.Intn(62)])
	}
	return shortURL
}

// Converter converts an user URL to short version
func Converter(URL string) (shortURL string) {

	value, ok := urlToShort[URL]
	switch ok {
	case true:
		return value
	default:
		shortURL := randomizer()
		urlToShort[URL] = shortURL
		urlToLong[shortURL] = URL
		return shortURL
	}
}

// GetShortURL returns long url from short
func GetShortURL(shortURL string) (longURL string) {
	longURL, ok := urlToLong[shortURL]
	switch ok {
	case true:
		return longURL
	default:
		return ""
	}
}
