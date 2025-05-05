package subnet

import (
	"cider/internal/cidr"
	"cider/internal/ip"
	"cider/internal/list"
	"cider/internal/must"
	"fmt"
	"strconv"
	"strings"
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

	s := strings.Split("/", rangeToSplit)

	ip := ip.NewIp(s[0])

	host := int(must.Must(strconv.ParseInt(s[1], 10, 32)))

	block := cidr.NewBlock(ip, host)

	subnets, err := block.Subnet(sizes)

	if err != nil {
		return err
	}

	for _, subnet := range subnets {
		fmt.Println(subnet)
	}

	return nil
}
