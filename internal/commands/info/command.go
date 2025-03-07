package info

import (
	"cider/internal/cidr"
	"fmt"
	"os"
	"text/tabwriter"
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

	w := tabwriter.NewWriter(os.Stdout, 2, 4, 1, ' ', 0)

	fmt.Fprint(w, "Address range\tStart of next block\tCidr\tSubnet mask\tAddresses\n")
	fmt.Fprintf(w, "%v - %v\t%s\t/%v\t%v\t%v\n", block.NetworkAddress(), block.BroadcastAddress(), block.StartAddressOfNextBlock(), block.HostPortion, block.SubnetMask(), block.AvailableHosts())

	return w.Flush()
}
