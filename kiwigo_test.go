package kiwigo

import (
	"bytes"
	"os"
	"testing"
)

const modelPath = "./models/base"

func TestKiwi_Analyze(t *testing.T) {
	kb, err := NewBuilder(modelPath, 0, OptionsBuildKiwiBuildDefault)
	if err != nil {
		t.Fatalf("failed to create kiwi: %v", err)
	}

	kb.AddWord("맛집", "NNG", 0)
	k := kb.Build(0)

	body, _ := os.ReadFile("./testdata/keywords.txt")
	lines := bytes.Split(body, []byte("\n"))
	category := make(map[string]int, len(lines))
	for _, line := range lines {
		result, err := k.Analyze(string(line), 1, OptionsAnalyzeKiwiMatchAll)
		if err != nil {
			t.Fatalf("failed to analyze: %v", err)
		}

		token := result[0].Tokens[len(result[0].Tokens)-1]
		if token.Tag[0] == 'N' {
			category[token.Form]++
		}
	}

	for item, v := range category {
		if v >= 3 {
			t.Logf("%s : %d", item, v)
		}
	}
}
