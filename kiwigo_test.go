package kiwigo

import (
	"os"
	"testing"
)

const modelPath = "./models/base"

func TestKiwi_Analyze(t *testing.T) {
	k, err := New(modelPath, 0, OptionsBuildKiwiBuildDefault)
	if err != nil {
		t.Fatalf("failed to create kiwi: %v", err)
	}

	body, _ := os.ReadFile("./testdata/keywords.txt")
	result, err := k.Analyze(string(body), 1, OptionsAnalyzeKiwiMatchAll)
	if err != nil {
		t.Fatalf("failed to analyze: %v", err)
	}

	for _, token := range result {
		t.Logf("token: %v", token)
	}
}

//func TestKiwiBuilder_ExtractWords(t *testing.T) {
//	kb, err := NewBuilder(modelPath, 0, OptionsBuildKiwiBuildDefault)
//	if err != nil {
//		t.Fatalf("failed to create kiwi builder: %v", err)
//	}
//
//	body, _ := os.ReadFile("./testdata/keywords.txt")
//	result, err := kb.ExtractWords(string(body), 5, 5, 0.0, -3.0)
//	if err != nil {
//		t.Fatalf("failed to extract words: %v", err)
//	}
//
//	for _, token := range result {
//		t.Logf("token: %v", token)
//	}
//}
