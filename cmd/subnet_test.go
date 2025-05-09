package cmd_test

import "testing"

func TestSubnet(t *testing.T) {
	testCases := []testCase{
		{
			name:      "missing args",
			input:     []string{"cider", "subnet"},
			stdOutput: "",
			stdErr:    "command expects at least 2 arguments",
		},
		{
			name:      "single arg",
			input:     []string{"cider", "subnet", "10.163.0.0/16"},
			stdOutput: "",
			stdErr:    "command expects at least 2 arguments",
		},
		{
			name:      "valid range",
			input:     []string{"cider", "subnet", "10.163.0.0/16", "19", "19", "19", "19"},
			stdOutput: "10.163.0.0/19\n10.163.32.0/19\n10.163.64.0/19\n10.163.96.0/19\n",
			stdErr:    "",
		},
	}
	executeTestCases(t, testCases)
}
