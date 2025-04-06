package kiwigo

//#include "kiwi/capi.h"
import "C"

type kiwiError string

func (ke kiwiError) Error() string {
	return string(ke)
}

func getKiwiError() kiwiError {
	return kiwiError(C.GoString(C.kiwi_error()))
}
