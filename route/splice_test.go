package route

import (
	"net"
	"testing"

	M "github.com/sagernet/sing/common/metadata"
	"github.com/sagernet/sing/protocol/socks"
)

func TestUnwrapUDPSpliceTarget(t *testing.T) {
	listener, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { listener.Close() })

	manager := new(ConnectionManager)
	t.Run("direct", func(t *testing.T) {
		tracked := manager.TrackPacketConn(listener)
		wrapped := &transparentPacketConn{PacketConn: tracked}
		target, ok := unwrapSpliceTarget(wrapped)
		if !ok || target.socket != tracked.(*trackedPacketConn) {
			t.Fatalf("unexpected direct target: %+v, ok=%v", target, ok)
		}
		if wrapped.readChecks != 1 || wrapped.writeChecks != 1 {
			t.Fatalf("replaceability checked %d times for read and %d times for write", wrapped.readChecks, wrapped.writeChecks)
		}
	})

	t.Run("SOCKS5", func(t *testing.T) {
		conn, err := net.DialUDP("udp4", nil, listener.LocalAddr().(*net.UDPAddr))
		if err != nil {
			t.Fatal(err)
		}
		tracked := manager.TrackConn(conn)
		control, peer := net.Pipe()
		t.Cleanup(func() { peer.Close() })
		association := socks.NewAssociatePacketConn(tracked, M.SocksaddrFromNet(listener.LocalAddr()), control)
		t.Cleanup(func() { association.Close() })

		wrapped := &transparentPacketConn{PacketConn: association}
		if _, ok := unwrapSpliceTarget(wrapped); ok {
			t.Fatal("packet splice accepted a socket that still needs the SOCKS5 codec")
		}
	})
}

type transparentPacketConn struct {
	net.PacketConn
	readChecks  int
	writeChecks int
}

func (c *transparentPacketConn) ReaderReplaceable() bool {
	c.readChecks++
	return true
}

func (c *transparentPacketConn) WriterReplaceable() bool {
	c.writeChecks++
	return true
}

func (c *transparentPacketConn) Upstream() any {
	return c.PacketConn
}
