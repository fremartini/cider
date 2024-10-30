package in

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	INT_SIZE = 32
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Handle(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("command expects exactly 2 arguments")
	}

	fmt.Println(isInRange(args[0], args[1]))

	return nil
}

// https://stackoverflow.com/questions/9622967/how-to-see-if-an-ip-address-belongs-inside-of-a-range-of-ips-using-cidr-notation
func isInRange(ipAddress, CIDRmask string) bool {

	parts := strings.Split(CIDRmask, "/")

	IP_addr := ipToDecimal(ipAddress)
	CIDR_addr := ipToDecimal(parts[0])
	CIDR_mask := -1 << (INT_SIZE - must(strconv.ParseInt(parts[1], 10, INT_SIZE)))

	return (IP_addr & CIDR_mask) == (CIDR_addr & CIDR_mask)
}

// http://www.aboutmyip.com/AboutMyXApp/IP2Integer.jsp
func ipToDecimal(ip string) int {
	parts := strings.Split(ip, ".")

	base := 10
	octet1 := must(strconv.ParseInt(parts[0], base, INT_SIZE))
	octet2 := must(strconv.ParseInt(parts[1], base, INT_SIZE))
	octet3 := must(strconv.ParseInt(parts[2], base, INT_SIZE))
	octet4 := must(strconv.ParseInt(parts[3], base, INT_SIZE))

	return int((octet1 * 16777216) + (octet2 * 65536) + (octet3 * 256) + octet4)
}

func must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}

	return x
}
