package ranges

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (*handler) Handle() error {
	table := generateCIDRTable()

	return printOutput(table)
}

func generateCIDRTable() []*CIDRNetwork {
	subnets := []*CIDRNetwork{}
	for cidr := 0; cidr < 33; cidr++ {
		numAddresses := calculateNumAddresses(32, cidr)
		mask := calculateSubnetMask(cidr)

		snet := &CIDRNetwork{
			NetworkPortion: uint(cidr),
			SubnetMask:     mask,
			AvailableHosts: numAddresses,
		}

		subnets = append(subnets, snet)
	}

	return subnets
}

func calculateSubnetMask(cidr int) string {
	ones := strings.Repeat("1", cidr)
	zeroes := strings.Repeat("0", 32-cidr)

	mask := ones + zeroes

	octet1 := must(strconv.ParseInt(mask[0:8], 2, 32))
	octet2 := must(strconv.ParseInt(mask[8:16], 2, 32))
	octet3 := must(strconv.ParseInt(mask[16:24], 2, 32))
	octet4 := must(strconv.ParseInt(mask[24:32], 2, 32))

	return fmt.Sprintf("%v.%v.%v.%v", octet1, octet2, octet3, octet4)
}

func calculateNumAddresses(addressLength, prefixLength int) uint {
	numAddresses := math.Pow(2, float64(addressLength)-float64(prefixLength))

	return uint(numAddresses)
}

func printOutput(subnets []*CIDRNetwork) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "CIDR\tSubnet Mask\tAddresses\n")
	for _, subnet := range subnets {
		fmt.Fprintf(w, "/%v\t%s\t%v\n", subnet.NetworkPortion, subnet.SubnetMask, subnet.AvailableHosts)
	}

	return w.Flush()
}

func must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}

	return x
}
