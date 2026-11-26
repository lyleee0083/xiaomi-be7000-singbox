package include

import (
	"context"

	"github.com/sagernet/sing-box"
	"github.com/sagernet/sing-box/adapter/certificate"
	"github.com/sagernet/sing-box/adapter/endpoint"
	"github.com/sagernet/sing-box/adapter/inbound"
	"github.com/sagernet/sing-box/adapter/outbound"
	"github.com/sagernet/sing-box/adapter/service"
	"github.com/sagernet/sing-box/dns"
	"github.com/sagernet/sing-box/dns/transport"
	"github.com/sagernet/sing-box/dns/transport/hosts"
	"github.com/sagernet/sing-box/dns/transport/local"
	"github.com/sagernet/sing-box/protocol/direct"
	"github.com/sagernet/sing-box/protocol/group"
	"github.com/sagernet/sing-box/protocol/socks"
	"github.com/sagernet/sing-box/protocol/tun"
	"github.com/sagernet/sing-box/protocol/vless"
)

func Context(ctx context.Context) context.Context {
	return box.Context(ctx, InboundRegistry(), OutboundRegistry(), EndpointRegistry(), DNSTransportRegistry(), ServiceRegistry(), CertificateProviderRegistry())
}

func InboundRegistry() *inbound.Registry {
	registry := inbound.NewRegistry()
	direct.RegisterInbound(registry)
	socks.RegisterInbound(registry)
	tun.RegisterInbound(registry)
	return registry
}

func OutboundRegistry() *outbound.Registry {
	registry := outbound.NewRegistry()
	direct.RegisterOutbound(registry)
	group.RegisterSelector(registry)
	vless.RegisterOutbound(registry)
	return registry
}

func EndpointRegistry() *endpoint.Registry {
	return endpoint.NewRegistry()
}

func DNSTransportRegistry() *dns.TransportRegistry {
	registry := dns.NewTransportRegistry()
	transport.RegisterUDP(registry)
	transport.RegisterHTTPS(registry)
	hosts.RegisterTransport(registry)
	local.RegisterTransport(registry)
	return registry
}

func ServiceRegistry() *service.Registry {
	return service.NewRegistry()
}

func CertificateProviderRegistry() *certificate.Registry {
	return certificate.NewRegistry()
}
