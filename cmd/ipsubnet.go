package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/arnobroekhof/ipsubnet"
)

var (
	mask    int
	network string
)

type NetworkInfo struct {
	numberOfIpaddresses      int      `json:"number-of-addresses"`
	numberOfAddressableHosts int      `json:"number-of-available_addresses"`
	ipRange                  []string `json:"ip-range"`
	networkBits              int      `json:"network-bits"`
	broadcastAddress         string   `json:"broadcast-address"`
}

func init() {
	flag.IntVar(&mask, "mask", 24, "set the subnet mask")
	flag.StringVar(&network, "network", "192.168.0.0", "network or host")
	flag.Parse()
}

func main() {

	sub := ipsubnet.SubnetCalculator(network, mask)

	n := NetworkInfo{
		numberOfIpaddresses:      sub.GetNumberIPAddresses(),
		numberOfAddressableHosts: sub.GetNumberAddressableHosts(),
		ipRange:                  sub.GetIPAddressRange(),
		networkBits:              sub.GetNetworkSize(),
		broadcastAddress:         sub.GetBroadcastAddress(),
	}

	fmt.Printf("network:\t\t\t%s/%d\n", network, mask)
	fmt.Printf("number-of-addresses:\t\t%d\n", n.numberOfIpaddresses)
	fmt.Printf("number-of-available-addresses:\t%d\n", n.numberOfAddressableHosts)
	fmt.Printf("ip-range:\t\t\t[ %s - %s ]\n", n.ipRange[0], n.ipRange[1])
	fmt.Printf("broadcast-address:\t\t%s\n",n.broadcastAddress)

}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
