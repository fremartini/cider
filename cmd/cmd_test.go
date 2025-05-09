package cmd_test

import (
	"bytes"
	"cider/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name      string
	input     []string
	stdOutput string
	stdErr    string
}

func executeTestCases(t *testing.T, testCases []testCase) {
	executeTestCasesWithCustomAssertion(
		t,
		testCases,
		func(t *testing.T, tc testCase, stdout, stderr string) {
			assert.Equal(t, tc.stdOutput, stdout, "std output not as expected")
			assert.Equal(t, tc.stdErr, stderr, "err output not as expected")
		},
	)
}

func executeTestCasesWithCustomAssertion(
	t *testing.T,
	testCases []testCase,
	assertion func(t *testing.T, tc testCase, stdout, stderr string),
) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stdOut := new(bytes.Buffer)
			stdErr := new(bytes.Buffer)

			cmd.Execute(stdOut, stdErr, tc.input)

			assertion(t, tc, stdOut.String(), stdErr.String())
		})
	}
}
