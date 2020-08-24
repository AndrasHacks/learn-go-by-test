package maps

import "errors"

type Dictionary map[string]string

const (
	wordNotFoundErrorMessage     string = "could not find the word you were looking for"
	keyAlreadyExistsErrorMessage string = "key already exists!"
)

var (
	WordNotFoundError     = errors.New(wordNotFoundErrorMessage)
	KeyAlreadyExistsError = errors.New(keyAlreadyExistsErrorMessage)
)

func (d Dictionary) Search(key string) (string, error) {
	if d[key] != "" {
		return d[key], nil
	} else {
		return "", WordNotFoundError
	}
}

// Map is a reference type!
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err != WordNotFoundError {
		return KeyAlreadyExistsError
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == WordNotFoundError {
		return WordNotFoundError
	}
	d[word] = definition
	return nil
}
