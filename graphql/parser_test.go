package graphql

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/charithe/scylla/graphql/symbols"
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name        string `json:"name"`
	Input       string `json:"input"`
	ExpectedVal string `json:"expected_val"`
	ExpectedErr string `json:"expected_err"`
}

func TestParser(t *testing.T) {
	dataFiles := []string{"testdata/selection_sets.yaml", "testdata/error_cases.yaml"}
	for _, dataFile := range dataFiles {
		testData, err := ioutil.ReadFile(dataFile)
		assert.NoError(t, err)

		var testCases []TestCase
		err = yaml.Unmarshal(testData, &testCases)
		assert.NoError(t, err)

		suiteName := strings.TrimSuffix(filepath.Base(dataFile), filepath.Ext(dataFile))
		t.Run(suiteName, func(t *testing.T) {
			//t.Parallel()
			suite(t, testCases)
		})
	}
}

func suite(t *testing.T, testCases []TestCase) {
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.Name, func(t *testing.T) {
			//t.Parallel()
			parser, err := NewParser(strings.NewReader(tc.Input))
			assert.NoError(t, err)
			result, err := parser.Start()
			if tc.ExpectedErr != "" {
				if result != nil {
					fmt.Println(symbols.Stringify(result))
				}
				assert.Error(t, err)
				assert.EqualError(t, err, tc.ExpectedErr)
			} else {
				assert.NoError(t, err)
				assert.JSONEq(t, tc.ExpectedVal, symbols.Stringify(result))
			}
		})
	}
}
