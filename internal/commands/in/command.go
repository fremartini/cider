package in

import (
	"cider/internal/cidr"
	"cider/internal/ip"
	"cider/internal/list"
	"cider/internal/must"
	"fmt"
	"strconv"
	"strings"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Handle(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("command expects at least 2 arguments")
	}

	ranges := list.Map(args[1:], func(i string) *cidr.CidrBlock {
		s := strings.Split(i, "/")

		ip := ip.NewIp(s[0])

		host := int(must.Must(strconv.ParseInt(s[1], 10, 32)))

		return cidr.NewBlock(ip, host)
	})

	ipOrRange := args[0]

	var blocksInRange []*cidr.CidrBlock
	if strings.Contains(ipOrRange, "/") {
		s := strings.Split(ipOrRange, "/")

		ip := ip.NewIp(s[0])

		host := int(must.Must(strconv.ParseInt(s[1], 10, 32)))

		block := cidr.NewBlock(ip, host)

		blocksInRange = list.Filter(ranges, func(cidr *cidr.CidrBlock) bool {
			return cidr.ContainsRange(block)
		})
	} else {
		ip := ip.NewIp(ipOrRange)
		blocksInRange = list.Filter(ranges, func(cidr *cidr.CidrBlock) bool {
			return cidr.ContainsIp(ip)
		})
	}

	if len(blocksInRange) == 0 {
		fmt.Printf("%s is not in any of the provided ranges\n", ipOrRange)
		return nil
	}

	for _, block := range blocksInRange {
		fmt.Printf("%s/%d\n", block.Ip, block.Host)
	}

	return nil
}
