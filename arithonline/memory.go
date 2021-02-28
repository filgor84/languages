package arithonline

import (
	"strings"
)

type memory interface {
	save(interface{}) int
	load(int) interface{}
	toString() string
}

//ListOfWords is a dummy memory example to get a representation of the lexer results
type ListOfWords struct {
	words []string
}

func (l *ListOfWords) save(data interface{}) int {
	top := len(l.words)
	l.words = append(l.words, data.(string))
	return top
}

func (l ListOfWords) load(id int) interface{} {
	return l.words[id]
}

func (l ListOfWords) toString() string {
	return strings.Join(l.words, " ")
}
