package cmd_test

import (
	"testing"
)

func TestVersion(t *testing.T) {
	testCases := []testCase{
		{
			name:      "returns version",
			input:     []string{"cider", "version"},
			stdOutput: "version\n",
			stdErr:    "",
		},
	}
	executeTestCases(t, testCases)
}
