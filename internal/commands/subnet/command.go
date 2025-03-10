package subnet

import (
	"cider/internal/cidr"
	"cider/internal/list"
	"fmt"
	"strconv"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (*handler) Handle(args []string) error {
	rangeToSplit := args[0]
	sizes := list.Map(args[1:], func(s string) int {
		n, _ := strconv.Atoi(s)

		return n
	})

	block := cidr.NewBlock(rangeToSplit)

	subnets, err := block.Subnet(sizes)

	if err != nil {
		return err
	}

	for _, subnet := range subnets {
		fmt.Println(subnet)
	}

	return nil
}
