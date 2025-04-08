package kiwigo

//#include "kiwi/capi.h"
import "C"

type OptionsAnalyze int

var _ Options = (*OptionsAnalyze)(nil)

func (o OptionsAnalyze) toInt() C.int {
	if !o.validate() {
		return C.int(OptionsAnalyzeKiwiMatchAll)
	}
	return C.int(o)
}

func (o OptionsAnalyze) validate() bool {
	switch o {
	case OptionsAnalyzeKiwiMatchURL,
		OptionsAnalyzeKiwiMatchEmail,
		OptionsAnalyzeKiwiMatchHashtag,
		OptionsAnalyzeKiwiMatchMention,
		OptionsAnalyzeKiwiMatchSerial,
		OptionsAnalyzeKiwiMatchNormalizeCoda,
		OptionsAnalyzeKiwiMatchJoinNounPrefix,
		OptionsAnalyzeKiwiMatchJoinNounSuffix,
		OptionsAnalyzeKiwiMatchJoinVerbSuffix,
		OptionsAnalyzeKiwiMatchJoinAdjSuffix,
		OptionsAnalyzeKiwiMatchJoinAdvSuffix,
		OptionsAnalyzeKiwiMatchJoinVSuffix,
		OptionsAnalyzeKiwiMatchJoinAffix,
		OptionsAnalyzeKiwiMatchSplitComplex,
		OptionsAnalyzeKiwiMatchZCoda,
		OptionsAnalyzeKiwiMatchCompatibleJamo,
		OptionsAnalyzeKiwiMatchSplitSaisiot,
		OptionsAnalyzeKiwiMatchMergeSaisiot,
		OptionsAnalyzeKiwiMatchAll,
		OptionsAnalyzeKiwiMatchAllWithNormalizing:
		return true
	default:
		return false
	}
}

const (
	OptionsAnalyzeKiwiMatchURL                OptionsAnalyze = C.KIWI_MATCH_URL
	OptionsAnalyzeKiwiMatchEmail              OptionsAnalyze = C.KIWI_MATCH_EMAIL
	OptionsAnalyzeKiwiMatchHashtag            OptionsAnalyze = C.KIWI_MATCH_HASHTAG
	OptionsAnalyzeKiwiMatchMention            OptionsAnalyze = C.KIWI_MATCH_MENTION
	OptionsAnalyzeKiwiMatchSerial             OptionsAnalyze = C.KIWI_MATCH_SERIAL
	OptionsAnalyzeKiwiMatchNormalizeCoda      OptionsAnalyze = C.KIWI_MATCH_NORMALIZE_CODA
	OptionsAnalyzeKiwiMatchJoinNounPrefix     OptionsAnalyze = C.KIWI_MATCH_JOIN_NOUN_PREFIX
	OptionsAnalyzeKiwiMatchJoinNounSuffix     OptionsAnalyze = C.KIWI_MATCH_JOIN_NOUN_SUFFIX
	OptionsAnalyzeKiwiMatchJoinVerbSuffix     OptionsAnalyze = C.KIWI_MATCH_JOIN_VERB_SUFFIX
	OptionsAnalyzeKiwiMatchJoinAdjSuffix      OptionsAnalyze = C.KIWI_MATCH_JOIN_ADJ_SUFFIX
	OptionsAnalyzeKiwiMatchJoinAdvSuffix      OptionsAnalyze = C.KIWI_MATCH_JOIN_ADV_SUFFIX
	OptionsAnalyzeKiwiMatchJoinVSuffix        OptionsAnalyze = C.KIWI_MATCH_JOIN_V_SUFFIX
	OptionsAnalyzeKiwiMatchJoinAffix          OptionsAnalyze = C.KIWI_MATCH_JOIN_AFFIX
	OptionsAnalyzeKiwiMatchSplitComplex       OptionsAnalyze = C.KIWI_MATCH_SPLIT_COMPLEX
	OptionsAnalyzeKiwiMatchZCoda              OptionsAnalyze = C.KIWI_MATCH_Z_CODA
	OptionsAnalyzeKiwiMatchCompatibleJamo     OptionsAnalyze = C.KIWI_MATCH_COMPATIBLE_JAMO
	OptionsAnalyzeKiwiMatchSplitSaisiot       OptionsAnalyze = C.KIWI_MATCH_SPLIT_SAISIOT
	OptionsAnalyzeKiwiMatchMergeSaisiot       OptionsAnalyze = C.KIWI_MATCH_MERGE_SAISIOT
	OptionsAnalyzeKiwiMatchAll                OptionsAnalyze = C.KIWI_MATCH_ALL
	OptionsAnalyzeKiwiMatchAllWithNormalizing OptionsAnalyze = C.KIWI_MATCH_ALL_WITH_NORMALIZING
)
