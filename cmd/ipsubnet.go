package main

import (
	"flag"
	"fmt"
	"github.com/arnobroekhof/ipsubnet"
)

var (
	mask    int
	network string
)

type NetworkInfo struct {
	NumberOfIpaddresses      int      `json:"number-of-addresses"`
	NumberOfAddressableHosts int      `json:"number-of-available_addresses"`
	IpRange                  []string `json:"ip-range"`
	NetworkBits              int      `json:"network-bits"`
	BroadcastAddress         string   `json:"broadcast-address"`
}

func init() {
	flag.IntVar(&mask, "mask", 24, "set the subnet mask")
	flag.StringVar(&network, "network", "192.168.0.0", "network or host")
	flag.Parse()
}

func main() {

	sub := ipsubnet.SubnetCalculator(network, mask)

	n := NetworkInfo{
		NumberOfIpaddresses:      sub.GetNumberIPAddresses(),
		NumberOfAddressableHosts: sub.GetNumberAddressableHosts(),
		IpRange:                  sub.GetIPAddressRange(),
		NetworkBits:              sub.GetNetworkSize(),
		BroadcastAddress:         sub.GetBroadcastAddress(),
	}

	fmt.Printf("network:\t\t\t%s/%d\n", network, mask)
	fmt.Printf("number-of-addresses:\t\t%d\n", n.NumberOfIpaddresses)
	fmt.Printf("number-of-available-addresses:\t%d\n", n.NumberOfAddressableHosts)
	fmt.Printf("ip-range:\t\t\t[ %s - %s ]\n", n.IpRange[0], n.IpRange[1])
	fmt.Printf("broadcast-address:\t\t%s\n", n.BroadcastAddress)

}
