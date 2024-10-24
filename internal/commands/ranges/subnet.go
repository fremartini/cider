package ranges

import "strconv"

type subnet struct {
	CIDR uint   // 0
	Mask string // 0.0.0.0
}

func SubnetFromText(s []string) *subnet {
	cidr, err := strconv.ParseUint(s[0], 10, 64)

	if err != nil {
		panic(err)
	}

	return &subnet{
		CIDR: uint(cidr),
		Mask: s[1],
	}
}
