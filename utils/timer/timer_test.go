// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package timer

import (
	"sync"
	"testing"
	"time"
)

func TestTimer(*testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()

	timer := NewTimer(wg.Done)
	go timer.Dispatch()

	timer.SetTimeoutIn(time.Millisecond)
}
