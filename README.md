# IPv4 Subnet Calculator (GO)

Network calculator for subnet mask and other classless (CIDR) network information.


### Calculations
 * IP address network subnet masks, network and host portions, and provides aggregated reports.
 * Subnet mask
 * Network portion
 * Host portion
 * Number of IP addresses in the network
 * Number of addressable hosts in the network
 * IP address range
 * Broadcast address

Provides each data in dotted quads, hexadecimal, and binary formats, as well as array of quads.

## Usages

### Install

```go
go get -u github.com/brotherpowers/ipsubnet
```

### Get Started

```go
import "github.com/brotherpowers/ipsubnet"

sub := ipsubnet.SubnetCalculator("192.168.112.203", 23)
```


### Get Network Information

```go
sub.GetNumberIPAddresses();      // 512
sub.GetNumberAddressableHosts(); // 510
sub.GetIPAddressRange();         // [ 192.168.112.0, 192.168.113.255 ]
sub.GetNetworkSize();            // 23
sub.GetBroadcastAddress();       // 192.168.113.255
```

### Get IP Address
```go
sub.getIPAddress();       // 192.168.112.203
sub.getIPAddressQuads();  // [ 192, 168, 112, 203 ]
sub.getIPAddressHex();    // C0A870CB
sub.getIPAddressBinary(); // 11000000101010000111000011001011
```

### Get Subnet Mask
```go
sub.getSubnetMask();       // 255.255.254.0
sub.getSubnetMaskQuads();  // [ 255, 255, 254, 0 ]
sub.getSubnetMaskHex();    // FFFFFE00
sub.getSubnetMaskBinary(); // 11111111111111111111111000000000
```

### Get Network Portion
```go
sub.getNetworkPortion();       // 192.168.112.0
sub.getNetworkPortionQuads();  // [ 192, 168, 112, 0 ]
sub.getNetworkPortionHex();    // C0A87000
sub.getNetworkPortionBinary(); // 11000000101010000111000000000000
```

### Get Host Portion
```go
sub.getHostPortion();       // 0.0.0.203
sub.getHostPortionQuads();  // [ 0, 0, 0, 203 ]
sub.getHostPortionHex();    // 000000CB
sub.getHostPortionBinary(); // 00000000000000000000000011001011
```


## Based On

[IPv4 Subnet Calculator ( PHP )](https://github.com/markrogoyski/ipv4-subnet-calculator-php)

## License

IPv4 Subnet Calculator (GO) is licensed under the MIT License.
