package jsonnetunit_test

import (
	"encoding/json/v2"
	"os"
	"testing"

	"github.com/bastjan/jsonnunit-go/jsonnetunit"
)

func Test_TestCase_Unmarshal(t *testing.T) {
	f, err := os.Open("./testdata/unmarshal.json")
	if err != nil {
		t.Fatalf("failed to open test case file: %v", err)
	}
	defer f.Close()

	var tc jsonnetunit.TestCaseResult
	if err := json.UnmarshalRead(f, &tc); err != nil {
		t.Fatalf("failed to unmarshal test case: %v", err)
	}

	if len(tc.TestCases) != 2 {
		t.Fatalf("expected 2 test cases, got %d", len(tc.TestCases))
	}
	if tc.TestCases[0].Name != "TestCase1" {
		t.Errorf("expected Name to be 'TestCase1', got '%s'", tc.TestCases[0].Name)
	}
	if len(tc.TestCases[0].Result.Tests) != 2 {
		t.Errorf("expected 2 tests in TestCase1, got %d", len(tc.TestCases[0].Result.Tests))
	} else {
		if tc.TestCases[0].Result.Tests[0].Name != "Test with result array" {
			t.Errorf("expected Name to be 'Test with result array', got '%s'", tc.TestCases[0].Result.Tests[0].Name)
		}
		if len(tc.TestCases[0].Result.Tests[0].Result) != 2 {
			t.Errorf("expected 2 results in 'Test with result array', got %d", len(tc.TestCases[0].Result.Tests[0].Result))
		} else {
			if tc.TestCases[0].Result.Tests[0].Result[0] != jsonnetunit.Pass {
				t.Errorf("expected first result to be Pass, got %s", tc.TestCases[0].Result.Tests[0].Result[0])
			}
			if tc.TestCases[0].Result.Tests[0].Result[1] != jsonnetunit.Fail {
				t.Errorf("expected second result to be Fail, got %s", tc.TestCases[0].Result.Tests[0].Result[1])
			}
			if tc.TestCases[0].Result.Tests[1].Name != "Test with result string" {
				t.Errorf("expected Name to be 'Test with result string', got '%s'", tc.TestCases[0].Result.Tests[1].Name)
			}
		}
		if len(tc.TestCases[0].Result.Tests[1].Result) != 1 {
			t.Errorf("expected 1 result in 'Test with result string', got %d", len(tc.TestCases[0].Result.Tests[1].Result))
		} else if tc.TestCases[0].Result.Tests[1].Result[0] != jsonnetunit.Pass {
			t.Errorf("expected result to be Pass, got %s", tc.TestCases[0].Result.Tests[1].Result[0])
		}
	}

	if tc.TestCases[1].Name != "TestCase2" {
		t.Errorf("expected Name to be 'TestCase2', got '%s'", tc.TestCases[1].Name)
	}
	if len(tc.TestCases[1].Result.TestCases) != 1 {
		t.Errorf("expected 1 sub-test case in TestCase2, got %d", len(tc.TestCases[1].Result.TestCases))
	} else {
		if tc.TestCases[1].Result.TestCases[0].Name != "TestCase2/1" {
			t.Errorf("expected Name to be 'TestCase2/1', got '%s'", tc.TestCases[1].Result.TestCases[0].Name)
		}
		if len(tc.TestCases[1].Result.TestCases[0].Result.Tests) != 1 {
			t.Errorf("expected 1 test in TestCase2/1, got %d", len(tc.TestCases[1].Result.TestCases[0].Result.Tests))
		} else if tc.TestCases[1].Result.TestCases[0].Result.Tests[0].Name != "Test with result null" {
			t.Errorf("expected Name to be 'Test with result null', got '%s'", tc.TestCases[1].Result.TestCases[0].Result.Tests[0].Name)
		}
		if len(tc.TestCases[1].Result.TestCases[0].Result.Tests[0].Result) != 0 {
			t.Errorf("expected 0 results in 'Test with result null', got %d", len(tc.TestCases[1].Result.TestCases[0].Result.Tests[0].Result))
		}
	}
}

func Test_TestResults_IsPass(t *testing.T) {
	tests := []struct {
		in       jsonnetunit.TestResults
		expected bool
	}{
		{jsonnetunit.TestResults{jsonnetunit.Pass, jsonnetunit.Pass}, true},
		{jsonnetunit.TestResults{jsonnetunit.Pass}, true},
		{jsonnetunit.TestResults{jsonnetunit.Pass, jsonnetunit.Fail}, false},
		{jsonnetunit.TestResults{jsonnetunit.Fail}, false},
		{jsonnetunit.TestResults{}, true},
	}

	for _, tt := range tests {
		if got := tt.in.IsPass(); got != tt.expected {
			t.Errorf("TestResults.IsPass() = %v, want %v", got, tt.expected)
		}
	}
}
