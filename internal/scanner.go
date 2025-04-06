package internal

import (
	"bufio"
	"strings"
)

type Scanner interface {
	Scan() bool
	Text() string
	Bytes() []byte
	First()
}

type scanner struct {
	lines []string
	idx   int
}

func NewScanner(input string) Scanner {
	sc := bufio.NewScanner(strings.NewReader(input))
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return &scanner{
		lines: lines,
		idx:   -1,
	}
}

func (s *scanner) First() {
	s.idx = -1
}

func (s *scanner) Scan() bool {
	if s.idx+1 >= len(s.lines) {
		return false
	}

	s.idx++
	return true
}

func (s *scanner) Text() string {
	if s.idx >= 0 && s.idx < len(s.lines) {
		return s.lines[s.idx]
	}
	return ""
}

func (s *scanner) Bytes() []byte {
	return []byte(s.Text())
}
