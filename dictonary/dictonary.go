package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	if _, ok := d[word]; !ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}
func (d Dictionary) Delete(s string) {
	/*
		_, err := d.Search(word)
		switch err {
		case ErrNotFound:
			return ErrWordDoesNotExist
		case nil:
			d[word] = definition
		default:
			return err
		}
		return nil
	*/
	delete(d, s)
}
