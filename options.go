package kiwigo

import "C"

type Options interface {
	toInt() C.int
	validate() bool
}
