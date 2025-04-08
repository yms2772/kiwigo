package kiwigo

//#include "kiwi/capi.h"
import "C"

type OptionsBuild int

var _ Options = (*OptionsBuild)(nil)

func (o OptionsBuild) toInt() C.int {
	if !o.validate() {
		return C.int(OptionsBuildKiwiBuildDefault)
	}
	return C.int(o)
}

func (o OptionsBuild) validate() bool {
	switch o {
	case OptionsBuildKiwiBuildIntegrateAllomorph,
		OptionsBuildKiwiBuildLoadDefaultDict,
		OptionsBuildKiwiBuildLoadTypoDict,
		OptionsBuildKiwiBuildLoadMultiDict,
		OptionsBuildKiwiBuildDefault:
		return true
	default:
		return false
	}
}

const (
	OptionsBuildKiwiBuildIntegrateAllomorph OptionsBuild = C.KIWI_BUILD_INTEGRATE_ALLOMORPH
	OptionsBuildKiwiBuildLoadDefaultDict    OptionsBuild = C.KIWI_BUILD_LOAD_DEFAULT_DICT
	OptionsBuildKiwiBuildLoadTypoDict       OptionsBuild = C.KIWI_BUILD_LOAD_TYPO_DICT
	OptionsBuildKiwiBuildLoadMultiDict      OptionsBuild = C.KIWI_BUILD_LOAD_MULTI_DICT
	OptionsBuildKiwiBuildDefault            OptionsBuild = C.KIWI_BUILD_DEFAULT
)
