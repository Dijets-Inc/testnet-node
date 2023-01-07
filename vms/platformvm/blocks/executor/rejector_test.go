// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow"
	"github.com/lasthyphen/dijetsnodego/snow/choices"
	"github.com/lasthyphen/dijetsnodego/utils/logging"
	"github.com/lasthyphen/dijetsnodego/vms/components/verify"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/blocks"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/state"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/txs"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/txs/mempool"
	"github.com/lasthyphen/dijetsnodego/vms/secp256k1fx"
)

func TestRejectBlock(t *testing.T) {
	type test struct {
		name         string
		newBlockFunc func() (blocks.Block, error)
		rejectFunc   func(*rejector, blocks.Block) error
	}

	tests := []test{
		{
			name: "proposal block",
			newBlockFunc: func() (blocks.Block, error) {
				return blocks.NewBanffProposalBlock(
					time.Now(),
					ids.GenerateTestID(),
					1,
					&txs.Tx{
						Unsigned: &txs.AddDelegatorTx{
							// Without the line below, this function will error.
							DelegationRewardsOwner: &secp256k1fx.OutputOwners{},
						},
						Creds: []verify.Verifiable{},
					},
				)
			},
			rejectFunc: func(r *rejector, b blocks.Block) error {
				return r.BanffProposalBlock(b.(*blocks.BanffProposalBlock))
			},
		},
		{
			name: "atomic block",
			newBlockFunc: func() (blocks.Block, error) {
				return blocks.NewApricotAtomicBlock(
					ids.GenerateTestID(),
					1,
					&txs.Tx{
						Unsigned: &txs.AddDelegatorTx{
							// Without the line below, this function will error.
							DelegationRewardsOwner: &secp256k1fx.OutputOwners{},
						},
						Creds: []verify.Verifiable{},
					},
				)
			},
			rejectFunc: func(r *rejector, b blocks.Block) error {
				return r.ApricotAtomicBlock(b.(*blocks.ApricotAtomicBlock))
			},
		},
		{
			name: "standard block",
			newBlockFunc: func() (blocks.Block, error) {
				return blocks.NewBanffStandardBlock(
					time.Now(),
					ids.GenerateTestID(),
					1,
					[]*txs.Tx{
						{
							Unsigned: &txs.AddDelegatorTx{
								// Without the line below, this function will error.
								DelegationRewardsOwner: &secp256k1fx.OutputOwners{},
							},
							Creds: []verify.Verifiable{},
						},
					},
				)
			},
			rejectFunc: func(r *rejector, b blocks.Block) error {
				return r.BanffStandardBlock(b.(*blocks.BanffStandardBlock))
			},
		},
		{
			name: "commit",
			newBlockFunc: func() (blocks.Block, error) {
				return blocks.NewBanffCommitBlock(time.Now(), ids.GenerateTestID() /*parent*/, 1 /*height*/)
			},
			rejectFunc: func(r *rejector, blk blocks.Block) error {
				return r.BanffCommitBlock(blk.(*blocks.BanffCommitBlock))
			},
		},
		{
			name: "abort",
			newBlockFunc: func() (blocks.Block, error) {
				return blocks.NewBanffAbortBlock(time.Now(), ids.GenerateTestID() /*parent*/, 1 /*height*/)
			},
			rejectFunc: func(r *rejector, blk blocks.Block) error {
				return r.BanffAbortBlock(blk.(*blocks.BanffAbortBlock))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			blk, err := tt.newBlockFunc()
			require.NoError(err)

			mempool := mempool.NewMockMempool(ctrl)
			state := state.NewMockState(ctrl)
			blkIDToState := map[ids.ID]*blockState{
				blk.Parent(): nil,
				blk.ID():     nil,
			}
			rejector := &rejector{
				backend: &backend{
					ctx: &snow.Context{
						Log: logging.NoLog{},
					},
					blkIDToState: blkIDToState,
					Mempool:      mempool,
					state:        state,
				},
			}

			// Set expected calls on dependencies.
			for _, tx := range blk.Txs() {
				mempool.EXPECT().Add(tx).Return(nil).Times(1)
			}
			gomock.InOrder(
				state.EXPECT().AddStatelessBlock(blk, choices.Rejected).Times(1),
				state.EXPECT().Commit().Return(nil).Times(1),
			)

			err = tt.rejectFunc(rejector, blk)
			require.NoError(err)
			// Make sure block and its parent are removed from the state map.
			require.NotContains(rejector.blkIDToState, blk.ID())
		})
	}
}
