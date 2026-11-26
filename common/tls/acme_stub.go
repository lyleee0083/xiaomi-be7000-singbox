package tls

import (
	"context"

	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/logger"
)

func startACME(ctx context.Context, logger logger.Logger, options option.InboundACMEOptions) (managedCertificateProvider, error) {
	return nil, E.New(`ACME is unavailable in the BE7000 build`)
}
