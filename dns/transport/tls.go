package transport

import (
	"context"
	"net"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/dialer"
	"github.com/sagernet/sing-box/common/tls"
	"github.com/sagernet/sing-box/dns"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/logger"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"

	mDNS "github.com/miekg/dns"
)

var (
	_ adapter.DNSTransport         = (*TLSTransport)(nil)
	_ adapter.IdleConnectionKeeper = (*TLSTransport)(nil)
)

type TLSTransport struct {
	dns.TransportAdapter
	logger      logger.ContextLogger
	dialer      tls.Dialer
	serverAddr  M.Socksaddr
	multiplexer *queryMultiplexer
}

func NewTLSRaw(logger logger.ContextLogger, adapter dns.TransportAdapter, dialer N.Dialer, serverAddr M.Socksaddr, tlsConfig tls.Config) *TLSTransport {
	t := &TLSTransport{
		TransportAdapter: adapter,
		logger:           logger,
		dialer:           tls.NewDialer(dialer, tlsConfig),
		serverAddr:       serverAddr,
	}
	t.multiplexer = newQueryMultiplexer(queryMultiplexerOptions{
		dial: func(ctx context.Context) (net.Conn, error) {
			conn, err := t.dialer.DialTLSContext(ctx, t.serverAddr)
			if err != nil {
				return nil, E.Cause(err, "dial TLS connection")
			}
			return conn, nil
		},
		write: func(conn net.Conn, message *mDNS.Msg, queryId uint16) error {
			return WriteMessage(conn, queryId, message)
		},
		readNext: func(conn net.Conn) (*mDNS.Msg, error) {
			return ReadMessage(conn)
		},
		retryReadError: true,
		probeReuse:     true,
	})
	return t
}

func (t *TLSTransport) Start(stage adapter.StartStage, scope *adapter.Scope) error {
	if stage != adapter.StartStateStart {
		return nil
	}
	scope.Add(t.multiplexer.Close)
	return dialer.InitializeDetour(t.dialer)
}

func (t *TLSTransport) Reset() {
	t.multiplexer.Reset()
}

func (t *TLSTransport) SetKeepIdleConnections(keep bool) {
	t.multiplexer.SetKeepIdleConnections(keep)
}

func (t *TLSTransport) CloseIdleConnections() {
	t.multiplexer.CloseIdleConnections()
}

func (t *TLSTransport) Exchange(ctx context.Context, message *mDNS.Msg) (*mDNS.Msg, error) {
	return t.multiplexer.Exchange(ctx, message)
}

func (t *TLSTransport) ExchangeAsync(ctx context.Context, message *mDNS.Msg, callback func(response *mDNS.Msg, err error)) {
	t.multiplexer.ExchangeAsync(ctx, message, callback)
}
