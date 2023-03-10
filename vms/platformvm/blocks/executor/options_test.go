// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lasthyphen/dijetsnodego/snow/consensus/snowman"
	"github.com/lasthyphen/dijetsnodego/vms/platformvm/blocks"
)

func TestOptionsUnexpectedBlockType(t *testing.T) {
	tests := []blocks.Block{
		&blocks.BanffAbortBlock{},
		&blocks.BanffCommitBlock{},
		&blocks.BanffStandardBlock{},
		&blocks.ApricotAbortBlock{},
		&blocks.ApricotCommitBlock{},
		&blocks.ApricotStandardBlock{},
		&blocks.ApricotAtomicBlock{},
	}

	for _, blk := range tests {
		t.Run(fmt.Sprintf("%T", blk), func(t *testing.T) {
			err := blk.Visit(&options{})
			require.ErrorIs(t, err, snowman.ErrNotOracle)
		})
	}
}
