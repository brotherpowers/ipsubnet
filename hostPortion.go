package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetHostPortion() string {
	return s.hostCalculation("%d", ".")
}

func (s *Ip) GetHostPortionQuards() []byte {
	return convertQuardsToInt(strings.Split(s.hostCalculation("%d", "."), "."))
}

func (s *Ip) GetHostPortionHex() string {
	return s.hostCalculation("%02X", "")
}

func (s *Ip) GetHostPortionBinary() string {
	return s.hostCalculation("%08b", "")
}

func (s *Ip) hostCalculation(format, separator string) string {
	splits := s.GetIPAddressQuads()
	networkQuards := []string{}
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[0] & ^byte(s.subnet_mask>>24)))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[1] & ^byte(s.subnet_mask>>16)))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[2] & ^byte(s.subnet_mask>>8)))
	networkQuards = append(networkQuards, fmt.Sprintf(format, splits[3] & ^byte(s.subnet_mask>>0)))

	return strings.Join(networkQuards, separator)
}
