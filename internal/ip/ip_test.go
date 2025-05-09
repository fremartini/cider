package ip_test

import (
	"cider/internal/ip"
	"testing"
)

func Test_ToDecimal(t *testing.T) {
	tests := map[string]struct {
		ip       *ip.Ip
		expected int
	}{
		"decimal": {ip: ip.NewIp("10.0.0.5"), expected: 167772165},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.ip.ToDecimal()

			if actual != test.expected {
				t.Fatalf("%s: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}
