package main

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound     = DictionaryErr("could not find definition")
	ErrExistingWord = DictionaryErr("existing word")
	ErrWordNotFound = DictionaryErr("could not find word")
)

func (d Dictionary) Search(word string) (string, error) {
	if def, ok := d[word]; ok {
		return def, nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dictionary) Add(word string, def string) error {
	if _, err := d.Search(word); err == nil {
		return ErrExistingWord
	} else {
		d[word] = def
		return nil
	}
}

func (d Dictionary) Update(word string, def string) (string, error) {
	if old, err := d.Search(word); err == nil {
		d[word] = def
		return old, nil
	} else {
		return "", ErrWordNotFound
	}
}

func (d Dictionary) Delete(word string) (string, error) {
	if old, err := d.Search(word); err == nil {
		delete(d, word)
		return old, nil
	} else {
		return "", ErrWordNotFound
	}
}
