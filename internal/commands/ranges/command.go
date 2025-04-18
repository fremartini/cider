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
		table := calculateAllCidrBlocks()

		return printCidrBlocks(table)
	}

	// argument was given. Try to parse it
	hostPortion, err := strconv.ParseInt(arg, 10, INT_SIZE)

	if err != nil {
		return fmt.Errorf("%s is not a valid integer", arg)
	}

	if hostPortion < 0 || hostPortion > INT_SIZE {
		return fmt.Errorf("%v is not a valid size - must be between 0 and %d", hostPortion, INT_SIZE)
	}

	block := defaultCidrBlockFromHostPortion(int(hostPortion))

	blocks := []*cidr.CIDRBlock{block}

	return printCidrBlocks(blocks)
}

func calculateAllCidrBlocks() []*cidr.CIDRBlock {
	blocks := []*cidr.CIDRBlock{}
	for i := range INT_SIZE + 1 {
		block := defaultCidrBlockFromHostPortion(i)

		blocks = append(blocks, block)
	}

	return blocks
}

func defaultCidrBlockFromHostPortion(hostPortion int) *cidr.CIDRBlock {
	return cidr.NewBlock(fmt.Sprintf("10.0.0.0/%v", hostPortion))
}

func printCidrBlocks(blocks []*cidr.CIDRBlock) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "Cidr\tSubnet mask\tAddresses\tAzure addresses\n")
	for _, block := range blocks {
		fmt.Fprintf(w, "/%v\t%s\t%v\t%s\n", block.HostPortion, block.SubnetMask(), block.AvailableHosts(), block.AvailableAzureHosts())
	}

	return w.Flush()
}
