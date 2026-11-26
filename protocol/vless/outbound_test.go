package vless

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-vmess/vless"
	"github.com/sagernet/sing/common/buf"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"

	"github.com/stretchr/testify/require"
)

func TestVLESSXUDP(t *testing.T) {
	const user = "00000000-0000-4000-8000-000000000001"
	destination := M.ParseSocksaddr("192.0.2.1:53")
	payload := []byte("BE7000 XUDP")
	for _, connected := range []bool{false, true} {
		name := "ListenPacket"
		if connected {
			name = "DialContext"
		}
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			logger := log.NewNOPFactory().Logger()
			client, err := vless.NewClient(user, "", logger)
			require.NoError(t, err)
			service := vless.NewService[string](logger, xudpEcho{})
			service.UpdateUsers([]string{user}, []string{user}, []string{""})
			clientConn, serverConn := net.Pipe()
			defer clientConn.Close()
			defer serverConn.Close()
			deadline, _ := ctx.Deadline()
			require.NoError(t, clientConn.SetDeadline(deadline))
			require.NoError(t, serverConn.SetDeadline(deadline))
			go service.NewConnection(ctx, serverConn, destination, nil)

			outbound := &Outbound{
				logger: logger,
				dialer: xudpTestDialer{conn: clientConn},
				client: client,
			}
			response := make([]byte, 64)
			if connected {
				conn, err := outbound.DialContext(ctx, N.NetworkUDP, destination)
				require.NoError(t, err)
				defer conn.Close()
				_, err = conn.Write(payload)
				require.NoError(t, err)
				n, err := conn.Read(response)
				require.NoError(t, err)
				require.Equal(t, payload, response[:n])
			} else {
				conn, err := outbound.ListenPacket(ctx, destination)
				require.NoError(t, err)
				defer conn.Close()
				_, err = conn.WriteTo(payload, destination.UDPAddr())
				require.NoError(t, err)
				n, from, err := conn.ReadFrom(response)
				require.NoError(t, err)
				require.Equal(t, destination, M.SocksaddrFromNet(from))
				require.Equal(t, payload, response[:n])
			}
		})
	}
}

type xudpTestDialer struct {
	N.Dialer
	conn net.Conn
}

func (d xudpTestDialer) DialContext(context.Context, string, M.Socksaddr) (net.Conn, error) {
	return d.conn, nil
}

type xudpEcho struct{}

func (xudpEcho) NewConnectionEx(_ context.Context, conn net.Conn, _, _ M.Socksaddr, _ N.CloseHandlerFunc) {
	conn.Close()
}

func (xudpEcho) NewPacketConnectionEx(_ context.Context, conn N.PacketConn, _, _ M.Socksaddr, onClose N.CloseHandlerFunc) {
	go func() {
		packet := buf.NewPacket()
		destination, err := conn.ReadPacket(packet)
		if err == nil {
			err = conn.WritePacket(packet, destination)
		} else {
			packet.Release()
		}
		conn.Close()
		if onClose != nil {
			onClose(err)
		}
	}()
}
