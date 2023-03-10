// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"testing"

	"github.com/lasthyphen/dijetsnodego/database/memdb"
	"github.com/lasthyphen/dijetsnodego/database/prefixdb"
	"github.com/lasthyphen/dijetsnodego/ids"
)

func TestSharedMemory(t *testing.T) {
	chainID0 := ids.GenerateTestID()
	chainID1 := ids.GenerateTestID()

	for _, test := range SharedMemoryTests {
		baseDB := memdb.New()

		memoryDB := prefixdb.New([]byte{0}, baseDB)
		testDB := prefixdb.New([]byte{1}, baseDB)

		m := NewMemory(memoryDB)

		sm0 := m.NewSharedMemory(chainID0)
		sm1 := m.NewSharedMemory(chainID1)

		test(t, chainID0, chainID1, sm0, sm1, testDB)
	}
}
