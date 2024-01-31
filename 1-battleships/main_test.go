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
			params := strings.Split(input, " ")
			got := checkShips(params)
			want := strings.TrimSpace(answers[i])
			if got != want {
				t.Errorf("Test file no. %s: %s = %s, want %s", testNo, strings.TrimSpace(input), got, want)
			}
		}
	}
}

func readTestFiles() (inputs map[string]string, answers map[string]string) {
	inputs = map[string]string{}
	answers = map[string]string{}

	entries, err := os.ReadDir(testsDir)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		idx, isAnswer := strings.CutSuffix(entry.Name(), ".a")
		bytes, err := os.ReadFile(filepath.Join(testsDir, entry.Name()))
		if err != nil {
			panic(err)
		}
		if isAnswer {
			answers[idx] = string(bytes)
		} else {
			inputs[idx] = string(bytes)
		}
	}
	return inputs, answers
}
