package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetNumberIPAddresses() int {
	return 2 << uint(31-s.networkSize)
}

func (s *Ip) GetNumberAddressableHosts() int {
	if s.networkSize == 32 {
		return 1
	} else if s.networkSize == 31 {
		return 2
	}
	return (s.GetNumberIPAddresses() - 2)
}

func (s *Ip) GetNetworkSize() int {
	return s.networkSize
}

func (s *Ip) GetIPAddressRange() []string {
	return []string{s.GetNetworkPortion(), s.GetBroadcastAddress()}
}

func (s *Ip) GetBroadcastAddress() string {
	networkQuads := s.GetNetworkPortionQuards()
	numberIPAddress := s.GetNumberIPAddresses()
	networkRangeQuads := []string{}
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[0]&(s.subnet_mask>>24))+(((numberIPAddress-1)>>24)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[1]&(s.subnet_mask>>16))+(((numberIPAddress-1)>>16)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[2]&(s.subnet_mask>>8))+(((numberIPAddress-1)>>8)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[3]&(s.subnet_mask>>0))+(((numberIPAddress-1)>>0)&0xFF)))

	return strings.Join(networkRangeQuads, ".")
}

func (s *Ip) GetHostMaxAddress() string {
	networkQuads := s.GetNetworkPortionQuards()
	numberIPAddress := s.GetNumberIPAddresses()
	networkRangeQuads := []string{}
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[0]&(s.subnet_mask>>24))+(((numberIPAddress-2)>>24)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[1]&(s.subnet_mask>>16))+(((numberIPAddress-2)>>16)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[2]&(s.subnet_mask>>8))+(((numberIPAddress-2)>>8)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[3]&(s.subnet_mask>>0))+(((numberIPAddress-2)>>0)&0xFF)))

	return strings.Join(networkRangeQuads, ".")
}