// consensus.go

package consensus

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// PoAValidator represents a validator (authority node) in the PoA network.
type PoAValidator struct {
	Address    string
	PrivateKey string // Hex encoded private key for simplicity
}

// SignSeal signs the seal for the proposed block.
func (v *PoAValidator) SignSeal(chain types.ChainReader, hash types.Hash) ([]byte, error) {
	privKey, err := crypto.HexToECDSA(v.PrivateKey)
	if err != nil {
		return nil, err
	}

	signature, err := crypto.Sign(hash.Bytes(), privKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// Propose proposes a new block and returns true if the proposal is accepted.
func (v *PoAValidator) Propose(chain types.ChainReader, header *types.Header, seal *types.BlockSeal) bool {
	// Example: Check if the block has a valid signature
	hash := chain.GetHeaderByNumber(header.Number.Uint64() - 1).Hash()
	signature, err := v.SignSeal(chain, hash)
	if err != nil {
		return false
	}

	// Verify the signature
	pubKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return false
	}

	// Example: Verify that the signature matches the validator's address
	addr := crypto.PubkeyToAddress(*pubKey)
	if addr.Hex() != v.Address {
		return false
	}

	// Example: Additional validation checks
	// Implement your custom logic

	return true
}

// Seal seals the proposed block.
func (v *PoAValidator) Seal(chain types.ChainReader, block *types.Block, results []*types.Receipt) (*types.BlockSeal, error) {
	// Example: Create a new seal for the proposed block
	hash := chain.GetHeaderByNumber(block.NumberU64() - 1).Hash()
	signature, err := v.SignSeal(chain, hash)
	if err != nil {
		return nil, err
	}

	seal := &types.BlockSeal{
		Hash:      hash,
		Signature: signature,
	}

	return seal, nil
}
