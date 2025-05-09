package cmd_test

import "testing"

func TestIn(t *testing.T) {
	testCases := []testCase{
		{
			name:      "missing args",
			input:     []string{"cider", "in"},
			stdOutput: "",
			stdErr:    "command expects at least 2 arguments",
		},
		{
			name:      "ip inside range",
			input:     []string{"cider", "in", "10.164.214.32", "10.164.214.0/26"},
			stdOutput: "10.164.214.0/26\n",
			stdErr:    "",
		},
		{
			name:      "ip outside range",
			input:     []string{"cider", "in", "10.164.215.0", "10.164.214.0/26"},
			stdOutput: "10.164.215.0 is not in any of the provided ranges\n",
			stdErr:    "",
		},
	}
	executeTestCases(t, testCases)
}
