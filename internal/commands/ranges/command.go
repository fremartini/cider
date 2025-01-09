package ranges

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	cidr, err := strconv.ParseInt(arg, 10, INT_SIZE)

	if err != nil {
		return fmt.Errorf("%s is not a valid integer", arg)
	}

	if cidr < 0 || cidr > INT_SIZE {
		return fmt.Errorf("%v is not a valid size - must be between 0 and 32", cidr)
	}

	block := calculateCIDRBlock(int(cidr))

	table := []*CIDRBlock{block}

	return printCIDRBlocks(table)
}

func calculateAllCIDRBlocks() []*CIDRBlock {
	blocks := []*CIDRBlock{}
	for cidr := 0; cidr < INT_SIZE+1; cidr++ {
		block := calculateCIDRBlock(cidr)

		blocks = append(blocks, block)
	}

	return blocks
}

func calculateCIDRBlock(cidr int) *CIDRBlock {
	numAddresses := calculateNumAddresses(INT_SIZE, cidr)
	mask := calculateSubnetMask(cidr)

	block := &CIDRBlock{
		NetworkPortion: uint(cidr),
		SubnetMask:     mask,
		AvailableHosts: numAddresses,
	}

	return block
}

func calculateSubnetMask(cidr int) string {
	ones := strings.Repeat("1", cidr)
	zeroes := strings.Repeat("0", INT_SIZE-cidr)

	mask := ones + zeroes

	base := 2
	octet1 := must(strconv.ParseInt(mask[0:8], base, INT_SIZE))
	octet2 := must(strconv.ParseInt(mask[8:16], base, INT_SIZE))
	octet3 := must(strconv.ParseInt(mask[16:24], base, INT_SIZE))
	octet4 := must(strconv.ParseInt(mask[24:32], base, INT_SIZE))

	return fmt.Sprintf("%v.%v.%v.%v", octet1, octet2, octet3, octet4)
}

func calculateNumAddresses(addressLength, prefixLength int) uint {
	numAddresses := math.Pow(2, float64(addressLength)-float64(prefixLength))

	return uint(numAddresses)
}

func printCIDRBlocks(blocks []*CIDRBlock) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "CIDR\tSubnet Mask\tAddresses\tAzure Addresses\n")
	for _, block := range blocks {

		availableAzureAddresses := "N/A"

		// Azure reserves the first four addresses and the last address, for a total of five IP addresses within each subnet
		// https://learn.microsoft.com/en-us/azure/virtual-network/virtual-networks-faq#are-there-any-restrictions-on-using-ip-addresses-within-these-subnets
		if block.AvailableHosts >= 5 {
			availableAzureAddresses = fmt.Sprintf("%v", block.AvailableHosts-5)
		}

		fmt.Fprintf(w, "/%v\t%s\t%v\t%s\n", block.NetworkPortion, block.SubnetMask, block.AvailableHosts, availableAzureAddresses)
	}

	return w.Flush()
}

func must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}

	return x
}
