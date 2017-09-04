package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetNetworkPortion() string {
	return s.networkCalculation("%d", ".")
}

func (s *Ip) GetNetworkPortionQuards() []int {
	return convertQuardsToInt(strings.Split(s.networkCalculation("%d", "."), "."))
}

func (s *Ip) GetNetworkPortionHex() string {
	return s.networkCalculation("%02X", "")
}

func (s *Ip) GetNetworkPortionBinary() string {
	return s.networkCalculation("%08b", "")
}

func (s *Ip) networkCalculation(format, separator string) string {
	splits := s.GetIPAddressQuads()
	networkQuards := []string{}
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[0]&(s.subnet_mask>>24)))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[1]&(s.subnet_mask>>16)))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[2]&(s.subnet_mask>>8)))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[3]&(s.subnet_mask>>0)))

	return strings.Join(networkQuards, separator)
}
