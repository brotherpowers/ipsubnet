package ipsubnet

import "testing"

type TestCase struct {
	ip                     string
	subnet                 int64
	numberIPAddresses      int64
	numberAddressableHosts int64
	ipAddressRange         []string
	networkSize            int64
	broadCastAddress       string
	ipAddressQuards        []byte
	ipAddressHex           string
	ipAddressBinary        string
	subnetMask             string
	subnetMaskQuards       []byte
	subnetMaskHex          string
	subnetMaskBinary       string

	networkPortion       string
	networkPortionQuards []byte
	networkPortionHex    string
	networkPortionBinary string

	hostPortion       string
	hostPortionQuards []byte
	hostPortionHex    string
	hostPortionBinary string
}

func builder() TestCase {
	test := TestCase{}
	test.ip = "192.168.112.203"
	test.subnet = 23
	test.numberIPAddresses = 512
	test.numberAddressableHosts = 510
	test.ipAddressRange = []string{"192.168.112.0", "192.168.113.255"}
	test.networkSize = 23
	test.broadCastAddress = "192.168.113.255"
	test.ipAddressQuards = []byte{192, 168, 112, 203}
	test.ipAddressHex = "C0A870CB"
	test.ipAddressBinary = "11000000101010000111000011001011"
	// Subnet
	test.subnetMask = "255.255.254.0"
	test.subnetMaskQuards = []byte{255, 255, 254, 0}
	test.subnetMaskHex = "FFFFFE00"
	test.subnetMaskBinary = "11111111111111111111111000000000"
	// Network Portion
	test.networkPortion = "192.168.112.0"
	test.networkPortionQuards = []byte{192, 168, 112, 0}
	test.networkPortionHex = "C0A87000"
	test.networkPortionBinary = "11000000101010000111000000000000"
	// Network Portion
	test.hostPortion = "0.0.0.203"
	test.hostPortionQuards = []byte{0, 0, 0, 203}
	test.hostPortionHex = "000000CB"
	test.hostPortionBinary = "00000000000000000000000011001011"
	return test
}

func ip() *Ip {
	return SubnetCalculator(builder().ip, builder().subnet)
}

