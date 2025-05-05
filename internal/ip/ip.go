package ip

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	INT_SIZE = 32
)

type Ip struct {
	decimalParts []string
}

func NewIp(address string) *Ip {
	parts := strings.Split(address, ".")

	return &Ip{
		decimalParts: parts,
	}
}

func (ip *Ip) ToDecimal() int {
	base := 10
	octet1 := must(strconv.ParseUint(ip.decimalParts[0], base, INT_SIZE))
	octet2 := must(strconv.ParseUint(ip.decimalParts[1], base, INT_SIZE))
	octet3 := must(strconv.ParseUint(ip.decimalParts[2], base, INT_SIZE))
	octet4 := must(strconv.ParseUint(ip.decimalParts[3], base, INT_SIZE))

	// http://www.aboutmyip.com/AboutMyXApp/IP2Integer.jsp
	return int((octet1 * 16777216) + (octet2 * 65536) + (octet3 * 256) + octet4)
}

func (ip *Ip) Ip() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip.decimalParts[0], ip.decimalParts[1], ip.decimalParts[2], ip.decimalParts[3])
}

func must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}

	return x
}
