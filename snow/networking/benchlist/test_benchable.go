// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package benchlist

import (
	"testing"

	"github.com/lasthyphen/dijetsnodego/ids"
)

type TestBenchable struct {
	T *testing.T

	CantBenched, CantUnbenched bool
	BenchedF, UnbenchedF       func(chainID ids.ID, validatorID ids.NodeID)
}

// Default set the default callable value to [cant]
func (b *TestBenchable) Default(cant bool) {
	b.CantBenched = cant
	b.CantUnbenched = cant
}

func (b *TestBenchable) Benched(chainID ids.ID, validatorID ids.NodeID) {
	if b.BenchedF != nil {
		b.BenchedF(chainID, validatorID)
	} else if b.CantBenched && b.T != nil {
		b.T.Fatalf("Unexpectedly called Benched")
	}
}

func (b *TestBenchable) Unbenched(chainID ids.ID, validatorID ids.NodeID) {
	if b.UnbenchedF != nil {
		b.UnbenchedF(chainID, validatorID)
	} else if b.CantUnbenched && b.T != nil {
		b.T.Fatalf("Unexpectedly called Unbenched")
	}
}
