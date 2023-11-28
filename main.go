package main

import (
	"errors"
	"fmt"
)

type URLShortener struct {
	urlMap map[string]string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		urlMap: make(map[string]string),
	}
}

func (u *URLShortener) ShortenURL(originalURL string) string {
	// Your implementation here
}

func (u *URLShortener) Redirect(shortURL string) (string, error) {
	// Your implementation here
}

func main() {
	// Test your implementation here
}
