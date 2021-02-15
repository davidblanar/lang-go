package main

import (
	"regexp"
	"strings"
	"strconv"
	"log"
	"fmt"
)

type Token struct {
	tokenType string
	val interface{}
	lineno int
	col int
}

type Tokenizer struct {
	rs ReadStream
	symbols map[string]bool
	identifierChars *regexp.Regexp
	numberChars *regexp.Regexp
	whiteSpace *regexp.Regexp
	tokens []Token
}

func (t *Tokenizer) generate() []Token {
	for {
		if t.rs.hasNext() {
			current := t.rs.peek()
			if current == "#" {
				t._skipComment()
			} else if t._isWhiteSpace(current) {
				t._skip()
			} else if t._isIdentifier(current) {
				t._readIdentifier()
			} else if _, ok := t.symbols[current]; ok {
				t._readSymbol()
			} else if t._isNumber(current) {
				t._readNumber()
			} else if current == "\"" {
				t._readString()
			} else {
				throwError(fmt.Sprintf("Unrecognized character %s", current), t.rs.lineno, t.rs.col)
			}
		} else {
			break
		}
	}
	return t.tokens
}

func (t *Tokenizer) _skipComment() {
	for {
		if t.rs.hasNext() && t.rs.next() != "\n" {
			
		} else {
			break
		}
	}
}

func (t *Tokenizer) _skip() {
	t.rs.next()
}

func (t *Tokenizer) _isWhiteSpace(char string) bool {
	return t.whiteSpace.MatchString(char)
}

func (t *Tokenizer) _isIdentifier(char string) bool {
	return t.identifierChars.MatchString(char)
}

func (t *Tokenizer) _isNumber(char string) bool {
	return t.numberChars.MatchString(char)
}

func (t *Tokenizer) _readIdentifier() {
	var sb strings.Builder
	for {
		if t.rs.hasNext() {
			current := t.rs.peek()
			if t._isWhiteSpace(current) || current == ")" {
				break
			}
			sb.WriteString(t.rs.next())
		} else {
			break
		}
	}
	token := Token{TOKEN_IDENTIFIER, sb.String(), t.rs.lineno, t.rs.col}
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) _readSymbol() {
	current := t.rs.next()
	token := Token{TOKEN_SYMBOL, current, t.rs.lineno, t.rs.col}
	t.tokens = append(t.tokens, token)	
}

func (t *Tokenizer) _readNumber() {
	var sb strings.Builder
	for {
		if t.rs.hasNext() {
			current := t.rs.peek()
			if current == "." || t._isNumber(current) {
				sb.WriteString(t.rs.next())
			} else {
				break
			}
		} else {
			break
		}
	}
	val, err := strconv.ParseFloat(sb.String(), 64)
	if err != nil {
		log.Fatal(err)
	}
	token := Token{TOKEN_NUMBER, val, t.rs.lineno, t.rs.col}
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) _readString() {
	// skip beginning double quote
	t._skip()
	var sb strings.Builder
	for {
		if t.rs.hasNext() {
			current := t.rs.peek()
			if current == "\"" {
				// skip closing double quote
				t._skip()
				break
			} else {
				sb.WriteString(t.rs.next())
			}
		} else {
			break
		}
	}
	token := Token{TOKEN_STRING, sb.String(), t.rs.lineno, t.rs.col}
	t.tokens = append(t.tokens, token)
}
