package in

import (
	"cider/internal/cidr"
	"cider/internal/list"
	"fmt"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (h *handler) Handle(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("command expects at least 2 arguments")
	}

	ip := args[0]

	ranges := list.Map(args[1:], func(i string) *cidr.CIDRBlock {
		return cidr.NewBlock(i)
	})

	blocksInRange := list.Filter(ranges, func(cidr *cidr.CIDRBlock) bool {
		return cidr.Contains(ip)
	})

	if len(blocksInRange) == 0 {
		fmt.Printf("%s is not in any of the provided ranges\n", ip)
		return nil
	}

	for _, block := range blocksInRange {
		fmt.Printf("%s/%d\n", block.Network, block.HostPortion)
	}

	return nil
}
