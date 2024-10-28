package ranges

type CIDRBlock struct {
	NetworkPortion uint   // 0
	SubnetMask     string // 0.0.0.0
	AvailableHosts uint   // 4294967296
}
