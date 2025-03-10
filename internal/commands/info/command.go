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
	a, b string
}

func (h *handler) Handle(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("command expects exactly one argument")
	}

	ip := args[0]

	block := cidr.NewBlock(ip)

	entries := []pair{
		{a: "Address range", b: fmt.Sprintf("%s - %s", block.NetworkAddress(), block.BroadcastAddress())},
		{a: "Start of next block", b: block.StartAddressOfNextBlock()},
		{a: "Cidr", b: fmt.Sprintf("/%v", block.HostPortion)},
		{a: "Subnet mask", b: block.SubnetMask()},
		{a: "Addresses", b: fmt.Sprintf("%v", block.AvailableHosts())},
		{a: "Azure addresses", b: block.AvailableAzureHosts()},
	}

	printOutput(entries)

	return nil
}

func printOutput(entries []pair) {
	keys := []string{}
	for _, pair := range entries {
		keys = append(keys, pair.a)
	}

	longestTitle := list.Fold(keys, 0, func(title string, acc int) int {
		return int(math.Max(float64(acc), float64(len(title))))
	})

	for _, pair := range entries {
		fmt.Printf("%s : %v\n", utils.PadRight(pair.a, ' ', longestTitle), pair.b)
	}
}
