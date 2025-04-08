package kiwigo

/*
#cgo LDFLAGS: -lkiwi

#include <stdlib.h>
#include <string.h>
#include "kiwi/capi.h"

extern int KiwiReader(int idx, char* buf, void* userData);
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"

	"github.com/yms2772/kiwigo/internal"
)

//export KiwiReaderImpl
func KiwiReaderImpl(idx C.int, buf *C.char, userData unsafe.Pointer) C.int {
	scanner := cgo.Handle(userData).Value().(internal.Scanner)
	if buf == nil {
		if idx == 0 {
			scanner.First()
		}

		if !scanner.Scan() {
			return C.int(0)
		}

		text := scanner.Bytes()
		return C.int(len(text) + 1)
	}

	textCString := C.CString(scanner.Text())
	defer C.free(unsafe.Pointer(textCString))

	C.strcpy(buf, textCString)
	return C.int(0)
}

type Kiwi interface {
	Analyze(text string, topN int, options OptionsAnalyze) ([]AnalyzeResult, error)
}

type kiwi struct {
	h C.kiwi_h
}

var _ Kiwi = (*kiwi)(nil)

type AnalyzeToken struct {
	Position int
	Tag      string
	Form     string
}

type AnalyzeResult struct {
	Tokens []AnalyzeToken
	Score  float32
}

func (k *kiwi) Analyze(text string, topN int, options OptionsAnalyze) ([]AnalyzeResult, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	kiwiResH := C.kiwi_analyze(k.h, cText, C.int(topN), C.int(options), nil, nil)
	if kiwiResH == nil {
		return nil, getKiwiError()
	}
	defer C.kiwi_res_close(kiwiResH)

	result := make([]AnalyzeResult, int(C.kiwi_res_size(kiwiResH)))
	for i := range result {
		tokens := make([]AnalyzeToken, int(C.kiwi_res_word_num(kiwiResH, C.int(i))))

		for j := 0; j < len(tokens); j++ {
			tokens[j] = AnalyzeToken{
				Form:     C.GoString(C.kiwi_res_form(kiwiResH, C.int(i), C.int(j))),
				Tag:      C.GoString(C.kiwi_res_tag(kiwiResH, C.int(i), C.int(j))),
				Position: int(C.kiwi_res_position(kiwiResH, C.int(i), C.int(j))),
			}
		}

		result[i] = AnalyzeResult{
			Tokens: tokens,
			Score:  float32(C.kiwi_res_prob(kiwiResH, C.int(i))),
		}
	}
	return result, nil
}

func New(modelPath string, numThread int, options OptionsBuild) (Kiwi, error) {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))
	k := &kiwi{
		h: C.kiwi_init(cModelPath, C.int(numThread), options.toInt()),
	}
	if k.h == nil {
		return nil, getKiwiError()
	}
	return k, nil
}

type KiwiBuilder interface {
	Build(typoCostThreshold float32) Kiwi
	AddWord(form string, tag string, score float32)
	ExtractWords(input string, minCnt, maxWordLen int, minScore, posThreshold float32) ([]WordInfo, error)
}

type kiwiBuilder struct {
	h C.kiwi_builder_h
}

var _ KiwiBuilder = (*kiwiBuilder)(nil)

type WordInfo struct {
	Form     string
	Freq     int
	POSScore float32
	Score    float32
}

func (kb *kiwiBuilder) Build(typoCostThreshold float32) Kiwi {
	kiwiH := C.kiwi_builder_build(kb.h, nil, C.float(typoCostThreshold))
	if kiwiH == nil {
		return nil
	}
	return &kiwi{h: kiwiH}
}

func (kb *kiwiBuilder) AddWord(form string, tag string, score float32) {
	cForm := C.CString(form)
	defer C.free(unsafe.Pointer(cForm))
	cTag := C.CString(tag)
	defer C.free(unsafe.Pointer(cTag))

	C.kiwi_builder_add_word(kb.h, cForm, cTag, C.float(score))
}

func (kb *kiwiBuilder) ExtractWords(input string, minCnt, maxWordLen int, minScore, posThreshold float32) ([]WordInfo, error) {
	scanner := internal.NewScanner(input)
	h := cgo.NewHandle(scanner)
	defer h.Delete()

	kiwiWsH := C.kiwi_builder_extract_words(
		kb.h,
		C.kiwi_reader_t(C.KiwiReader),
		unsafe.Pointer(h),
		C.int(minCnt),
		C.int(maxWordLen),
		C.float(minScore),
		C.float(posThreshold),
	)
	if kiwiWsH == nil {
		return nil, getKiwiError()
	}
	defer C.kiwi_ws_close(kiwiWsH)

	result := make([]WordInfo, int(C.kiwi_ws_size(kiwiWsH)))
	for i := range result {
		idx := C.int(i)
		result[i] = WordInfo{
			Form:     C.GoString(C.kiwi_ws_form(kiwiWsH, idx)),
			Freq:     int(C.kiwi_ws_freq(kiwiWsH, idx)),
			POSScore: float32(C.kiwi_ws_pos_score(kiwiWsH, idx)),
			Score:    float32(C.kiwi_ws_score(kiwiWsH, idx)),
		}
	}
	return result, nil
}

func NewBuilder(modelPath string, numThread int, options OptionsBuild) (KiwiBuilder, error) {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	k := &kiwiBuilder{
		h: C.kiwi_builder_init(cModelPath, C.int(numThread), options.toInt()),
	}
	if k.h == nil {
		return nil, getKiwiError()
	}
	return k, nil
}
