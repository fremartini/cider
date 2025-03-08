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

func (h *handler) Handle(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("command expects exactly one argument")
	}

	ip := args[0]

	block := cidr.NewBlock(ip)

	entries := map[string]any{
		"Address range":       fmt.Sprintf("%s - %s", block.NetworkAddress(), block.BroadcastAddress()),
		"Start of next block": block.StartAddressOfNextBlock(),
		"Cidr":                fmt.Sprintf("/%v", block.HostPortion),
		"Subnet mask":         block.SubnetMask(),
		"Addresses":           block.AvailableHosts(),
		"Azure addresses":     block.AvailableAzureHosts(),
	}

	keys := []string{}
	for k := range entries {
		keys = append(keys, k)
	}

	widestTitle := list.Fold(keys, 0, func(e string, acc int) int {
		m := int(math.Max(float64(acc), float64(len(e))))

		return m
	})

	for k, v := range entries {
		fmt.Printf("%s : %v\n", utils.PadRight(k, ' ', widestTitle), v)
	}

	return nil
}
