package cidr_test

import (
	"cider/internal/cidr"
	"reflect"
	"testing"
)

func Test_SubnetSuccess(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		sizes    []int
		expected []string
	}{
		"two even subnets":           {input: cidr.NewBlock("10.0.0.0/16"), sizes: []int{17, 17}, expected: []string{"10.0.0.0/17", "10.0.128.0/17"}},
		"subnets of different sizes": {input: cidr.NewBlock("10.0.0.0/16"), sizes: []int{18, 17, 20}, expected: []string{"10.0.0.0/17", "10.0.128.0/18", "10.0.192.0/20"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, _ := test.input.Subnet(test.sizes)

			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("%s returns correct subnet: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_SubnetFailure(t *testing.T) {
	tests := map[string]struct {
		input *cidr.CIDRBlock
		sizes []int
	}{
		"invalid configuration": {input: cidr.NewBlock("10.0.0.0/16"), sizes: []int{16, 16}},
		"insufficient space":    {input: cidr.NewBlock("10.0.0.0/30"), sizes: []int{32, 32, 32, 32, 32}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := test.input.Subnet(test.sizes)

			if err == nil {
				t.Fatalf("%s expected error: got %v", name, actual)
			}
		})
	}
}

func Test_Contains(t *testing.T) {
	tests := map[string]struct {
		ip       string
		ipRange  *cidr.CIDRBlock
		expected bool
	}{
		"ip inside range":  {ip: "10.50.30.7", ipRange: cidr.NewBlock("10.0.0.0/8"), expected: true},
		"ip outside range": {ip: "10.50.30.7", ipRange: cidr.NewBlock("10.0.0.0/28"), expected: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.ipRange.Contains(test.ip)

			if actual != test.expected {
				t.Fatalf("%s: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_NetworkPortionBinary(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		expected string
	}{
		"/0": {input: cidr.NewBlock("10.0.0.0/0"), expected: "00001010.00000000.00000000.00000000"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.input.NetworkPortionBinary()

			if actual != test.expected {
				t.Fatalf("%s returns correct binary representation: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_SubnetMask(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		expected string
	}{
		"/0":  {input: cidr.NewBlock("10.0.0.0/0"), expected: "0.0.0.0"},
		"/1":  {input: cidr.NewBlock("10.0.0.0/1"), expected: "128.0.0.0"},
		"/2":  {input: cidr.NewBlock("10.0.0.0/2"), expected: "192.0.0.0"},
		"/3":  {input: cidr.NewBlock("10.0.0.0/3"), expected: "224.0.0.0"},
		"/4":  {input: cidr.NewBlock("10.0.0.0/4"), expected: "240.0.0.0"},
		"/5":  {input: cidr.NewBlock("10.0.0.0/5"), expected: "248.0.0.0"},
		"/6":  {input: cidr.NewBlock("10.0.0.0/6"), expected: "252.0.0.0"},
		"/7":  {input: cidr.NewBlock("10.0.0.0/7"), expected: "254.0.0.0"},
		"/8":  {input: cidr.NewBlock("10.0.0.0/8"), expected: "255.0.0.0"},
		"/9":  {input: cidr.NewBlock("10.0.0.0/9"), expected: "255.128.0.0"},
		"/10": {input: cidr.NewBlock("10.0.0.0/10"), expected: "255.192.0.0"},
		"/11": {input: cidr.NewBlock("10.0.0.0/11"), expected: "255.224.0.0"},
		"/12": {input: cidr.NewBlock("10.0.0.0/12"), expected: "255.240.0.0"},
		"/13": {input: cidr.NewBlock("10.0.0.0/13"), expected: "255.248.0.0"},
		"/14": {input: cidr.NewBlock("10.0.0.0/14"), expected: "255.252.0.0"},
		"/15": {input: cidr.NewBlock("10.0.0.0/15"), expected: "255.254.0.0"},
		"/16": {input: cidr.NewBlock("10.0.0.0/16"), expected: "255.255.0.0"},
		"/17": {input: cidr.NewBlock("10.0.0.0/17"), expected: "255.255.128.0"},
		"/18": {input: cidr.NewBlock("10.0.0.0/18"), expected: "255.255.192.0"},
		"/19": {input: cidr.NewBlock("10.0.0.0/19"), expected: "255.255.224.0"},
		"/20": {input: cidr.NewBlock("10.0.0.0/20"), expected: "255.255.240.0"},
		"/21": {input: cidr.NewBlock("10.0.0.0/21"), expected: "255.255.248.0"},
		"/22": {input: cidr.NewBlock("10.0.0.0/22"), expected: "255.255.252.0"},
		"/23": {input: cidr.NewBlock("10.0.0.0/23"), expected: "255.255.254.0"},
		"/24": {input: cidr.NewBlock("10.0.0.0/24"), expected: "255.255.255.0"},
		"/25": {input: cidr.NewBlock("10.0.0.0/25"), expected: "255.255.255.128"},
		"/26": {input: cidr.NewBlock("10.0.0.0/26"), expected: "255.255.255.192"},
		"/27": {input: cidr.NewBlock("10.0.0.0/27"), expected: "255.255.255.224"},
		"/28": {input: cidr.NewBlock("10.0.0.0/28"), expected: "255.255.255.240"},
		"/29": {input: cidr.NewBlock("10.0.0.0/29"), expected: "255.255.255.248"},
		"/30": {input: cidr.NewBlock("10.0.0.0/30"), expected: "255.255.255.252"},
		"/31": {input: cidr.NewBlock("10.0.0.0/31"), expected: "255.255.255.254"},
		"/32": {input: cidr.NewBlock("10.0.0.0/32"), expected: "255.255.255.255"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.input.SubnetMask()

			if actual != test.expected {
				t.Fatalf("%s returns correct subnet mask: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_AvailableHosts(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		expected uint
	}{
		"/0":  {input: cidr.NewBlock("10.0.0.0/0"), expected: 4294967296},
		"/1":  {input: cidr.NewBlock("10.0.0.0/1"), expected: 2147483648},
		"/2":  {input: cidr.NewBlock("10.0.0.0/2"), expected: 1073741824},
		"/3":  {input: cidr.NewBlock("10.0.0.0/3"), expected: 536870912},
		"/4":  {input: cidr.NewBlock("10.0.0.0/4"), expected: 268435456},
		"/5":  {input: cidr.NewBlock("10.0.0.0/5"), expected: 134217728},
		"/6":  {input: cidr.NewBlock("10.0.0.0/6"), expected: 67108864},
		"/7":  {input: cidr.NewBlock("10.0.0.0/7"), expected: 33554432},
		"/8":  {input: cidr.NewBlock("10.0.0.0/8"), expected: 16777216},
		"/9":  {input: cidr.NewBlock("10.0.0.0/9"), expected: 8388608},
		"/10": {input: cidr.NewBlock("10.0.0.0/10"), expected: 4194304},
		"/11": {input: cidr.NewBlock("10.0.0.0/11"), expected: 2097152},
		"/12": {input: cidr.NewBlock("10.0.0.0/12"), expected: 1048576},
		"/13": {input: cidr.NewBlock("10.0.0.0/13"), expected: 524288},
		"/14": {input: cidr.NewBlock("10.0.0.0/14"), expected: 262144},
		"/15": {input: cidr.NewBlock("10.0.0.0/15"), expected: 131072},
		"/16": {input: cidr.NewBlock("10.0.0.0/16"), expected: 65536},
		"/17": {input: cidr.NewBlock("10.0.0.0/17"), expected: 32768},
		"/18": {input: cidr.NewBlock("10.0.0.0/18"), expected: 16384},
		"/19": {input: cidr.NewBlock("10.0.0.0/19"), expected: 8192},
		"/20": {input: cidr.NewBlock("10.0.0.0/20"), expected: 4096},
		"/21": {input: cidr.NewBlock("10.0.0.0/21"), expected: 2048},
		"/22": {input: cidr.NewBlock("10.0.0.0/22"), expected: 1024},
		"/23": {input: cidr.NewBlock("10.0.0.0/23"), expected: 512},
		"/24": {input: cidr.NewBlock("10.0.0.0/24"), expected: 256},
		"/25": {input: cidr.NewBlock("10.0.0.0/25"), expected: 128},
		"/26": {input: cidr.NewBlock("10.0.0.0/26"), expected: 64},
		"/27": {input: cidr.NewBlock("10.0.0.0/27"), expected: 32},
		"/28": {input: cidr.NewBlock("10.0.0.0/28"), expected: 16},
		"/29": {input: cidr.NewBlock("10.0.0.0/29"), expected: 8},
		"/30": {input: cidr.NewBlock("10.0.0.0/30"), expected: 4},
		"/31": {input: cidr.NewBlock("10.0.0.0/31"), expected: 2},
		"/32": {input: cidr.NewBlock("10.0.0.0/32"), expected: 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.input.AvailableHosts()

			if actual != test.expected {
				t.Fatalf("%s returns correct number of addresses: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_StartAddressOfNextBlock(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		expected string
	}{
		"/17": {input: cidr.NewBlock("10.0.0.0/17"), expected: "10.0.128.0"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.input.StartAddressOfNextBlock()

			if actual != test.expected {
				t.Fatalf("%s returns next block address: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_NetworkAddress(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		expected string
	}{
		"/26": {input: cidr.NewBlock("192.168.33.64/26"), expected: "192.168.33.64"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.input.NetworkAddress()

			if actual != test.expected {
				t.Fatalf("%s returns correct network address: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}

func Test_BroadcastAddress(t *testing.T) {
	tests := map[string]struct {
		input    *cidr.CIDRBlock
		expected string
	}{
		"/26": {input: cidr.NewBlock("192.168.33.64/26"), expected: "192.168.33.127"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := test.input.BroadcastAddress()

			if actual != test.expected {
				t.Fatalf("%s returns correct broadcast address: got %v expected %v", name, actual, test.expected)
			}
		})
	}
}
