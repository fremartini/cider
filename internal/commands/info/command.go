package info

import (
	"cider/internal/cidr"
	"cider/internal/ip"
	"cider/internal/list"
	"cider/internal/must"
	"cider/internal/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
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

	ipString := args[0]

	s := strings.Split("/", ipString)

	ip := ip.NewIp(s[0])

	host := int(must.Must(strconv.ParseInt(s[1], 10, 32)))

	block := cidr.NewBlock(ip, host)

	entries := []pair{
		{item1: "Address range", item2: fmt.Sprintf("%s - %s", block.NetworkAddress(), block.BroadcastAddress())},
		{item1: "Start of next block", item2: block.StartAddressOfNextBlock()},
		{item1: "Mask", item2: fmt.Sprintf("%s (%s)", fmt.Sprintf("/%v", block.Host), block.Mask())},
		{item1: "Addresses", item2: fmt.Sprintf("%v", block.AvailableHosts())},
		{item1: "Azure addresses", item2: block.AvailableAzureHosts()},
		{item1: "Binary", item2: block.NetworkPortionBinary()},
		{item1: "Decimal", item2: fmt.Sprintf("%v", block.ToDecimal())},
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
