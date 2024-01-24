package main

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
	return ""
}

func (u *URLShortener) Redirect(shortURL string) (string, error) {
	// Your implementation here
	return "", nil
}

func main() {
	// Test your implementation here
}
