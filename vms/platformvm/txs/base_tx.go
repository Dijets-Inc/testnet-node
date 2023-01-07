// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"
	"fmt"

	"github.com/lasthyphen/dijetsnodego/ids"
	"github.com/lasthyphen/dijetsnodego/snow"
	"github.com/lasthyphen/dijetsnodego/utils"
	"github.com/lasthyphen/dijetsnodego/utils/set"
	"github.com/lasthyphen/dijetsnodego/vms/components/djtx"
	"github.com/lasthyphen/dijetsnodego/vms/secp256k1fx"
)

var (
	ErrNilTx = errors.New("tx is nil")

	errOutputsNotSorted      = errors.New("outputs not sorted")
	errInputsNotSortedUnique = errors.New("inputs not sorted and unique")
)

// BaseTx contains fields common to many transaction types. It should be
// embedded in transaction implementations.
type BaseTx struct {
	djtx.BaseTx `serialize:"true"`

	// true iff this transaction has already passed syntactic verification
	SyntacticallyVerified bool `json:"-"`

	unsignedBytes []byte // Unsigned byte representation of this data
}

func (tx *BaseTx) Initialize(unsignedBytes []byte) {
	tx.unsignedBytes = unsignedBytes
}

func (tx *BaseTx) Bytes() []byte {
	return tx.unsignedBytes
}

func (tx *BaseTx) InputIDs() set.Set[ids.ID] {
	inputIDs := set.NewSet[ids.ID](len(tx.Ins))
	for _, in := range tx.Ins {
		inputIDs.Add(in.InputID())
	}
	return inputIDs
}

func (tx *BaseTx) Outputs() []*djtx.TransferableOutput {
	return tx.Outs
}

// InitCtx sets the FxID fields in the inputs and outputs of this [BaseTx]. Also
// sets the [ctx] to the given [vm.ctx] so that the addresses can be json
// marshalled into human readable format
func (tx *BaseTx) InitCtx(ctx *snow.Context) {
	for _, in := range tx.BaseTx.Ins {
		in.FxID = secp256k1fx.ID
	}
	for _, out := range tx.BaseTx.Outs {
		out.FxID = secp256k1fx.ID
		out.InitCtx(ctx)
	}
}

// SyntacticVerify returns nil iff this tx is well formed
func (tx *BaseTx) SyntacticVerify(ctx *snow.Context) error {
	switch {
	case tx == nil:
		return ErrNilTx
	case tx.SyntacticallyVerified: // already passed syntactic verification
		return nil
	}
	if err := tx.BaseTx.Verify(ctx); err != nil {
		return fmt.Errorf("metadata failed verification: %w", err)
	}
	for _, out := range tx.Outs {
		if err := out.Verify(); err != nil {
			return fmt.Errorf("output failed verification: %w", err)
		}
	}
	for _, in := range tx.Ins {
		if err := in.Verify(); err != nil {
			return fmt.Errorf("input failed verification: %w", err)
		}
	}
	switch {
	case !djtx.IsSortedTransferableOutputs(tx.Outs, Codec):
		return errOutputsNotSorted
	case !utils.IsSortedAndUniqueSortable(tx.Ins):
		return errInputsNotSortedUnique
	default:
		return nil
	}
}
