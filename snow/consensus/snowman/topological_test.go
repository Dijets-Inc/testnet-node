// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"testing"
)

func TestTopological(t *testing.T) {
	runConsensusTests(t, TopologicalFactory{})
}
