package subnet

import (
	"cider/internal/cidr"
	"cider/internal/ip"
	"cider/internal/list"
	"cider/internal/must"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type handler struct {
	stdout io.Writer
}

func New(stdout io.Writer) *handler {
	return &handler{
		stdout: stdout,
	}
}

func (h *handler) Handle(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("command expects at least 2 arguments")
	}

	rangeToSplit := args[0]
	sizes := list.Map(args[1:], func(s string) int {
		n, _ := strconv.Atoi(s)

		return n
	})

	s := strings.Split(rangeToSplit, "/")

	ip := ip.NewIp(s[0])

	host := int(must.Must(strconv.ParseInt(s[1], 10, 32)))

	block := cidr.NewBlock(ip, host)

	subnets, err := block.Subnet(sizes)

	if err != nil {
		return err
	}

	for _, subnet := range subnets {
		fmt.Fprintln(h.stdout, subnet)
	}

	return nil
}
