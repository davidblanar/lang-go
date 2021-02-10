package main

import "fmt"

type ReadStream struct {
	input string
	pos int
}

func (rs *ReadStream) next() string {
	ch := string(rs.input[rs.pos])
	rs.pos++
	return ch
}

func (rs *ReadStream) hasNext() bool {
	return rs.pos < len(rs.input)
}

func (rs *ReadStream) peek() string {
	return string(rs.input[rs.pos])
}
