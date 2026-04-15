package main

import (
	"encoding/json/v2"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/bastjan/jsonnunit-go/jsonnetunit"
	"github.com/google/go-jsonnet"
)

var jpaths stringSliceFlag

func main() {
	flag.Var(&jpaths, "jpath", "TODO")
	flag.Var(&jpaths, "J", "TODO")

	flag.Parse()

	jpEnv := filepath.SplitList(os.Getenv("JSONNET_PATH"))
	slices.Reverse(jpaths)

	jvm := jsonnet.MakeVM()
	jvm.Importer(&jsonnet.FileImporter{
		JPaths: append(jpaths, jpEnv...),
	})

	testFile := flag.Arg(0)

	rawRes, err := jvm.EvaluateFile(testFile)
	if err != nil {
		panic(err)
	}

	var res jsonnetunit.TestCaseResult
	if err := json.Unmarshal([]byte(rawRes), &res); err != nil {
		panic(err)
	}

	if err := jsonnetunit.FormatTestCaseResult(os.Stdout, res); err != nil {
		panic(err)
	}
}

type stringSliceFlag []string

func (f stringSliceFlag) String() string {
	return fmt.Sprint([]string(f))
}

func (f *stringSliceFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}
