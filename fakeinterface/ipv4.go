package fakeinterface

import (
	"net"
	"sync/atomic"

	"github.com/newtools/zsocket/inet"
	"github.com/newtools/zsocket/nettypes"
)

const _StartgingID = 12345

var id uint32 = _StartgingID

func IPv4Packet(source, dest *net.IPAddr, proto nettypes.IPProtocol, pay []byte, len uint16) (nettypes.IPv4Packet, uint16) {
	l := 20 + len
	ipv4 := nettypes.IPv4Packet(make([]byte, l))
	copy(ipv4[20:], pay[:len])
	// 0100 0101
	ipv4[0] = 69
	ipv4[1] = 0
	inet.PutHToNS(ipv4[2:4], l)
	inet.PutHToNS(ipv4[4:6], getNextId())
	ipv4[6] = 64
	ipv4[7] = 0
	ipv4[8] = 64
	ipv4[9] = byte(proto)
	ipv4[10] = 0
	ipv4[11] = 0
	copy(ipv4[12:16], source.IP)
	copy(ipv4[16:20], dest.IP)
	inet.PutShort(ipv4[10:12], ipv4.CalculateChecksum())
	return ipv4, l
}

func getNextId() uint16 {
	n := atomic.AddUint32(&id, 1)
	if n <= 65535 {
		return uint16(n)
	}
	atomic.CompareAndSwapUint32(&id, n, _StartgingID)
	return getNextId()
}
