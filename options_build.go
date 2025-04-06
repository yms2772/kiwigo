package kiwigo

//#include "kiwi/capi.h"
import "C"

type OptionsBuild int

var _ Options = (*OptionsBuild)(nil)

func (o OptionsBuild) toInt() C.int {
	return C.int(o)
}

const (
	OptionsBuildKiwiBuildIntegrateAllomorph  OptionsBuild = C.KIWI_BUILD_INTEGRATE_ALLOMORPH
	OptionsBuildKiwiBuildLoadDefaultDict     OptionsBuild = C.KIWI_BUILD_LOAD_DEFAULT_DICT
	OptionsBuildKiwiBuildLoadTypoDict        OptionsBuild = C.KIWI_BUILD_LOAD_TYPO_DICT
	OptionsBuildKiwiBuildLoadMultiDict       OptionsBuild = C.KIWI_BUILD_LOAD_MULTI_DICT
	OptionsBuildKiwiBuildDefault             OptionsBuild = C.KIWI_BUILD_DEFAULT
	OptionsBuildKiwiBuildModelTypeDefault    OptionsBuild = C.KIWI_BUILD_MODEL_TYPE_DEFAULT
	OptionsBuildKiwiBuildModelTypeKNLM       OptionsBuild = C.KIWI_BUILD_MODEL_TYPE_KNLM
	OptionsBuildKiwiBuildModelTypeSBG        OptionsBuild = C.KIWI_BUILD_MODEL_TYPE_SBG
	OptionsBuildKiwiBuildModelTypeCONG       OptionsBuild = C.KIWI_BUILD_MODEL_TYPE_CONG
	OptionsBuildKiwiBuildModelTypeCONGGlobal OptionsBuild = C.KIWI_BUILD_MODEL_TYPE_CONG_GLOBAL
)
