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
	inputStrings := readTestFiles()

	for testNo, inputString := range inputStrings {
		t.Logf("Starting test case %s", testNo)
		r := strings.NewReader(inputString)
		scanner := bufio.NewScanner(r)
		scanner.Scan()
		var setCount int
		fmt.Sscan(scanner.Text(), &setCount)
		w, err := os.Create(filepath.Join("tests", testNo+"answer"))
		defer w.Close()
		if err != nil {
			panic(err)
		}
		for i := 0; i < setCount; i++ {
			if i > 0 {
				w.WriteString("\n")
			}
			w.WriteString(process(scanner))
		}
	}

}

func readTestFiles() map[string]string {
	inputs := map[string]string{}

	entries, _ := os.ReadDir(testsDir)
	for _, entry := range entries {
		idx := entry.Name()
		if strings.HasSuffix(idx, ".a") || strings.HasSuffix(idx, "answer") {
			continue
		}
		bytes, _ := os.ReadFile(filepath.Join(testsDir, idx))
		inputs[idx] = string(bytes)
	}
	return inputs
}
