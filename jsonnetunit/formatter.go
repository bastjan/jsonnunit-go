package jsonnetunit

import (
	"fmt"
	"io"
	"strings"
)

func FormatTestCaseResult(w io.Writer, result TestCaseResult) error {
	return formatTestCase(w, TestCase{Result: result}, -1)
}

func formatTestCase(w io.Writer, result TestCase, level int) error {
	if level > -1 {
		fmt.Printf("CASE %s%s\n", indent(level), result.Name)
	}
	for _, test := range result.Result.Tests {
		header := "FAIL"
		if test.Result.IsPass() {
			header = "PASS"
		}
		fmt.Printf("%s%s  %s\n", indent(level+1), header, test.Name)
	}
	for _, sub := range result.Result.TestCases {
		if err := formatTestCase(w, sub, level+1); err != nil {
			return err
		}
	}
	return nil
}

func indent(level int) string {
	return strings.Repeat("  ", level)
}
