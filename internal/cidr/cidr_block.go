package cidr

import (
	"cider/internal/ip"
	"cider/internal/list"
	"cider/internal/must"
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

type CidrBlock struct {
	Ip   *ip.Ip
	Host int
}

func NewBlock(ip *ip.Ip, host int) *CidrBlock {
	return &CidrBlock{
		Ip:   ip,
		Host: host,
	}
}

func (b *CidrBlock) Subnet(sizes []int) ([]string, error) {
	// sort subnets largest to smallest to prevent fragmentation
	slices.Sort(sizes)

	next := b
	subnets := []string{}
	for _, size := range sizes {
		subnetBlock := NewBlock(next.Ip, size)

		if !b.ContainsIp(subnetBlock.Ip) {
			return nil, fmt.Errorf("invalid configuration: subnet %s/%v is outside provided network range %s/%v", next.Ip, size, b.Ip, b.Host)
		}

		subnets = append(subnets, fmt.Sprintf("%s/%v", subnetBlock.Ip.Ip(), subnetBlock.Host))

		next = NewBlock(ip.NewIp(subnetBlock.StartAddressOfNextBlock()), size)
	}

	return subnets, nil
}

// https://stackoverflow.com/questions/9622967/how-to-see-if-an-ip-address-belongs-inside-of-a-range-of-ips-using-cidr-notation
func (b *CidrBlock) ContainsIp(inner *ip.Ip) bool {
	IP_addr := inner.ToDecimal()
	CIDR_addr := b.Ip.ToDecimal()
	CIDR_mask := -1 << (INT_SIZE - b.Host)

	return (IP_addr & CIDR_mask) == (CIDR_addr & CIDR_mask)
}

func (b *CidrBlock) ContainsRange(inner *CidrBlock) bool {
	return b.ContainsIp(inner.Ip)
}

func (b *CidrBlock) NetworkPortionBinary() string {
	octets := strings.Split(b.Ip.Ip(), ".")
	octets = list.Map(octets, toBin)

	return fmt.Sprintf("%s.%s.%s.%s", octets[0], octets[1], octets[2], octets[3])
}

func toBin(s string) string {
	asInt := must.Must(strconv.ParseUint(s, 10, INT_SIZE))
	asBinaryString := strconv.FormatUint(asInt, 2)
	paddedBinaryString := utils.PadLeft(asBinaryString, '0', 8)
	return paddedBinaryString
}

func (b *CidrBlock) ToDecimal() int {
	return b.Ip.ToDecimal()
}

func (b *CidrBlock) Mask() string {
	ones := strings.Repeat("1", b.Host)
	zeroes := strings.Repeat("0", INT_SIZE-b.Host)

	mask := ones + zeroes

	octets := stringToOctets(mask)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func (b *CidrBlock) AvailableHosts() uint {
	numAddresses := math.Pow(2, float64(INT_SIZE)-float64(b.Host))

	return uint(numAddresses)
}

func (b *CidrBlock) AvailableAzureHosts() string {
	// Azure reserves the first four addresses and the last address, for a total of five IP addresses within each subnet
	// https://learn.microsoft.com/en-us/azure/virtual-network/virtual-networks-faq#are-there-any-restrictions-on-using-ip-addresses-within-these-subnets
	if b.AvailableHosts() >= 5 {
		return fmt.Sprintf("%v", b.AvailableHosts()-5)
	}

	return "N/A"
}

func (b *CidrBlock) StartAddressOfNextBlock() string {
	octets := strings.Split(b.BroadcastAddress(), ".")
	octets = list.Map(octets, toBin)
	binStr := strings.Join(octets, "")

	next := must.Must(strconv.ParseUint(binStr, 2, INT_SIZE)) + 1

	asBinaryString := strconv.FormatUint(next, 2)
	asBinaryString = utils.PadLeft(asBinaryString, '0', 32)

	octetsInt := stringToOctets(asBinaryString)

	return fmt.Sprintf("%v.%v.%v.%v", octetsInt[0], octetsInt[1], octetsInt[2], octetsInt[3])
}

func (b *CidrBlock) NetworkAddress() string {
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.Host]

	broadcast := ipBin + strings.Repeat("0", INT_SIZE-b.Host)

	octets := stringToOctets(broadcast)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func (b *CidrBlock) BroadcastAddress() string {
	// https://stackoverflow.com/questions/1470792/how-to-calculate-the-ip-range-when-the-ip-address-and-the-netmask-is-given
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.Host]

	broadcast := ipBin + strings.Repeat("1", INT_SIZE-b.Host)

	octets := stringToOctets(broadcast)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func stringToOctets(ipString string) []uint64 {
	octets := make([]uint64, 4)

	base := 2
	octets[0] = must.Must(strconv.ParseUint(ipString[0:8], base, INT_SIZE))
	octets[1] = must.Must(strconv.ParseUint(ipString[8:16], base, INT_SIZE))
	octets[2] = must.Must(strconv.ParseUint(ipString[16:24], base, INT_SIZE))
	octets[3] = must.Must(strconv.ParseUint(ipString[24:32], base, INT_SIZE))

	return octets
}
