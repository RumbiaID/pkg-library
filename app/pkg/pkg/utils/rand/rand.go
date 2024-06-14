package rand

import gonanoid "github.com/matoous/go-nanoid"

// Generate generates a random string with the given alphabet and size.
// If the alphabet is "*", it will use the default alphabet.
// If the size is less than or equal to 0, it will use the default size.
func Generate(rawAlphabet string, size int, prefix ...string) (string, error) {
	if rawAlphabet == "*" {
		rawAlphabet = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if size <= 0 {
		size = 21
	}
	r, e := gonanoid.Generate(rawAlphabet, size)
	if e != nil {
		return "", e
	}
	if len(prefix) > 0 {
		return prefix[0] + r, nil
	}
	return r, nil
}

// GenerateNumeric generates a random string with the numeric alphabet and size.
// If the size is less than or equal to 0, it will use the default size.
func GenerateNumeric(size int, prefix ...string) (string, error) {
	if size <= 0 {
		size = 21
	}
	r, e := gonanoid.Generate("123456789", size)
	if e != nil {
		return "", e
	}
	if len(prefix) > 0 {
		return prefix[0] + r, nil
	}
	return r, nil
}
