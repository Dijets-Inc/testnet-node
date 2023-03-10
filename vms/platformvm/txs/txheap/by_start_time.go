// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package txheap

import (
	"time"

	"github.com/lasthyphen/dijetsnodego/vms/platformvm/txs"
)

var _ TimedHeap = (*byStartTime)(nil)

type TimedHeap interface {
	Heap

	Timestamp() time.Time
}

type byStartTime struct {
	txHeap
}

func NewByStartTime() TimedHeap {
	h := &byStartTime{}
	h.initialize(h)
	return h
}

func (h *byStartTime) Less(i, j int) bool {
	iTime := h.txs[i].tx.Unsigned.(txs.Staker).StartTime()
	jTime := h.txs[j].tx.Unsigned.(txs.Staker).StartTime()
	return iTime.Before(jTime)
}

func (h *byStartTime) Timestamp() time.Time {
	return h.Peek().Unsigned.(txs.Staker).StartTime()
}
