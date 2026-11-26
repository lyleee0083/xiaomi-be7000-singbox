package vless

import (
	"context"
	"net"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/adapter/outbound"
	"github.com/sagernet/sing-box/common/dialer"
	"github.com/sagernet/sing-box/common/tls"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing-vmess"
	"github.com/sagernet/sing-vmess/vless"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/logger"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

func RegisterOutbound(registry *outbound.Registry) {
	outbound.Register[option.VLESSOutboundOptions](registry, C.TypeVLESS, NewOutbound)
}

type Outbound struct {
	outbound.Adapter
	logger     logger.ContextLogger
	dialer     N.Dialer
	client     *vless.Client
	serverAddr M.Socksaddr
	tlsDialer  tls.Dialer
}

func NewOutbound(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.VLESSOutboundOptions) (adapter.Outbound, error) {
	if options.PacketEncoding != nil && *options.PacketEncoding != "xudp" {
		return nil, E.New("BE7000 VLESS requires XUDP packet encoding")
	}
	outboundDialer, err := dialer.New(ctx, options.DialerOptions, options.ServerIsDomain())
	if err != nil {
		return nil, err
	}
	outbound := &Outbound{
		Adapter:    outbound.NewAdapterWithDialerOptions(C.TypeVLESS, tag, options.Network.Build(), options.DialerOptions),
		logger:     logger,
		dialer:     outboundDialer,
		serverAddr: options.ServerOptions.Build(),
	}
	if options.TLS != nil {
		tlsConfig, err := tls.NewClientWithOptions(tls.ClientOptions{
			Context:        ctx,
			Logger:         logger,
			ServerAddress:  options.Server,
			Options:        *options.TLS,
			DomainResolver: options.DomainResolver,
			KTLSCompatible: options.Flow == "",
		})
		if err != nil {
			return nil, err
		}
		if tlsConfig != nil {
			outbound.tlsDialer = tls.NewDialer(outboundDialer, tlsConfig)
		}
	}
	outbound.client, err = vless.NewClient(options.UUID, options.Flow, logger)
	if err != nil {
		return nil, err
	}
	return outbound, nil
}

func (h *Outbound) DialContext(ctx context.Context, network string, destination M.Socksaddr) (net.Conn, error) {
	switch N.NetworkName(network) {
	case N.NetworkTCP:
		ctx, conn, err := h.dialProxy(ctx, destination)
		if err != nil {
			return nil, err
		}
		h.logger.InfoContext(ctx, "outbound connection to ", destination)
		relayConn, err := h.client.DialEarlyConn(conn, destination)
		if err != nil {
			conn.Close()
			return nil, err
		}
		return relayConn, nil
	case N.NetworkUDP:
		return h.listenXUDP(ctx, destination)
	default:
		return nil, E.Extend(N.ErrUnknownNetwork, network)
	}
}

func (h *Outbound) ListenPacket(ctx context.Context, destination M.Socksaddr) (net.PacketConn, error) {
	return h.listenXUDP(ctx, destination)
}

func (h *Outbound) listenXUDP(ctx context.Context, destination M.Socksaddr) (vmess.PacketConn, error) {
	ctx, conn, err := h.dialProxy(ctx, destination)
	if err != nil {
		return nil, err
	}
	h.logger.InfoContext(ctx, "outbound packet connection to ", destination)
	packetConn, err := h.client.DialEarlyXUDPPacketConn(conn, destination)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return packetConn, nil
}

func (h *Outbound) dialProxy(ctx context.Context, destination M.Socksaddr) (context.Context, net.Conn, error) {
	ctx, metadata := adapter.ExtendContext(ctx)
	metadata.Outbound = h.Tag()
	metadata.Destination = destination
	var conn net.Conn
	var err error
	if h.tlsDialer != nil {
		conn, err = h.tlsDialer.DialTLSContext(ctx, h.serverAddr)
	} else {
		conn, err = h.dialer.DialContext(ctx, N.NetworkTCP, h.serverAddr)
	}
	return ctx, conn, err
}
