package ipsubnet

import (
	"fmt"
	"strings"
)

func (s *Ip) GetIPAddress() string {
	return s.ip
}

func (s *Ip) GetIPAddressHex() string {
	return s.ipAddressCalculation("%X", "")
}

func (s *Ip) GetIPAddressBinary() string {
	return s.ipAddressCalculation("%08b", "")
}

func (s *Ip) GetIPAddressQuads() []byte {
	splits := strings.Split(s.ip, ".")

	return convertQuardsToInt(splits)
}

func (s *Ip) ipAddressCalculation(format, separator string) string {

	splits := s.GetIPAddressQuads()
	formatted := []string{}

	for _, quard := range splits {
		formatted = append(formatted, fmt.Sprintf(format, quard))
	}

	return strings.Join(formatted, separator)
}
