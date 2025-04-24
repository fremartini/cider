package cidr

import (
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

func (b *CIDRBlock) Subnet(sizes []int) ([]string, error) {
	// sort subnets largest to smallest to prevent fragmentation
	slices.Sort(sizes)

	next := b
	subnets := []string{}
	for _, size := range sizes {
		subnetBlock := NewBlock(fmt.Sprintf("%s/%v", next.Network, size))

		if !b.ContainsIp(subnetBlock.Network) {
			return nil, fmt.Errorf("invalid configuration: subnet %s/%v is outside provided network range %s/%v", next.Network, size, b.Network, b.HostPortion)
		}

		subnets = append(subnets, fmt.Sprintf("%s/%v", subnetBlock.Network, subnetBlock.HostPortion))

		next = NewBlock(fmt.Sprintf("%s/%v", subnetBlock.StartAddressOfNextBlock(), size))
	}

	return subnets, nil
}

// https://stackoverflow.com/questions/9622967/how-to-see-if-an-ip-address-belongs-inside-of-a-range-of-ips-using-cidr-notation
func (outer *CIDRBlock) ContainsIp(inner string) bool {
	innerIp := strings.Split(inner, "/")[0]

	IP_addr := ipToDecimal(innerIp)
	CIDR_addr := ipToDecimal(outer.Network)
	CIDR_mask := -1 << (INT_SIZE - outer.HostPortion)

	return (IP_addr & CIDR_mask) == (CIDR_addr & CIDR_mask)
}

// http://www.aboutmyip.com/AboutMyXApp/IP2Integer.jsp
func ipToDecimal(ip string) int {
	parts := strings.Split(ip, ".")

	base := 10
	octet1 := must(strconv.ParseUint(parts[0], base, INT_SIZE))
	octet2 := must(strconv.ParseUint(parts[1], base, INT_SIZE))
	octet3 := must(strconv.ParseUint(parts[2], base, INT_SIZE))
	octet4 := must(strconv.ParseUint(parts[3], base, INT_SIZE))

	return int((octet1 * 16777216) + (octet2 * 65536) + (octet3 * 256) + octet4)
}

func (b *CIDRBlock) NetworkPortionBinary() string {
	octets := strings.Split(b.Network, ".")
	octets = list.Map(octets, toBin)

	return fmt.Sprintf("%s.%s.%s.%s", octets[0], octets[1], octets[2], octets[3])
}

func toBin(s string) string {
	asInt := must(strconv.ParseUint(s, 10, INT_SIZE))
	asBinaryString := strconv.FormatUint(asInt, 2)
	paddedBynaryString := utils.PadLeft(asBinaryString, '0', 8)
	return paddedBynaryString
}

func (b *CIDRBlock) Mask() string {
	ones := strings.Repeat("1", b.HostPortion)
	zeroes := strings.Repeat("0", INT_SIZE-b.HostPortion)

	mask := ones + zeroes

	octets := stringToOctets(mask)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func (b *CIDRBlock) AvailableHosts() uint {
	numAddresses := math.Pow(2, float64(INT_SIZE)-float64(b.HostPortion))

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
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.HostPortion]

	broadcast := ipBin + strings.Repeat("0", INT_SIZE-b.HostPortion)

	octets := stringToOctets(broadcast)

	return fmt.Sprintf("%v.%v.%v.%v", octets[0], octets[1], octets[2], octets[3])
}

func (b *CIDRBlock) BroadcastAddress() string {
	// https://stackoverflow.com/questions/1470792/how-to-calculate-the-ip-range-when-the-ip-address-and-the-netmask-is-given
	ipBin := strings.ReplaceAll(b.NetworkPortionBinary(), ".", "")[0:b.HostPortion]

	broadcast := ipBin + strings.Repeat("1", INT_SIZE-b.HostPortion)

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
