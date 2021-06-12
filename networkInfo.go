package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetNumberIPAddresses() int64 {
	return 2 << uint(31-s.networkSize)
}

func (s *Ip) GetNumberAddressableHosts() int64 {
	if s.networkSize == 32 {
		return 1
	} else if s.networkSize == 31 {
		return 2
	}
	return s.GetNumberIPAddresses() - 2
}

func (s *Ip) GetNetworkSize() int64 {
	return s.networkSize
}

func (s *Ip) GetIPAddressRange() []string {
	return []string{s.GetNetworkPortion(), s.GetBroadcastAddress()}
}

func (s *Ip) GetBroadcastAddress() string {
	networkQuads := s.GetNetworkPortionQuards()
	numberIPAddress := s.GetNumberIPAddresses()
	networkRangeQuads := []string{}
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[0]&byte(s.subnet_mask>>24))+byte(((numberIPAddress-1)>>24)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[1]&byte(s.subnet_mask>>16))+byte(((numberIPAddress-1)>>16)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[2]&byte(s.subnet_mask>>8))+byte(((numberIPAddress-1)>>8)&0xFF)))
	networkRangeQuads = append(networkRangeQuads, fmt.Sprintf("%d", (networkQuads[3]&byte(s.subnet_mask>>0))+byte(((numberIPAddress-1)>>0)&0xFF)))

	return strings.Join(networkRangeQuads, ".")
}
