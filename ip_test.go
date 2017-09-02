package ipsubnet

import "testing"

type TestCase struct {
	ip                     string
	subnet                 int
	numberIPAddresses      int
	numberAddressableHosts int
	ipAddressRange         []string
	networkSize            int
}

func builder() TestCase {
	test := TestCase{}
	test.ip = "192.168.112.203"
	test.subnet = 23
	test.numberIPAddresses = 512
	test.numberAddressableHosts = 510
	test.ipAddressRange = []string{"192.168.112.0", "192.168.113.255"}
	test.networkSize = 23

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
		t.Errorf("First Value of GetIPAddressRange Failed: Wants %v Got %v", wants, got)
	}

}

func TestGetIPAddress(t *testing.T) {
	wants := builder().ip
	got := ip().GetIPAddress()

	if wants != got {
		t.Errorf("GetIPAddress Failed: Wants %v Got %v", wants, got)
	}
}
