package cidr

import (
	"cider/internal/ip"
	"cider/internal/list"
	"cider/internal/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const (
	INT_SIZE = 32
)

type CIDRBlock struct {
	Ip   *ip.Ip
	Host int
}

func NewBlock(ip *ip.Ip, host string) *CIDRBlock {
	return &CIDRBlock{
		Ip:   ip,
		Host: must(strconv.Atoi(host)),
	}
}

func (b *CIDRBlock) Subnet(sizes []int) ([]string, error) {
	// sort subnets largest to smallest to prevent fragmentation
	slices.Sort(sizes)

	next := b
	subnets := []string{}
	for _, size := range sizes {
		subnetBlock := NewBlock(next.Ip, fmt.Sprintf("%v", size))

		if !b.Contains(subnetBlock.Ip.Ip()) {
			return nil, fmt.Errorf("invalid configuration: subnet %s/%v is outside provided network range %s/%v", next.Ip, size, b.Ip, b.Host)
		}

		subnets = append(subnets, fmt.Sprintf("%s/%v", subnetBlock.Ip.Ip(), subnetBlock.Host))

		next = NewBlock(ip.NewIp(subnetBlock.StartAddressOfNextBlock()), fmt.Sprintf("%v", size))
	}

	return subnets, nil
}

// https://stackoverflow.com/questions/9622967/how-to-see-if-an-ip-address-belongs-inside-of-a-range-of-ips-using-cidr-notation
func (outer *CIDRBlock) Contains(inner string) bool {
	innerNetwork := strings.Split(inner, "/")[0]

	IP_addr := ip.NewIp(innerNetwork).ToDecimal()
	CIDR_addr := outer.Ip.ToDecimal()
	CIDR_mask := -1 << (INT_SIZE - outer.Host)

	return (IP_addr & CIDR_mask) == (CIDR_addr & CIDR_mask)
}

func (b *CIDRBlock) NetworkPortionBinary() string {
	octets := strings.Split(b.Ip.Ip(), ".")
	octets = list.Map(octets, toBin)

	return fmt.Sprintf("%s.%s.%s.%s", octets[0], octets[1], octets[2], octets[3])
}

func toBin(s string) string {
	asInt := must(strconv.ParseUint(s, 10, INT_SIZE))
	asBinaryString := strconv.FormatUint(asInt, 2)
	paddedBinaryString := utils.PadLeft(asBinaryString, '0', 8)
	return paddedBinaryString
}

func (b *CIDRBlock) ToDecimal() int {
	return b.Ip.ToDecimal()
}

func (b *CIDRBlock) Mask() string {
	ones := strings.Repeat("1", b.Host)
	zeroes := strings.Repeat("0", INT_SIZE-b.Host)

	mask := ones + zeroes

	octets := stringToOctets(mask)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func (b *CIDRBlock) AvailableHosts() uint {
	numAddresses := math.Pow(2, float64(INT_SIZE)-float64(b.Host))

	return uint(numAddresses)
}

func (b *CIDRBlock) AvailableAzureHosts() string {
	// Azure reserves the first four addresses and the last address, for a total of five IP addresses within each subnet
	// https://learn.microsoft.com/en-us/azure/virtual-network/virtual-networks-faq#are-there-any-restrictions-on-using-ip-addresses-within-these-subnets
	if b.AvailableHosts() >= 5 {
		return fmt.Sprintf("%v", b.AvailableHosts()-5)
	}

	return "N/A"
}

func (b *CIDRBlock) StartAddressOfNextBlock() string {
	octets := strings.Split(b.BroadcastAddress(), ".")
	octets = list.Map(octets, toBin)
	binStr := strings.Join(octets, "")

	next := must(strconv.ParseUint(binStr, 2, INT_SIZE)) + 1

	asBinaryString := strconv.FormatUint(next, 2)
	asBinaryString = utils.PadLeft(asBinaryString, '0', 32)

	octetsInt := stringToOctets(asBinaryString)

	return fmt.Sprintf("%v.%v.%v.%v", octetsInt[0], octetsInt[1], octetsInt[2], octetsInt[3])
}

func (b *CIDRBlock) NetworkAddress() string {
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.Host]

	broadcast := ipBin + strings.Repeat("0", INT_SIZE-b.Host)

	octets := stringToOctets(broadcast)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func (b *CIDRBlock) BroadcastAddress() string {
	// https://stackoverflow.com/questions/1470792/how-to-calculate-the-ip-range-when-the-ip-address-and-the-netmask-is-given
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.Host]

	broadcast := ipBin + strings.Repeat("1", INT_SIZE-b.Host)

	octets := stringToOctets(broadcast)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func stringToOctets(ipString string) []uint64 {
	octets := make([]uint64, 4)

	base := 2
	octets[0] = must(strconv.ParseUint(ipString[0:8], base, INT_SIZE))
	octets[1] = must(strconv.ParseUint(ipString[8:16], base, INT_SIZE))
	octets[2] = must(strconv.ParseUint(ipString[16:24], base, INT_SIZE))
	octets[3] = must(strconv.ParseUint(ipString[24:32], base, INT_SIZE))

	return octets
}

func must[T any](x T, e error) T {
	if e != nil {
		panic(e)
	}

	return x
}
