package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
