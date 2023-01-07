// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"context"

	"github.com/lasthyphen/dijetsnodego/snow/consensus/snowman"
)

var _ snowman.Block = (*BlockWrapper)(nil)

// BlockWrapper wraps a snowman Block while adding a smart caching layer to improve
// VM performance.
type BlockWrapper struct {
	snowman.Block

	state *State
}

// Verify verifies the underlying block, evicts from the unverified block cache
// and if the block passes verification, adds it to [cache.verifiedBlocks].
// Note: it is guaranteed that if a block passes verification it will be added to
// consensus and eventually be decided ie. either Accept/Reject will be called
// on [bw] removing it from [verifiedBlocks].
func (bw *BlockWrapper) Verify(ctx context.Context) error {
	if err := bw.Block.Verify(ctx); err != nil {
		// Note: we cannot cache blocks failing verification in case
		// the error is temporary and the block could become valid in
		// the future.
		return err
	}

	blkID := bw.ID()
	bw.state.unverifiedBlocks.Evict(blkID)
	bw.state.verifiedBlocks[blkID] = bw
	return nil
}

// Accept accepts the underlying block, removes it from verifiedBlocks, caches it as a decided
// block, and updates the last accepted block.
func (bw *BlockWrapper) Accept(ctx context.Context) error {
	blkID := bw.ID()
	delete(bw.state.verifiedBlocks, blkID)
	bw.state.decidedBlocks.Put(blkID, bw)
	bw.state.lastAcceptedBlock = bw

	return bw.Block.Accept(ctx)
}

// Reject rejects the underlying block, removes it from processing blocks, and caches it as a
// decided block.
func (bw *BlockWrapper) Reject(ctx context.Context) error {
	blkID := bw.ID()
	delete(bw.state.verifiedBlocks, blkID)
	bw.state.decidedBlocks.Put(blkID, bw)
	return bw.Block.Reject(ctx)
}
