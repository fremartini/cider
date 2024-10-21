package ranges

import "strconv"

type subnet struct {
	CIDR         string // /0
	Mask         string // 0.0.0.0
	NumAddresses uint   // 4294967296
	Wildcard     string // 255.255.255.255
}

func SubnetFromText(s []string) *subnet {
	numAddresses, err := strconv.ParseUint(s[2], 10, 64)

	if err != nil {
		panic(err)
	}

	return &subnet{
		CIDR:         s[0],
		Mask:         s[1],
		NumAddresses: uint(numAddresses),
		Wildcard:     s[3],
	}
}
