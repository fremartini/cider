package cmd_test

import "testing"

func TestInfo(t *testing.T) {
	testCases := []testCase{
		{
			name:      "missing args",
			input:     []string{"cider", "info"},
			stdOutput: "",
			stdErr:    "command expects exactly one argument",
		},
		{
			name:      "prints info",
			input:     []string{"cider", "info", "10.0.64.0/18"},
			stdOutput: "Address range       : 10.0.64.0 - 10.0.127.255\nStart of next block : 10.0.128.0\nMask                : /18 (255.255.192.0)\nAddresses           : 16384\nAzure addresses     : 16379\nBinary              : 00001010.00000000.01000000.00000000\nDecimal             : 167788544\n",
			stdErr:    "",
		},
	}
	executeTestCases(t, testCases)
}
