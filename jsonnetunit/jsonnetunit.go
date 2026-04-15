package jsonnetunit

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
)

type TestResult string

const (
	Pass TestResult = "JSONNUNIT--passed"
	Fail TestResult = "JSONNUNIT--fail"
)

func (tr *TestResult) UnmarshalJSONFrom(d *jsontext.Decoder) error {
	var s string
	if err := json.UnmarshalDecode(d, &s); err != nil {
		return err
	}
	switch s {
	case string(Pass):
		*tr = Pass
	case string(Fail):
		*tr = Fail
	default:
		return fmt.Errorf("invalid TestResult: %s", s)
	}
	return nil
}

type TestResults []TestResult

func (tr TestResults) IsPass() bool {
	for _, r := range tr {
		if r != Pass {
			return false
		}
	}
	return true
}

func (a *TestResults) UnmarshalJSONFrom(d *jsontext.Decoder) error {
	switch d.PeekKind() {
	case 'n':
		*a = nil
	case '[':
		var s []TestResult
		if err := json.UnmarshalDecode(d, &s); err != nil {
			return err
		}
		*a = s
	case '"':
		var s TestResult
		if err := json.UnmarshalDecode(d, &s); err != nil {
			return err
		}
		*a = []TestResult{s}
	default:
		return fmt.Errorf("invalid TestResults: expected null, array or string")
	}
	return nil
}

type TestCase struct {
	Name       string         `json:"name"`
	Result TestCaseResult `json:"tests"`
}

type TestCaseResult struct {
	TestCases []TestCase `json:"testcases"`
	Tests     []Test     `json:"tests"`
}

type Test struct {
	Name   string      `json:"name"`
	Result TestResults `json:"result"`
}
