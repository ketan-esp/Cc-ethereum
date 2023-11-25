// config.go

package consensus

import (
	"github.com/ethereum/go-ethereum/params"
)

// PoAConfig holds the configuration parameters for the PoA consensus.
type PoAConfig struct {
	params.ChainConfig
	Period uint64
}

// PoAChainConfig returns the PoA-specific chain configuration.
func PoAChainConfig(config *params.ChainConfig) PoAConfig {
	return PoAConfig{
		ChainConfig: *config,
		Period:      3, // Adjust as needed
	}
}
