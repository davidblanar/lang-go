package main

type ReadStream struct {
  input string
  pos int
  lineno int
  col int
}

func (rs *ReadStream) next() string {
  ch := string(rs.input[rs.pos])
  rs.pos++
  rs.col++
  if ch == "\n" {
    rs.lineno++
    rs.col = 1
  }
  return ch
}

func (rs *ReadStream) hasNext() bool {
  return rs.pos < len(rs.input)
}

func (rs *ReadStream) peek() string {
  return string(rs.input[rs.pos])
}
