package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetSubnetMask() string {
	return s.subnetCalculation("%d", ".")
}

func (s *Ip) GetSubnetMaskQuards() []byte {
	return convertQuardsToInt(strings.Split(s.subnetCalculation("%d", "."), "."))
}

func (s *Ip) GetSubnetMaskHex() string {
	return s.subnetCalculation("%02X", "")
}

func (s *Ip) GetSubnetMaskBinary() string {
	return s.subnetCalculation("%08b", "")
}

func (s *Ip) subnetCalculation(format, separator string) string {
	maskQuards := []string{}
	maskQuards = append(maskQuards, fmt.Sprintf(format, (s.subnet_mask>>24)&0xFF))
	maskQuards = append(maskQuards, fmt.Sprintf(format, (s.subnet_mask>>16)&0xFF))
	maskQuards = append(maskQuards, fmt.Sprintf(format, (s.subnet_mask>>8)&0xFF))
	maskQuards = append(maskQuards, fmt.Sprintf(format, (s.subnet_mask>>0)&0xFF))

	return strings.Join(maskQuards, separator)
}
