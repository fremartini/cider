package ranges

import (
	"fmt"
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
	t := strings.Split(ranges, ",")

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

	fmt.Fprint(w, "CIDR\tSubnet Mask\t# of addresses\tWildcard\n")
	for _, subnet := range subnets {
		fmt.Fprintf(w, "%s\t%s\t%v\t%s", subnet.CIDR, subnet.Mask, subnet.NumAddresses, subnet.Wildcard)
	}

	return w.Flush()
}
