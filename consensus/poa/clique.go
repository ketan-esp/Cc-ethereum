// clique.go

package consensus

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/crypto"
)

// PoA is a basic PoA-based consensus engine.
type PoA struct {
	clique.Clique
}

// NewPoA returns a new PoA consensus engine.
func NewPoA(config *params.ChainConfig, db ethdb.Database) *PoA {
	poa := &PoA{}
	poa.Clique = *clique.New(db, config, nil)
	return poa
}

// Author returns the PoA consensus author.
func (poa *PoA) Author(header *types.Header) (consensus.Author, error) {
	return poa.Clique.Author(header), nil
}
