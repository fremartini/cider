package cmd_test

import (
	"testing"
)

func TestRanges(t *testing.T) {
	testCases := []testCase{
		{
			name:      "single range",
			input:     []string{"cider", "ranges", "21"},
			stdOutput: "Cidr Mask          Addresses Azure addresses\n/21  255.255.248.0 2048      2043\n",
			stdErr:    "",
		},
		{
			name:      "single range shorthand",
			input:     []string{"cider", "r", "21"},
			stdOutput: "Cidr Mask          Addresses Azure addresses\n/21  255.255.248.0 2048      2043\n",
			stdErr:    "",
		},
		{
			name:      "single range no command",
			input:     []string{"cider", "21"},
			stdOutput: "Cidr Mask          Addresses Azure addresses\n/21  255.255.248.0 2048      2043\n",
			stdErr:    "",
		},
		{
			name:      "all ranges",
			input:     []string{"cider", "ranges"},
			stdOutput: "Cidr Mask            Addresses  Azure addresses\n/0   0.0.0.0         4294967296 4294967291\n/1   128.0.0.0       2147483648 2147483643\n/2   192.0.0.0       1073741824 1073741819\n/3   224.0.0.0       536870912  536870907\n/4   240.0.0.0       268435456  268435451\n/5   248.0.0.0       134217728  134217723\n/6   252.0.0.0       67108864   67108859\n/7   254.0.0.0       33554432   33554427\n/8   255.0.0.0       16777216   16777211\n/9   255.128.0.0     8388608    8388603\n/10  255.192.0.0     4194304    4194299\n/11  255.224.0.0     2097152    2097147\n/12  255.240.0.0     1048576    1048571\n/13  255.248.0.0     524288     524283\n/14  255.252.0.0     262144     262139\n/15  255.254.0.0     131072     131067\n/16  255.255.0.0     65536      65531\n/17  255.255.128.0   32768      32763\n/18  255.255.192.0   16384      16379\n/19  255.255.224.0   8192       8187\n/20  255.255.240.0   4096       4091\n/21  255.255.248.0   2048       2043\n/22  255.255.252.0   1024       1019\n/23  255.255.254.0   512        507\n/24  255.255.255.0   256        251\n/25  255.255.255.128 128        123\n/26  255.255.255.192 64         59\n/27  255.255.255.224 32         27\n/28  255.255.255.240 16         11\n/29  255.255.255.248 8          3\n/30  255.255.255.252 4          N/A\n/31  255.255.255.254 2          N/A\n/32  255.255.255.255 1          N/A\n",
			stdErr:    "",
		},
		{
			name:      "string as range",
			input:     []string{"cider", "ranges", "not a range"},
			stdOutput: "",
			stdErr:    "not a range is not a valid integer",
		},
	}
	executeTestCases(t, testCases)
}
