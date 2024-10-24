package ranges

import (
	"fmt"
	"math"
	"os"
	"strings"
	"text/tabwriter"
)

const (
	ENTRY_DELIMITER = ";"
	LINE_DELIMITER  = ","
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
	t := strings.Split(strings.ReplaceAll(ranges, "\n", ""), LINE_DELIMITER)

	subnets := []*subnet{}
	for _, line := range t {
		s := strings.Split(line, ENTRY_DELIMITER)

		snet := SubnetFromText(s)

		subnets = append(subnets, snet)
	}

	return subnets
}

func printOutput(subnets []*subnet) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "CIDR\tSubnet Mask\tAddresses\n")
	for _, subnet := range subnets {
		numAddresses := calculateNumAddresses(32, subnet.CIDR)

		fmt.Fprintf(w, "/%v\t%s\t%v\n", subnet.CIDR, subnet.Mask, numAddresses)
	}

	return w.Flush()
}

func calculateNumAddresses(addressLength, prefixLength uint) uint {
	numAddresses := math.Pow(2, float64(addressLength)-float64(prefixLength))

	return uint(numAddresses)
}
