package in

import (
	"cider/internal/cidr"
	"cider/internal/ip"
	"cider/internal/list"
	"fmt"
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

	ranges := list.Map(args[1:], func(i string) *cidr.CIDRBlock {
		s := strings.Split("/", i)

		ip := ip.NewIp(s[0])

		return cidr.NewBlock(ip, s[1])
	})

	ipOrRange := args[0]

	blocksInRange := list.Filter(ranges, func(cidr *cidr.CIDRBlock) bool {
		return cidr.Contains(ipOrRange)
	})

	if len(blocksInRange) == 0 {
		fmt.Printf("%s is not in any of the provided ranges\n", ipOrRange)
		return nil
	}

	for _, block := range blocksInRange {
		fmt.Printf("%s/%d\n", block.Ip, block.Host)
	}

	return nil
}
