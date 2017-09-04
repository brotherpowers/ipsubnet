package ipsubnet

import "testing"

type TestCase struct {
	ip                     string
	subnet                 int
	numberIPAddresses      int
	numberAddressableHosts int
	ipAddressRange         []string
	networkSize            int
	broadCastAddress       string
	ipAddressQuards        []int
	ipAddressHex           string
	ipAddressBinary        string
	subnetMask             string
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
	test.ipAddressQuards = []int{192, 168, 112, 203}
	test.ipAddressHex = "C0A870CB"
	test.ipAddressBinary = "11000000101010000111000011001011"
	test.subnetMask = "255.255.254.0"

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
