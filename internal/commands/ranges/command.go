package ranges

import (
	"cider/internal/cidr"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

const (
	INT_SIZE = 32
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (*handler) Handle(arg string) error {
	// no args
	if arg == "" {
		table := calculateAllCIDRBlocks()

		return printCIDRBlocks(table)
	}

	// argument was given - try to parse it
	hostPortion, err := strconv.ParseInt(arg, 10, INT_SIZE)

	if err != nil {
		return fmt.Errorf("%s is not a valid integer", arg)
	}

	if hostPortion < 0 || hostPortion > INT_SIZE {
		return fmt.Errorf("%v is not a valid size - must be between 0 and 32", hostPortion)
	}

	block := calculateCIDRBlock(int(hostPortion))

	table := []*cidr.CIDRBlock{block}

	return printCIDRBlocks(table)
}

func calculateAllCIDRBlocks() []*cidr.CIDRBlock {
	blocks := []*cidr.CIDRBlock{}
	for i := 0; i < INT_SIZE+1; i++ {
		block := calculateCIDRBlock(i)

		blocks = append(blocks, block)
	}

	return blocks
}

func calculateCIDRBlock(hostPortion int) *cidr.CIDRBlock {
	return cidr.NewBlock(fmt.Sprintf("10.0.0.0/%v", hostPortion))
}

func printCIDRBlocks(blocks []*cidr.CIDRBlock) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "Cidr\tSubnet mask\tAddresses\tAzure addresses\n")
	for _, block := range blocks {

		availableAzureAddresses := "N/A"

		// Azure reserves the first four addresses and the last address, for a total of five IP addresses within each subnet
		// https://learn.microsoft.com/en-us/azure/virtual-network/virtual-networks-faq#are-there-any-restrictions-on-using-ip-addresses-within-these-subnets
		if block.AvailableHosts() >= 5 {
			availableAzureAddresses = fmt.Sprintf("%v", block.AvailableHosts()-5)
		}

		fmt.Fprintf(w, "/%v\t%s\t%v\t%s\n", block.HostPortion, block.SubnetMask(), block.AvailableHosts(), availableAzureAddresses)
	}

	return w.Flush()
}
