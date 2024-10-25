package ranges

import (
	"fmt"
	"math"
	"os"
	"strings"
	"text/tabwriter"
)

const (
	LINE_DELIMITER = ","
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (*handler) Handle() error {
	subnets := parseText()

	return printOutput(subnets)
}

func parseText() []*subnet {
	masks := strings.Split(strings.ReplaceAll(ranges, "\n", ""), LINE_DELIMITER)

	subnets := []*subnet{}
	for cidr := 0; cidr < 33; cidr++ {
		numAddresses := calculateNumAddresses(32, cidr)

		snet := &subnet{
			CIDR:         uint(cidr),
			Mask:         masks[cidr],
			NumAddresses: numAddresses,
		}

		subnets = append(subnets, snet)
	}

	return subnets
}

func printOutput(subnets []*subnet) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "CIDR\tSubnet Mask\tAddresses\n")
	for _, subnet := range subnets {

		fmt.Fprintf(w, "/%v\t%s\t%v\n", subnet.CIDR, subnet.Mask, subnet.NumAddresses)
	}

	return w.Flush()
}

func calculateNumAddresses(addressLength, prefixLength int) uint {
	numAddresses := math.Pow(2, float64(addressLength)-float64(prefixLength))

	return uint(numAddresses)
}
