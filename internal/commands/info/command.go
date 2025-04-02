package info

import (
	"cider/internal/cidr"
	"cider/internal/list"
	"cider/internal/utils"
	"fmt"
	"math"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

type pair struct {
	item1, item2 string
}

func (h *handler) Handle(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("command expects exactly one argument")
	}

	ip := args[0]

	block := cidr.NewBlock(ip)

	entries := []pair{
		{item1: "Address range", item2: fmt.Sprintf("%s - %s", block.NetworkAddress(), block.BroadcastAddress())},
		{item1: "Start of next block", item2: block.StartAddressOfNextBlock()},
		{item1: "Cidr", item2: fmt.Sprintf("/%v", block.HostPortion)},
		{item1: "Subnet mask", item2: block.SubnetMask()},
		{item1: "Addresses", item2: fmt.Sprintf("%v", block.AvailableHosts())},
		{item1: "Azure addresses", item2: block.AvailableAzureHosts()},
	}

	printOutput(entries)

	return nil
}

func printOutput(entries []pair) {
	keys := []string{}
	for _, pair := range entries {
		keys = append(keys, pair.item1)
	}

	longestTitle := list.Fold(keys, 0, func(title string, acc int) int {
		return int(math.Max(float64(acc), float64(len(title))))
	})

	for _, pair := range entries {
		fmt.Printf("%s : %v\n", utils.PadRight(pair.item1, ' ', longestTitle), pair.item2)
	}
}
