package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	testsDir = "tests"
)

func TestMain(t *testing.T) {
	inputStrings, answerStrings := readTestFiles()

	for testNo, inputString := range inputStrings {
		inputs := strings.Split(inputString, "\n")[1:]
		answers := strings.Split(answerStrings[testNo], "\n")
		for i, input := range inputs {
			input = strings.TrimSpace(input)
			params := strings.Split(input, " ")
			got := boolString(checkDate(params))
			want := strings.TrimSpace(answers[i])
			if got != want {
				t.Fatalf("Test file no. %s: %s = %s, want %s", testNo, strings.TrimSpace(input), got, want)
			}
		}

	}

}

func TestIsLeap(t *testing.T) {
	cases := map[int]bool{
		2261: false,
		2109: false,
	}

	for year, want := range cases {
		got := isLeap(year)
		if want != got {
			t.Fatalf("%d = %v, want %v", year, got, want)
		}
	}
}

func readTestFiles() (inputs map[string]string, answers map[string]string) {
	inputs = map[string]string{}
	answers = map[string]string{}

	entries, _ := os.ReadDir(testsDir)
	for _, entry := range entries {
		idx, isAnswer := strings.CutSuffix(entry.Name(), ".a")
		bytes, _ := os.ReadFile(filepath.Join(testsDir, entry.Name()))
		if isAnswer {
			answers[idx] = string(bytes)
		} else {
			inputs[idx] = string(bytes)
		}
	}
	return inputs, answers
}
