package cuddlevar_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	cuddlevar "github.com/yusei-wy/go-cuddlevar"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, cuddlevar.Analyzer, "a")
}
