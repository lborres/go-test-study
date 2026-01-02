package dictionary

import "errors"

var ErrNotFound = errors.New("could not find definition for the given term")

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
