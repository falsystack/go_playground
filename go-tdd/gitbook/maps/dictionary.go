package maps

import "errors"

var ErrNotFound = errors.New("could not find word")

type Dictionary map[string]string

func (d *Dictionary) Search(word string) (string, error) {
	def, ok := (*d)[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}
