package kiwigo

/*
#cgo LDFLAGS: -l kiwi

int KiwiReader(int idx, char* buf, void* userData) {
	int KiwiReaderImpl(int, char*, void*);
  return KiwiReaderImpl(idx, buf, userData);
}
*/
import "C"
