package cidr

import (
	"cider/internal/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	INT_SIZE = 32
)

type CIDRBlock struct {
	Network     string
	HostPortion int
}

func NewBlock(network string) *CIDRBlock {
	networkAndHostPortion := strings.Split(network, "/")

	networkPortion := networkAndHostPortion[0]
	hostPortion := networkAndHostPortion[1]

	return &CIDRBlock{
		Network:     networkPortion,
		HostPortion: must(strconv.Atoi(hostPortion)),
	}
}

func padLeft(s string, paddingChar rune, totalWidth int) string {
	if len(s) >= totalWidth {
		return s
	}

	padding := totalWidth - len(s)

	return strings.Repeat(string(paddingChar), padding) + s
}

func (b *CIDRBlock) NetworkPortionBinary() string {
	octets := strings.Split(b.Network, ".")
	octets = list.Map(octets, toBin)

	return fmt.Sprintf("%s.%s.%s.%s", octets[0], octets[1], octets[2], octets[3])
}

func toBin(s string) string {
	asInt := must(strconv.ParseInt(s, 10, INT_SIZE))
	asBinaryString := strconv.FormatInt(asInt, 2)
	paddedBynaryString := padLeft(asBinaryString, '0', 8)
	return paddedBynaryString
}

func (b *CIDRBlock) SubnetMask() string {
	ones := strings.Repeat("1", b.HostPortion)
	zeroes := strings.Repeat("0", INT_SIZE-b.HostPortion)

	mask := ones + zeroes

	base := 2
	octet1 := must(strconv.ParseInt(mask[0:8], base, INT_SIZE))
	octet2 := must(strconv.ParseInt(mask[8:16], base, INT_SIZE))
	octet3 := must(strconv.ParseInt(mask[16:24], base, INT_SIZE))
	octet4 := must(strconv.ParseInt(mask[24:32], base, INT_SIZE))

	return fmt.Sprintf("%v.%v.%v.%v", octet1, octet2, octet3, octet4)
}

func (b *CIDRBlock) AvailableHosts() uint {
	numAddresses := math.Pow(2, float64(INT_SIZE)-float64(b.HostPortion))

	return uint(numAddresses)
}

func (b *CIDRBlock) StartAddressOfNextBlock() string {
	octets := strings.Split(b.BroadcastAddress(), ".")
	octets = list.Map(octets, toBin)
	binStr := strings.Join(octets, "")

	next := must(strconv.ParseInt(binStr, 2, INT_SIZE)) + 1

	asBinaryString := strconv.FormatInt(next, 2)
	asBinaryString = padLeft(asBinaryString, '0', 32)

	base := 2
	octet1 := must(strconv.ParseInt(asBinaryString[0:8], base, INT_SIZE))
	octet2 := must(strconv.ParseInt(asBinaryString[8:16], base, INT_SIZE))
	octet3 := must(strconv.ParseInt(asBinaryString[16:24], base, INT_SIZE))
	octet4 := must(strconv.ParseInt(asBinaryString[24:32], base, INT_SIZE))

	return fmt.Sprintf("%v.%v.%v.%v", octet1, octet2, octet3, octet4)
}

func (b *CIDRBlock) NetworkAddress() string {
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.HostPortion]

	broadcast := ipBin + strings.Repeat("0", INT_SIZE-b.HostPortion)

	base := 2
	octet1 := must(strconv.ParseInt(broadcast[0:8], base, INT_SIZE))
	octet2 := must(strconv.ParseInt(broadcast[8:16], base, INT_SIZE))
	octet3 := must(strconv.ParseInt(broadcast[16:24], base, INT_SIZE))
	octet4 := must(strconv.ParseInt(broadcast[24:32], base, INT_SIZE))

	return fmt.Sprintf("%v.%v.%v.%v", octet1, octet2, octet3, octet4)
}

func (b *CIDRBlock) BroadcastAddress() string {
	// https://stackoverflow.com/questions/1470792/how-to-calculate-the-ip-range-when-the-ip-address-and-the-netmask-is-given

	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.HostPortion]

	broadcast := ipBin + strings.Repeat("1", INT_SIZE-b.HostPortion)

	base := 2
	octet1 := must(strconv.ParseInt(broadcast[0:8], base, INT_SIZE))
	octet2 := must(strconv.ParseInt(broadcast[8:16], base, INT_SIZE))
	octet3 := must(strconv.ParseInt(broadcast[16:24], base, INT_SIZE))
	octet4 := must(strconv.ParseInt(broadcast[24:32], base, INT_SIZE))

	return fmt.Sprintf("%v.%v.%v.%v", octet1, octet2, octet3, octet4)
}

func must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}

	return x
}
