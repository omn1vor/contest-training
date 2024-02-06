package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	testsDir = "tests"
)

func TestMain(t *testing.T) {
	inputStrings, resultStrings := readTestFiles()

	for testNo, inputString := range inputStrings {
		t.Logf("Starting test case %s", testNo)
		r := strings.NewReader(inputString)
		scanner := bufio.NewScanner(r)
		scanner.Scan()
		var setCount int
		fmt.Sscan(scanner.Text(), &setCount)
		w := strings.Builder{}
		for i := 0; i < setCount; i++ {
			if i > 0 {
				w.WriteString("\n")
			}
			w.WriteString(process(scanner))
		}
		got := w.String()
		want := resultStrings[testNo]
		os.WriteFile(filepath.Join("tests", testNo+"answer"), []byte(got), 0666)
		if w.String() != want {
			t.Fatalf("Test no. %s. got:\n%s\nwant:\n%s", testNo, got, want)
		}
	}

}

func readTestFiles() (inputs map[string]string, results map[string]string) {
	inputs = map[string]string{}
	results = map[string]string{}
	entries, _ := os.ReadDir(testsDir)
	for _, entry := range entries {
		idx := entry.Name()
		if strings.HasSuffix(idx, "answer") {
			continue
		}
		if idx, found := strings.CutSuffix(entry.Name(), ".a"); found {
			bytes, _ := os.ReadFile(filepath.Join(testsDir, entry.Name()))
			results[idx] = string(bytes)
		} else {
			bytes, _ := os.ReadFile(filepath.Join(testsDir, idx))
			inputs[idx] = string(bytes)
		}
	}
	return inputs, results
}