func TestGetNumberIPAddresses(t *testing.T) {
	wants := builder().numberIPAddresses
	got := ip().GetNumberIPAddresses()

	if wants != got {
		t.Errorf("GetNumberIPAddresses Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetNumberAddressableHosts(t *testing.T) {
	wants := builder().numberAddressableHosts
	got := ip().GetNumberAddressableHosts()

	if wants != got {
		t.Errorf("GetNumberAddressableHosts Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetIPAddressRange(t *testing.T) {
	wants := builder().ipAddressRange
	got := ip().GetIPAddressRange()

	if wants[0] != got[0] {
		t.Errorf("First Value of GetIPAddressRange Failed: Wants %v Got %v", wants[0], got[0])
	}

	if wants[1] != got[1] {
		t.Errorf("First Value of GetIPAddressRange Failed: Wants %v Got %v", wants[1], got[1])
	}
}

func TestGetNetworSize(t *testing.T) {
	wants := builder().networkSize
	got := ip().GetNetworkSize()

	if wants != got {
		t.Errorf("GetNetworkSize Failed: Wants %v Got %v", wants, got)
	}

}

func TestGetBroadcastAddress(t *testing.T) {
	wants := builder().broadCastAddress
	got := ip().GetBroadcastAddress()

	if wants != got {
		t.Errorf("GetBroadcastAdress Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetIPAddress(t *testing.T) {
	wants := builder().ip
	got := ip().GetIPAddress()

	if wants != got {
		t.Errorf("GetIPAddress Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetIPAddressQuads(t *testing.T) {
	wants := builder().ipAddressQuards
	got := ip().GetIPAddressQuads()

	if got[0] != wants[0] {
		t.Errorf("First Value of GetIPAddressQuards Failed: Wants %v Got %v", wants[0], got[0])
	}
	if got[1] != wants[1] {
		t.Errorf("Second Value of GetIPAddressQuards Failed: Wants %v Got %v", wants[1], got[1])
	}
	if got[2] != wants[2] {
		t.Errorf("Third Value of GetIPAddressQuards Failed: Wants %v Got %v", wants[2], got[2])
	}
	if got[3] != wants[3] {
		t.Errorf("Fourth Value of GetIPAddressQuards Failed: Wants %v Got %v", wants[3], got[3])
	}
}

func TestGetIPAddressHex(t *testing.T) {
	wants := builder().ipAddressHex
	got := ip().GetIPAddressHex()

	if wants != got {
		t.Errorf("GetIpAddressHex Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetIPAddressBinary(t *testing.T) {
	wants := builder().ipAddressBinary
	got := ip().GetIPAddressBinary()

	if wants != got {
		t.Errorf("GetIpAddressBinary Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetSubnetMask(t *testing.T) {
	wants := builder().subnetMask
	got := ip().GetSubnetMask()

	if wants != got {
		t.Errorf("GetSubnetMask Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetSubnetMaskQuards(t *testing.T) {
	wants := builder().subnetMaskQuards
	got := ip().GetSubnetMaskQuards()

	if got[0] != wants[0] {
		t.Errorf("First Value of GetSubnetMaskQuards Failed: Wants %v Got %v", wants[0], got[0])
	}
	if got[1] != wants[1] {
		t.Errorf("Second Value of GetSubnetMaskQuards Failed: Wants %v Got %v", wants[1], got[1])
	}
	if got[2] != wants[2] {
		t.Errorf("Third Value of GetSubnetMaskQuards Failed: Wants %v Got %v", wants[2], got[2])
	}
	if got[3] != wants[3] {
		t.Errorf("Fourth Value of GetSubnetMaskQuards Failed: Wants %v Got %v", wants[3], got[3])
	}
}

func TestGetSubnetMaskHex(t *testing.T) {
	wants := builder().subnetMaskHex
	got := ip().GetSubnetMaskHex()

	if wants != got {
		t.Errorf("GetSubnetMaskHex Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetSubnetMaskBinary(t *testing.T) {
	wants := builder().subnetMaskBinary
	got := ip().GetSubnetMaskBinary()

	if wants != got {
		t.Errorf("GetSubnetMaskBinary Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetNetworkPortion(t *testing.T) {
	wants := builder().networkPortion
	got := ip().GetNetworkPortion()

	if wants != got {
		t.Errorf("GetNetworkPortion Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetNetworkPortionQuards(t *testing.T) {
	wants := builder().networkPortionQuards
	got := ip().GetNetworkPortionQuards()

	if got[0] != wants[0] {
		t.Errorf("First Value of GetNetworkPortionQuards Failed: Wants %v Got %v", wants[0], got[0])
	}
	if got[1] != wants[1] {
		t.Errorf("Second Value of GetNetworkPortionQuards Failed: Wants %v Got %v", wants[1], got[1])
	}
	if got[2] != wants[2] {
		t.Errorf("Third Value of GetNetworkPortionQuards Failed: Wants %v Got %v", wants[2], got[2])
	}
	if got[3] != wants[3] {
		t.Errorf("Fourth Value of GetNetworkPortionQuards Failed: Wants %v Got %v", wants[3], got[3])
	}
}

func TestGetNetworkPortionHex(t *testing.T) {
	wants := builder().networkPortionHex
	got := ip().GetNetworkPortionHex()

	if wants != got {
		t.Errorf("GetNetworkPortionHex Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetNetworkPortionBinary(t *testing.T) {
	wants := builder().networkPortionBinary
	got := ip().GetNetworkPortionBinary()

	if wants != got {
		t.Errorf("GetSubnetMaskBinary Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetHostPortion(t *testing.T) {
	wants := builder().hostPortion
	got := ip().GetHostPortion()

	if wants != got {
		t.Errorf("GetHostPortion Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetHostPortionQuards(t *testing.T) {
	wants := builder().hostPortionQuards
	got := ip().GetHostPortionQuards()

	if got[0] != wants[0] {
		t.Errorf("First Value of GetHostPortionQuards Failed: Wants %v Got %v", wants[0], got[0])
	}
	if got[1] != wants[1] {
		t.Errorf("Second Value of GetHostPortionQuards Failed: Wants %v Got %v", wants[1], got[1])
	}
	if got[2] != wants[2] {
		t.Errorf("Third Value of GetHostPortionQuards Failed: Wants %v Got %v", wants[2], got[2])
	}
	if got[3] != wants[3] {
		t.Errorf("Fourth Value of GetHostPortionQuards Failed: Wants %v Got %v", wants[3], got[3])
	}
}

func TestGetHostPortionHex(t *testing.T) {
	wants := builder().hostPortionHex
	got := ip().GetHostPortionHex()

	if wants != got {
		t.Errorf("GetHostPortionHex Failed: Wants %v Got %v", wants, got)
	}
}

func TestGetHostPortionBinary(t *testing.T) {
	wants := builder().hostPortionBinary
	got := ip().GetHostPortionBinary()

	if wants != got {
		t.Errorf("GetSubnetMaskBinary Failed: Wants %v Got %v", wants, got)
	}
}
