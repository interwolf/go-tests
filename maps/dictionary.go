package maps

// Dictionary waps map
type Dictionary map[string]string

const (
	// ErrorNotFound is a message for no-existence error
	ErrorNotFound = DictErr("cannot find the word")
	// ErrorExistence happens when adding existing words
	ErrorExistence = DictErr("cannot add the word which already exists")
	// ErrorNotExistence happens when updating non existing words
	ErrorNotExistence = DictErr("cannot update the word which not exists")
)

// DictErr implements the error interface
type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

// Search gives the explanation of input word
func (d Dictionary) Search(word string) (string, error) {
	var err error

	result, ok := d[word]
	if !ok {
		err = ErrorNotFound
	} else {
		err = nil
	}

	return result, err
}

// Add will add a word to dict
func (d Dictionary) Add(word, explain string) error {
	_, exist := d[word]
	if exist {
		return ErrorExistence
	}

	d[word] = explain
	return nil
}

// Update will modify the explanation of a word
func (d Dictionary) Update(word, explain string) error {
	_, exist := d[word]
	if !exist {
		return ErrorNotExistence
	}

	d[word] = explain
	return nil
}

// Delete will delete word from dict
// word must exist in dict
func (d Dictionary) Delete(word string) error {
	_, exist := d[word]
	if !exist {
		return ErrorNotExistence
	}

	delete(d, word)
	return nil
}
