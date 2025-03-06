package subnet

import (
	"fmt"
)

type handler struct{}

func New() *handler {
	return &handler{}
}

func (*handler) Handle(args []string) error {
	rangeToSplit := args[0]
	sizes := args[1:]

	fmt.Println(rangeToSplit, sizes)

	return nil
}

/*

    public IEnumerable<string> Subnet(string range, IEnumerable<int> sizes)
    {
        var vnet = ToCidrBlock(range);

        // the smallest vnet is a /32
        if (vnet.HostPortion > 32)
        {
            throw new InvalidRangeException("Smallest vnet size is 32");
        }

        // the smallest subnet is a /32
        if (sizes.Any(x => x > 32))
        {
            throw new InvalidRangeException("Smallest subnet size is 32");
        }

        var vnetLastIp = vnet.StartAndEndAddress().Item2;

        var result = new List<string>();

        sizes = sizes.Order();

        var nextIp = vnet.NetworkPortion;
        foreach (var subnetSize in sizes)
        {
            var cidrBlock = new CidrBlock {
                NetworkPortion = nextIp,
                HostPortion = subnetSize,
            };

            var c = cidrBlock.StartAddressOfNextBlock();

            var octets = c.Split(".");
            var octetsBytes = octets.Select(ToByte);
            nextIp = octetsBytes;

            result.Add($"{cidrBlock.ToIp()}/{subnetSize}");
        }

        return result;
    }

    private static CidrBlock ToCidrBlock(string ip)
    {
        var networkAndHostPortion = ip.Split("/");

        var networkPortion = networkAndHostPortion[0]; // 10.0.0.0
        var octets = networkPortion.Split(".");
        var octetsBytes = octets.Select(ToByte);

        var hostPortion = Convert.ToInt32(networkAndHostPortion[1]); // / 16

        return new CidrBlock
        {
            NetworkPortion = octetsBytes,
            HostPortion = hostPortion,
        };
    }

    public static byte ToByte(string n)
    {
        var padded = n.PadLeft(8, '0');

        return byte.Parse(padded);
    }

    public static string ToBinary(byte b)
    {
        return Convert.ToString(b, 2).PadLeft(8, '0');
    }
}


*/
