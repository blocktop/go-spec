package spec

import "context"

type Controller interface {
	Start(context.Context)
	StartBlockchain(ctx context.Context, blockchainType string)
	StopBlockchain(blockchainType string)
	Stop()
}
